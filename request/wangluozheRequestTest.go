package request

import (
	"github.com/wangluozhe/requests"
)

var WangluozheRequestSession = requests.NewSession()

func WangluozheRequest(href string) ([]byte, error) {
	// headers := new(http.Header)
	// headers.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
	// resp, err := WangluozheRequestSession.Get(href, &url.Request{
	// 	Headers: headers,
	// })
	resp, err := WangluozheRequestSession.Get(href, nil)
	if err != nil {
		return nil, err
	}
	return resp.Content, err
}
