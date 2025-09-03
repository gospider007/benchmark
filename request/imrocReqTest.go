package request

import (
	"github.com/imroc/req/v3"
)

type imrocReqSession struct {
	session *req.Client
}

var ImrocReqSession = new(imrocReqSession)

func (obj *imrocReqSession) Start() {
	obj.session = req.C()
	obj.session.Transport.MaxConnsPerHost = 1000000
	obj.session.Transport.MaxIdleConnsPerHost = 1000000
	obj.session.Transport.MaxIdleConns = 1000000
	obj.session.EnableInsecureSkipVerify()
}
func (obj *imrocReqSession) End() {}
func (obj *imrocReqSession) Request(href string) ([]byte, error) {
	resp := obj.session.R().MustGet(href) // Treat the package name as a Request, send GET request.
	return resp.Bytes(), nil
}
