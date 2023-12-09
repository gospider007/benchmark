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
		}
	} else if strings.Contains(href, "10k") {
		if len(con) == 10240 {
			sucess.Add(1)
		}
	} else if strings.Contains(href, "100k") {
		if len(con) == 102400 {
			sucess.Add(1)
		}
	}
}
func TestMain(f func(href string) ([]byte, error), total int, href string, threadNum ...int) time.Duration {
	href = "http://" + href
	// href = "https://" + href
	if len(threadNum) > 0 {
		return TestThreadMain(f, total, href, threadNum[0])
	}
	var sucess atomic.Int64
	// PrintBar(ctx, total, &sucess)
	t := time.Now()
	for i := 0; i < total; i++ {
		test(nil, f, href, &sucess)
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
	threadCli := thread.NewClient(nil, threadNum)
	for i := 0; i < total; i++ {
		threadCli.Write(&thread.Task{
			Func: test,
			Args: []any{f, href, &sucess},
		})
	}
	threadCli.JoinClose()
	t2 := time.Since(t)
	if sucess.Load() != int64(total) {
		log.Panic("fail")
	}
	return t2
}
