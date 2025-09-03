package request

import (
	"github.com/wangluozhe/requests"
)

type wangluozheRequestSession struct {
	session *requests.Session
}

var WangluozheRequestSession = new(wangluozheRequestSession)

func (obj *wangluozheRequestSession) Start() {
	obj.session = requests.NewSession()
}
func (obj *wangluozheRequestSession) End() {}
func (obj *wangluozheRequestSession) Request(href string) ([]byte, error) {
	resp, err := obj.session.Get(href, nil)
	if err != nil {
		return nil, err
	}
	return resp.Content, err
}
