package request

import (
	"io"

	"net/http"
)

var httpSession = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
	},
}

func HttpRequest(href string) ([]byte, error) {
	resp, err := httpSession.Get(href) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	con, err := io.ReadAll(resp.Body)
	return con, err
}
