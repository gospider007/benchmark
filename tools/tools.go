package tools

import (
	"context"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gospider007/bar"
	"github.com/gospider007/thread"
)

func PrintBar(ctx context.Context, total int, sucess *atomic.Int64) {
	go func() {
		barClient := bar.NewClient(int64(total))
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second * 3):
				barClient.Print(sucess.Load())
			}
		}
	}()
}
func test(ctx context.Context, f func(href string) ([]byte, error), href string, sucess *atomic.Int64) {
	con, err := f(href)
	if err != nil {
		log.Fatal(err)
		return
	}

	if strings.Contains(href, "1k") {
		if len(con) == 1024 {
			sucess.Add(1)
		} else {
			log.Print(len(con), "错误的长度")
		}
	} else if strings.Contains(href, "10k") {
		if len(con) == 10240 {
			sucess.Add(1)
		} else {
			log.Print(len(con))
		}
	} else if strings.Contains(href, "100k") {
		if len(con) == 102400 {
			sucess.Add(1)
		} else {
			log.Print(len(con))
		}
	} else if strings.Contains(href, "1000k") {
		if len(con) == 1024000 {
			sucess.Add(1)
		} else {
			log.Print(len(con))
		}
	} else if strings.Contains(href, "10000k") {
		if len(con) == 10240000 {
			sucess.Add(1)
		} else {
			log.Print(len(con))
		}
	} else if strings.Contains(href, "100000k") {
		if len(con) == 102400000 {
			sucess.Add(1)
		} else {
			log.Print(len(con))
		}
	} else if strings.Contains(href, "1000000k") {
		if len(con) == 1024000000 {
			sucess.Add(1)
		} else {
			log.Print(len(con))
		}
	} else {
		log.Panic(href, len(con))
	}
}

type Session interface {
	Start()
	End()
	Request(href string) ([]byte, error)
}

func TestMain(f Session, total int, href string, Http2 bool, threadNum ...int) time.Duration {
	f.Start()
	defer f.End()
	if Http2 {
		href = "https://" + href
	} else {
		href = "http://" + href
	}
	// href = "https://" + href
	if len(threadNum) > 0 && threadNum[0] > 1 {
		return TestThreadMain(f.Request, total, href, threadNum[0])
	}
	var sucess atomic.Int64
	// PrintBar(ctx, total, &sucess)
	t := time.Now()
	for i := 0; i < total; i++ {
		log.Print(i, " start")
		// if i%10000 == 0 {
		// 	log.Print(i, " start")
		// }
		test(nil, f.Request, href, &sucess)
		log.Print(i, " end")

		// if i%10000 == 0 {
		// 	log.Print(i, " end")
		// }
	}
	t2 := time.Since(t)
	if sucess.Load() != int64(total) {
		log.Print(sucess.Load())
		log.Panic("fail")
	}
	return t2
}
func TestThreadMain(f func(href string) ([]byte, error), total int, href string, threadNum int) time.Duration {
	var sucess atomic.Int64
	t := time.Now()
	threadCli := thread.NewClient(nil, threadNum, thread.ClientOption{
		Debug: true,
	})
	for i := 0; i < total; i++ {
		_, err := threadCli.Write(nil, &thread.Task{
			Func: test,
			Args: []any{f, href, &sucess},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	threadCli.JoinClose()
	t2 := time.Since(t)
	if sucess.Load() != int64(total) {
		log.Panic("fail")
	}
	return t2
}
