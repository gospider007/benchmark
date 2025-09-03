package request

import (
	"io"

	tls "github.com/refraction-networking/utls"
	chttp "github.com/wangluozhe/chttp"
)

type chttpSession struct {
	session *chttp.Client
}

var ChttpSession = new(chttpSession)

func (obj *chttpSession) Start() {
	obj.session = &chttp.Client{
		Transport: &chttp.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
func (obj *chttpSession) End() {}
func (obj *chttpSession) Request(href string) ([]byte, error) {
	resp, err := obj.session.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
