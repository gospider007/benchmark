package request

import (
	"github.com/gospider007/requests"
)

func GospiderRequest(href string) ([]byte, error) {
	resp, err := requests.Get(nil, href)
	if err != nil {
		return nil, err
	}
	return resp.Content(), nil
}
