package request

import (
	"io"

	"net/http"
)

var httpSession = &http.Client{}

func HttpRequest(href string) ([]byte, error) {
	resp, err := httpSession.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
