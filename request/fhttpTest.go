package request

import (
	"io"

	tls "github.com/refraction-networking/utls"
	fhttp "github.com/wangluozhe/fhttp"
)

type fhttpSession struct {
	session *fhttp.Client
}

func (obj *fhttpSession) Start() {
	obj.session = &fhttp.Client{
		Transport: &fhttp.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
func (obj *fhttpSession) End() {}
func (obj *fhttpSession) Request(href string) ([]byte, error) {
	resp, err := obj.session.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

var FhttpSession = new(fhttpSession)
