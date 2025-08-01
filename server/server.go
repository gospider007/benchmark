package server

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gospider007/gtls"
	"golang.org/x/net/http2"
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
		// if r.Body != nil && r.ContentLength > 0{
		// 	ccc, err := io.ReadAll(r.Body)
		// 	log.Print("post body len: ===============", err, len(ccc))
		// }
		randomBytes := make([]byte, 1024*100)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/1000k", func(w http.ResponseWriter, r *http.Request) {
		// if r.Body != nil && r.ContentLength > 0{
		// 	ccc, err := io.ReadAll(r.Body)
		// 	log.Print("post body len: ===============", err, len(ccc))
		// }
		randomBytes := make([]byte, 1024*1000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/10000k", func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil && r.ContentLength > 0 {
			io.Copy(io.Discard, r.Body)
		}
		randomBytes := make([]byte, 1024*10000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/100000k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*100000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/1000000k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*1000000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	Server := &http.Server{
		Addr: addr,
	}
	go Server.ListenAndServe()
	log.Print("准备")
	time.Sleep(time.Second * 4)
	log.Print("准备 ok")
	return Server
}
func RequestServer2(ctx context.Context, addr string) *http.Server {
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
		// if r.Body != nil && r.ContentLength > 0{
		// 	ccc, err := io.ReadAll(r.Body)
		// 	log.Print("post body len: ===============", err, len(ccc))
		// }
		randomBytes := make([]byte, 1024*100)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/1000k", func(w http.ResponseWriter, r *http.Request) {
		// if r.Body != nil && r.ContentLength > 0{
		// 	ccc, err := io.ReadAll(r.Body)
		// 	log.Print("post body len: ===============", err, len(ccc))
		// }
		randomBytes := make([]byte, 1024*1000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/10000k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*10000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/100000k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*100000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	http.HandleFunc("/1000000k", func(w http.ResponseWriter, r *http.Request) {
		randomBytes := make([]byte, 1024*1000000)
		io.ReadFull(rand.Reader, randomBytes)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(randomBytes)
	})
	Server := &http.Server{
		Addr: addr,
		TLSConfig: gtls.GetCertConfigForClient(&tls.Config{
			NextProtos:         []string{"h2"},
			InsecureSkipVerify: true,
		}),
	}
	// err = http.Http2ConfigureServer(Server, nil)
	err := http2.ConfigureServer(Server, nil)
	if err != nil {
		log.Fatal(err)
	}
	go Server.ListenAndServeTLS("", "")
	log.Print("准备")
	time.Sleep(time.Second * 4)
	log.Print("准备2 ok")
	return Server
}

// _,ok:=f.(*http2RSTStreamFrame)
// if ok{
// 	log.Print("我开始行动了")
// 	var val string
// 	fmt.Scanln(&val)
// 	debug.PrintStack()
// }

// 2025/02/16 20:45:38 endWrite: [0 0 64 1 4 0 0 0 1 65 138 8 157 92 11 129 112 220 101 150 90 131 69 133 96 32 0 58 255 135 95 141 66 108 40 233 96 228 73 82 178 19 97 71 79 92 134 8 4 208 0 0 127 80 131 155 217 171 122 141 196 117 167 74 107 88 148 24 181 37 129 46 15]
// 2025/02/16 20:45:38 endWrite: [0 0 64 1 4 0 0 0 1 65 138 8 157 92 11 129 112 220 101 150 90 131 69 133 96 32 0 58 255 135 95 141 66 108 40 233 96 228 73 82 178 19 97 71 79 92 134 8 4 208 0 0 127 80 131 155 217 171 122 141 196 117 167 74 107 88 148 24 181 37 129 46 15]
