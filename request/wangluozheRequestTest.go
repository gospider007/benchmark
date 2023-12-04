package request

import (
	"github.com/wangluozhe/requests"
)

var WangluozheRequestSession = requests.NewSession()

func WangluozheRequest(href string) ([]byte, error) {
	resp, err := WangluozheRequestSession.Get(href, nil)
	if err != nil {
		return nil, err
	}
	return resp.Content, err
}
