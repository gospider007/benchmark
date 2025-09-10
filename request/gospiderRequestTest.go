package request

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
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
			if strings.Contains(err.Error(), "executable file not found") {
				log.Print("缺少依赖执行如下命令安装:  pip install --force-reinstall git+https://gitee.com/gospider007/fingerproxy.git")
				os.Exit(1)
			} else {
				log.Panic(err)
			}
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
