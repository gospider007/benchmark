package request

import (
	"crypto/rand"
	"io"
	"log"

	"github.com/gospider007/requests"
)

func GospiderRequest(href string) ([]byte, error) {
	resp, err := requests.Get(nil, href, requests.RequestOption{
		ClientOption: requests.ClientOption{
			ErrCallBack: func(ctx *requests.Response) error {
				log.Print(ctx.Err())
				return nil
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.Content(), nil
}
func GospiderRequest2(href string) ([]byte, error) {
	randomBytes := make([]byte, 1024*100000)
	// randomBytes := make([]byte, 1024*100000)
	io.ReadFull(rand.Reader, randomBytes)
	resp, err := requests.Post(nil, href, requests.RequestOption{
		Body: randomBytes,
		ClientOption: requests.ClientOption{
			ErrCallBack: func(ctx *requests.Response) error {
				log.Print(ctx.Err())
				return nil
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.Content(), nil
}
