package request

import (
	"bytes"
	"io"

	"net/http"

	"crypto/rand"
	"crypto/tls"
)

var httpSession = &http.Client{
	Transport: &http.Transport{
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
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
func HttpRequest2(href string) ([]byte, error) {
	// randomBytes := make([]byte, 1024*10000)
	randomBytes := make([]byte, 1024*100000)
	// randomBytes := make([]byte, 1024*100000)
	io.ReadFull(rand.Reader, randomBytes)
	resp, err := httpSession.Post(href, "stream/octet-stream", bytes.NewReader(randomBytes)) // Treat the package name as a Request, send GET request.
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	con, err := io.ReadAll(resp.Body)
	return con, err
}
