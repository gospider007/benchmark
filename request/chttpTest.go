package request

import (
	"io"

	http "github.com/wangluozhe/chttp"
)

var chttpSession = &http.Client{}

func ChttpRequest(href string) ([]byte, error) {
	resp, err := chttpSession.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
