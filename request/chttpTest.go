package request

import (
	"io"

	tls "github.com/refraction-networking/utls"
	chttp "github.com/wangluozhe/chttp"
)

var chttpSession = &chttp.Client{
	Transport: &chttp.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func ChttpRequest(href string) ([]byte, error) {
	resp, err := chttpSession.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
