package request

import (
	azuretls "github.com/Noooste/azuretls-client"
	http "github.com/Noooste/fhttp"
	tls "github.com/Noooste/utls"
)

type azureSession struct {
	session *azuretls.Session
}

var AzureSession = new(azureSession)

func (obj *azureSession) Start() {
	obj.session = azuretls.NewSession()
	obj.session.Transport = &http.Transport{
		MaxConnsPerHost:     10000000,
		MaxIdleConns:        1000000,
		MaxIdleConnsPerHost: 1000000,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}
func (obj *azureSession) End() {}
func (obj *azureSession) Request(href string) ([]byte, error) {
	response, err := obj.session.Get(href)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
