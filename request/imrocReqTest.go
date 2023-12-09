package request

import (
	"github.com/imroc/req/v3"
)

var imrocReqSession = req.C()

func init() {
	imrocReqSession.Transport.MaxConnsPerHost = 1000000
	imrocReqSession.Transport.MaxIdleConnsPerHost = 1000000
	imrocReqSession.Transport.MaxIdleConns = 1000000
	imrocReqSession.EnableInsecureSkipVerify()
}
func ImrocReq(href string) ([]byte, error) {
	resp := imrocReqSession.R().MustGet(href) // Treat the package name as a Request, send GET request.
	return resp.Bytes(), nil
}
