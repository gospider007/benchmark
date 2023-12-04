package tools

import (
	"context"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gospider007/bar"
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
func test(f func(href string) ([]byte, error), href string, sucess *atomic.Int64) {
	con, err := f(href)
	if err != nil {
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
func TestMain(f func(href string) ([]byte, error), total int, href string) time.Duration {
	var sucess atomic.Int64
	// PrintBar(ctx, total, &sucess)
	t := time.Now()
	for i := 0; i < total; i++ {
		test(f, href, &sucess)
	}
	t2 := time.Since(t)
	if sucess.Load() != int64(total) {
		log.Panic("fail")
	}
	return t2
}
