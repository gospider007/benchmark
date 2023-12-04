package main

import (
	"benchmark/request"
	"benchmark/server"
	"benchmark/tools"
	"context"
	"log"
	"time"
)

func main() {
	ctx, cnl := context.WithCancel(context.TODO())
	defer cnl()
	svr := server.RequestServer(ctx, "127.0.0.1:3334")
	defer svr.Close()
	time.Sleep(time.Hour)
	// log.Print("GospiderRequest 1k time: ", tools.TestMain(request.GospiderRequest, 100000, "http://127.0.0.1:3334/1k"))
	// log.Print("WangluozheRequest 1k time: ", tools.TestMain(request.WangluozheRequest, 100000, "http://127.0.0.1:3334/1k"))
	// log.Print("ImrocReq 1k time: ", tools.TestMain(request.ImrocReq, 100000, "http://127.0.0.1:3334/1k"))
	// log.Print("GoRestyRequest 1k time: ", tools.TestMain(request.GoRestyRequest, 100000, "http://127.0.0.1:3334/1k"))

	// log.Print("GospiderRequest 10k time: ", tools.TestMain(request.GospiderRequest, 100000, "http://127.0.0.1:3334/10k"))
	// log.Print("WangluozheRequest 10k time: ", tools.TestMain(request.WangluozheRequest, 100000, "http://127.0.0.1:3334/10k"))
	// log.Print("ImrocReq 10k time: ", tools.TestMain(request.ImrocReq, 100000, "http://127.0.0.1:3334/10k"))
	// log.Print("GoRestyRequest 10k time: ", tools.TestMain(request.GoRestyRequest, 100000, "http://127.0.0.1:3334/10k"))

	log.Print("GospiderRequest 100k time: ", tools.TestMain(request.GospiderRequest, 100000, "http://127.0.0.1:3334/100k"))
	log.Print("WangluozheRequest 100k time: ", tools.TestMain(request.WangluozheRequest, 100000, "http://127.0.0.1:3334/100k"))
	log.Print("ImrocReq 100k time: ", tools.TestMain(request.ImrocReq, 100000, "http://127.0.0.1:3334/100k"))
	log.Print("GoRestyRequest 100k time: ", tools.TestMain(request.GoRestyRequest, 100000, "http://127.0.0.1:3334/100k"))
}
