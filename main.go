package main

import (
	"benchmark/request"
	"benchmark/server"
	"benchmark/tools"
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func oneThreadTest(requestFuncs []requestFunc, routes []string, http2 bool, total int, threadNum int) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	addr := "127.0.0.1:3334"
	ctx, cnl := context.WithCancel(context.TODO())
	defer cnl()
	// svr := server.RequestServer2(ctx, DefaultTestData.Addr)
	// defer svr.Close()
	if http2 {
		log.Print("进行http2测试")
		svr := server.RequestServer2(ctx, addr)
		defer svr.Close()
	} else {
		log.Print("进行http1测试")
		svr := server.RequestServer(ctx, addr)
		defer svr.Close()
	}
	// log.Print("GospiderRequest 100k time: ", tools.TestMain(request.GospiderRequest2, DefaultTestData.Total, DefaultTestData.Addr+"/1000k", threadNum))
	// log.Print("GospiderRequest 100k time: ", tools.TestMain(request.HttpRequest2, DefaultTestData.Total, DefaultTestData.Addr+"/1000k", threadNum))
	// log.Print("GospiderRequest 1k time: ", tools.TestMain(request.HttpRequest2, DefaultTestData.Total, DefaultTestData.Addr+"/"+DefaultTestData.K, threadNum))
	// log.Print("GospiderRequest 1k time: ", tools.TestMain(request.GospiderRequest2, DefaultTestData.Total, DefaultTestData.Addr+"/"+DefaultTestData.K, threadNum))
	for _, route := range routes {
		log.Print("=======================")
		for _, rf := range requestFuncs {
			log.Print(rf.name, "  ", route, "  time: ", tools.TestMain(rf.f, total, addr+"/"+route, http2, threadNum))
		}
	}
}

type requestFunc struct {
	name string
	f    tools.Session
}

var requestFuncs = []requestFunc{
	{"Fingerproxy", request.FingerproxySession},
	// {"AzureTest", request.AzureSession},
	{"WangluozheRequest", request.WangluozheRequestSession},
	// {"ImrocReq", request.ImrocReqSession},
	{"GoResty", request.GoRestySession},
	{"Http", request.HttpSession},
}
var routes = []string{"1k", "10k", "100k"}

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	// }()
	// oneThreadTest(requestFuncs, routes, true, 10000, 100)
	oneThreadTest(requestFuncs, routes, false, 100000, 100)
	// oneThreadTest(requestFuncs, routes, false, 100000, 100)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// time.Sleep(time.Second * 10)
	fmt.Printf("程序结束时内存占用:\n")
	fmt.Printf("Alloc = %.2f MB\n", float64(m.Alloc)/1024/1024)           // 当前堆上分配的内存
	fmt.Printf("TotalAlloc = %.2f MB\n", float64(m.TotalAlloc)/1024/1024) // 程序运行过程中总分配的内存
	fmt.Printf("Sys = %.2f MB\n", float64(m.Sys)/1024/1024)               // Go运行时向操作系统申请的总内存
}
