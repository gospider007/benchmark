package request

import (
	"crypto/tls"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type goRestySession struct {
	session *resty.Client
}

var GoRestySession = new(goRestySession)

func (obj *goRestySession) Start() {
	obj.session = resty.New()
	obj.session.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	obj.session.SetTransport(&http.Transport{
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        100000,
		MaxIdleConnsPerHost: 100000,
		MaxConnsPerHost:     100000,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
}
func (obj *goRestySession) End() {}
func (obj *goRestySession) Request(href string) ([]byte, error) {
	resp, err := obj.session.R().Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
