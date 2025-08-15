package main

import (
	"benchmark/request"
	"benchmark/server"
	"benchmark/tools"
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
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
		svr := server.RequestServer2(ctx, addr)
		defer svr.Close()
	} else {
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
	f    func(href string) ([]byte, error)
}

// {"GospiderRequest2", request.GospiderRequest2},

var requestFuncs = []requestFunc{
	{"GospiderRequest", request.GospiderRequest},
	{"AzureTest", request.AzureTest},
	{"WangluozheRequest", request.WangluozheRequest},
	{"ImrocReq", request.ImrocReq},
	{"GoRestyRequest", request.GoRestyRequest},
}
var routes = []string{"1k", "10k", "100k"}

func main() {
	// oneThreadTest(requestFuncs, routes, true, 10000, 10)
	// oneThreadTest(requestFuncs, routes, false, 10000, 10)
}
