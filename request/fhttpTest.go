package request

import (
	"io"

	http "github.com/wangluozhe/fhttp"
)

var fhttpSession = &http.Client{}

func FhttpRequest(href string) ([]byte, error) {
	resp, err := fhttpSession.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
