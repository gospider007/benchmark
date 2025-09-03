package request

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"net/http"
	"net/url"

	"crypto/tls"
)

type fingerproxySession struct {
	session *http.Client
	cmd     *exec.Cmd
	closed  bool
}

var FingerproxySession = new(fingerproxySession)

func (obj *fingerproxySession) Start() {
	obj.cmd = exec.Command("fingerproxy") // 你想执行的命令
	go func() {
		err := obj.cmd.Run() // Run 会阻塞，直到命令执行完
		if !obj.closed && err != nil {
			fmt.Println("命令执行出错:")
			log.Print("请确保安装了 fingerproxy :\n pip install --force-reinstall git+https://gitee.com/gospider007/fingerproxy.git")
			log.Panic(err)
		}
	}()
	time.Sleep(time.Second * 2)
	obj.session = &http.Client{
		Transport: &http.Transport{
			Proxy: func(*http.Request) (*url.URL, error) {
				return url.Parse("http://127.0.0.1:8080")
			},
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 1000,
			MaxConnsPerHost:     1000,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
func (obj *fingerproxySession) End() {
	obj.closed = true
	obj.cmd.Process.Kill() // 结束命令
}
func (obj *fingerproxySession) Request(href string) ([]byte, error) {
	resp, err := obj.session.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	con, err := io.ReadAll(resp.Body)
	return con, err
}
