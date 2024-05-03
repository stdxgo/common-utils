package httprpc

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Header 请求过程中header的数据类型
type Header map[string][]string

// Response Request请求返回结果
type Response interface {
	GetBody() []byte
	UnmarshalBody(m interface{}) error
	GetStatus() string
	GetStatusCode() int
	GetHeader() Header
	Request() *http.Request
}

type response struct {
	status     string
	statusCode int
	body       []byte
	header     Header
	request    *http.Request
}

func (resp *response) Request() *http.Request {
	if resp == nil {
		return nil
	}
	return resp.request
}
func (resp *response) GetBody() []byte {

	if resp == nil || resp.body == nil {
		return []byte{}
	}
	return resp.body
}

func (resp *response) UnmarshalBody(m interface{}) error {
	if resp == nil || resp.body == nil {
		return errors.New("resp body blank")
	}
	return json.Unmarshal(resp.body, m)
}

func (resp *response) GetStatus() string {
	if resp == nil {
		return ""
	}
	return resp.status
}
func (resp *response) GetStatusCode() int {
	if resp == nil {
		return 0
	}
	return resp.statusCode
}

func (resp *response) GetHeader() Header {
	if resp == nil {
		return nil
	}
	return resp.header
}
