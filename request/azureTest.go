package request

import (
	"github.com/Noooste/azuretls-client"
	http "github.com/Noooste/fhttp"
	tls "github.com/Noooste/utls"
)

var azureSession *azuretls.Session

func init() {
	azureSession = azuretls.NewSession()
	azureSession.Transport = &http.Transport{
		MaxConnsPerHost:     10000000,
		MaxIdleConns:        1000000,
		MaxIdleConnsPerHost: 1000000,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}

func AzureTest(href string) ([]byte, error) {
	response, err := azureSession.Get(href)
	// response.CloseBody()
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
