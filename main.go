package main

import (
	"benchmark/request"
	"benchmark/server"
	"benchmark/tools"
	"context"
	"log"
)

func oneThreadTest() {
	ctx, cnl := context.WithCancel(context.TODO())
	defer cnl()
	svr := server.RequestServer(ctx, "127.0.0.1:3334")
	defer svr.Close()
	// time.Sleep(time.Hour)
	log.Print("GospiderRequest 1k time: ", tools.TestMain(request.GospiderRequest, 100000, "127.0.0.1:3334/1k"))
	log.Print("WangluozheRequest 1k time: ", tools.TestMain(request.WangluozheRequest, 100000, "127.0.0.1:3334/1k"))
	log.Print("ImrocReq 1k time: ", tools.TestMain(request.ImrocReq, 100000, "127.0.0.1:3334/1k"))
	log.Print("GoRestyRequest 1k time: ", tools.TestMain(request.GoRestyRequest, 100000, "127.0.0.1:3334/1k"))
	log.Print("=======================")
	log.Print("GospiderRequest 10k time: ", tools.TestMain(request.GospiderRequest, 100000, "127.0.0.1:3334/10k"))
	log.Print("WangluozheRequest 10k time: ", tools.TestMain(request.WangluozheRequest, 100000, "127.0.0.1:3334/10k"))
	log.Print("ImrocReq 10k time: ", tools.TestMain(request.ImrocReq, 100000, "127.0.0.1:3334/10k"))
	log.Print("GoRestyRequest 10k time: ", tools.TestMain(request.GoRestyRequest, 100000, "127.0.0.1:3334/10k"))
	log.Print("=======================")

	log.Print("GospiderRequest 100k time: ", tools.TestMain(request.GospiderRequest, 100000, "127.0.0.1:3334/100k"))
	log.Print("WangluozheRequest 100k time: ", tools.TestMain(request.WangluozheRequest, 100000, "127.0.0.1:3334/100k"))
	log.Print("ImrocReq 100k time: ", tools.TestMain(request.ImrocReq, 100000, "127.0.0.1:3334/100k"))
	log.Print("GoRestyRequest 100k time: ", tools.TestMain(request.GoRestyRequest, 100000, "127.0.0.1:3334/100k"))
}
func tempTest() {
	ctx, cnl := context.WithCancel(context.TODO())
	defer cnl()
	svr := server.RequestServer(ctx, "127.0.0.1:3334")
	defer svr.Close()
	// log.Print("WangluozheRequest 1k time: ", tools.TestMain(request.WangluozheRequest, 100000, "127.0.0.1:3334/1k")) //1并发
	log.Print("WangluozheRequest 1k time: ", tools.TestMain(request.WangluozheRequest, 100000, "127.0.0.1:3334/1k", 2)) //100并发

	// log.Print("GospiderRequest 1k time: ", tools.TestMain(request.GospiderRequest, 100000, "127.0.0.1:3334/1k")) //1并发
	// log.Print("GospiderRequest 1k time: ", tools.TestMain(request.GospiderRequest, 100000, "127.0.0.1:3334/1k", 100)) //100并发

	// log.Print("FhttpRequest 1k time: ", tools.TestMain(request.FhttpRequest, 100000, "127.0.0.1:3334/1k"))     //1 并发
	// log.Print("FhttpRequest 1k time: ", tools.TestMain(request.FhttpRequest, 100000, "127.0.0.1:3334/1k", 10)) //100并发

	// log.Print("ChttpRequest 1k time: ", tools.TestMain(request.ChttpRequest, 100000, "127.0.0.1:3334/1k"))     //1 并发
	// log.Print("ChttpRequest 1k time: ", tools.TestMain(request.ChttpRequest, 100000, "127.0.0.1:3334/1k", 10)) //100并发

	// log.Print("HttpRequest 1k time: ", tools.TestMain(request.HttpRequest, 100000, "127.0.0.1:3334/1k")) //1 并发
	// log.Print("HttpRequest 1k time: ", tools.TestMain(request.HttpRequest, 100000, "127.0.0.1:3334/1k", 10)) //100并发

	// log.Print("ImrocReq 1k time: ", tools.TestMain(request.ImrocReq, 100000, "127.0.0.1:3334/1k"))               //1 并发
	// log.Print("ImrocReq 1k time: ", tools.TestMain(request.ImrocReq, 100000, "127.0.0.1:3334/1k",100))               //1 并发
}
func main() {
	tempTest()
	// oneThreadTest()
}
