package main

import (
	"benchmark/request"
	"benchmark/server"
	"benchmark/tools"
	"context"
	"log"
)

func oneThreadTest(threadNum int) {
	ctx, cnl := context.WithCancel(context.TODO())
	defer cnl()
	svr := server.RequestServer2(ctx, tools.DefaultTestData.Addr)
	defer svr.Close()
	// if tools.DefaultTestData.Http2 {
	// 	svr := server.RequestServer2(ctx, tools.DefaultTestData.Addr)
	// 	defer svr.Close()
	// } else {
	// 	svr := server.RequestServer(ctx, tools.DefaultTestData.Addr)
	// 	defer svr.Close()
	// }
	log.Print("GospiderRequest 1k time: ", tools.TestMain(request.GospiderRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/"+tools.DefaultTestData.K, threadNum))
	log.Print("WangluozheRequest 1k time: ", tools.TestMain(request.WangluozheRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/"+tools.DefaultTestData.K, threadNum))
	log.Print("ImrocReq 1k time: ", tools.TestMain(request.ImrocReq, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/"+tools.DefaultTestData.K, threadNum))
	log.Print("GoRestyRequest 1k time: ", tools.TestMain(request.GoRestyRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/"+tools.DefaultTestData.K, threadNum))
	// log.Print("=======================")
	// log.Print("GospiderRequest 10k time: ", tools.TestMain(request.GospiderRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/10k", threadNum))
	// log.Print("WangluozheRequest 10k time: ", tools.TestMain(request.WangluozheRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/10k", threadNum))
	// log.Print("ImrocReq 10k time: ", tools.TestMain(request.ImrocReq, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/10k", threadNum))
	// log.Print("GoRestyRequest 10k time: ", tools.TestMain(request.GoRestyRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/10k", threadNum))
	// log.Print("=======================")

	// log.Print("GospiderRequest 100k time: ", tools.TestMain(request.GospiderRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/1000k", threadNum))
	// log.Print("WangluozheRequest 100k time: ", tools.TestMain(request.WangluozheRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/1000000k", threadNum))
	// log.Print("ImrocReq 100k time: ", tools.TestMain(request.ImrocReq, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/1000000k", threadNum))
	// log.Print("GoRestyRequest 100k time: ", tools.TestMain(request.GoRestyRequest, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/1000000k", threadNum))

	// log.Print("GospiderRequest 100k time: ", tools.TestMain(request.GospiderRequest2, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/1000k", threadNum))
	// log.Print("GospiderRequest 100k time: ", tools.TestMain(request.HttpRequest2, tools.DefaultTestData.Total, tools.DefaultTestData.Addr+"/1000k", threadNum))
}

func main() {
	oneThreadTest(tools.DefaultTestData.Thread)
}
