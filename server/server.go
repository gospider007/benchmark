package server

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"net/http"
	"time"
)

func RequestServer(ctx context.Context, addr string) *http.Server {
	if addr == "" {
		addr = ":3334"
	}
	http.HandleFunc("/1k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*1)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})

	http.HandleFunc("/10k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*10)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})

	http.HandleFunc("/100k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*100)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	Server := &http.Server{
		Addr: addr,
	}
	go Server.ListenAndServe()
	// cert, err := gtls.CreateProxyCertWithName("test")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// Server := &http.Server{
	// 	Addr: addr,
	// 	TLSConfig: &tls.Config{
	// 		GetCertificate: func(chi *tls.ClientHelloInfo) (*tls.Certificate, error) {
	// 			return &cert, nil
	// 		},
	// 	},
	// }
	// go Server.ListenAndServeTLS("", "")

	log.Print("准备")
	time.Sleep(time.Second * 4)
	log.Print("准备 ok")
	return Server
}
