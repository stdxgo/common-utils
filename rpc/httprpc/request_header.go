package httprpc

import (
	"github.com/stdxgo/common-utils/maputil"
	"net/http"
)

// RequestHeaderKey 请求header中，key的类型，等价于string
type RequestHeaderKey string

const (
	// ContentType header
	ContentType = "Content-Type"
)

var (
	defaultHeader = map[string]string{
		ContentType: "application/json",
	}
)

func (r *appRequest) SetRawHeaders(param http.Header) AppRequest {
	r.rawHeaders = param
	r.isRawHeaders = true
	return r
}

func (r *appRequest) ClearRawHeaders(param http.Header) AppRequest {
	r.rawHeaders = nil
	r.isRawHeaders = false
	return r
}

// 配置 headers
func (r *appRequest) SetHeaders(param map[string]string) AppRequest {

	if r.headers == nil {
		r.headers = param
		return r
	}
	for k, v := range param {
		r.headers[k] = v
	}
	return r
}

// SetHeadersIfNotExist 配置 headers，如果不存在才设置
func (r *appRequest) SetHeadersIfNotExist(param map[string]string) AppRequest {

	if r.headers == nil {
		r.headers = maputil.MapCopy(param)
		return r
	}
	for k, v := range param {
		if _, ok := r.headers[k]; !ok {
			r.headers[k] = v
		}
	}
	return r
}

// ClearAndSetHeaders 清除并设置header
func (r *appRequest) ClearAndSetHeaders(param map[string]string) AppRequest {

	r.ClearHeaders()
	r.SetHeaders(param)
	return r
}

// ClearHeaders 清除header
func (r *appRequest) ClearHeaders() AppRequest {
	r.headers = nil
	return r
}
