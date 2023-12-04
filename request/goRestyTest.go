package request

import (
	"github.com/go-resty/resty/v2"
)

var goRestySession = resty.New()

func GoRestyRequest(href string) ([]byte, error) {
	resp, err := goRestySession.R().Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
