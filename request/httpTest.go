package request

import (
	"io"

	"net/http"

	"crypto/tls"
)

type httpSession struct {
	session *http.Client
}

var HttpSession = new(httpSession)

func (obj *httpSession) Start() {
	obj.session = &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        100000,
			MaxIdleConnsPerHost: 100000,
			MaxConnsPerHost:     100000,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
func (obj *httpSession) End() {}
func (obj *httpSession) Request(href string) ([]byte, error) {
	resp, err := obj.session.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	con, err := io.ReadAll(resp.Body)
	return con, err
}
