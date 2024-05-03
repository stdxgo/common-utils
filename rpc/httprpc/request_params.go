package httprpc

import (
	"fmt"
	"net/url"
	"strings"
)

// 配置 param
func (r *appRequest) SetParam(param map[string]string) AppRequest {

	if r.param == nil {
		r.param = make(map[string]string)
	}
	for k, v := range param {
		r.param[k] = url.QueryEscape(v)
	}
	return r
}

// 配置 pathParam
func (r *appRequest) SetPathParam(param map[string]string) AppRequest {

	if r.pathParam == nil {
		r.pathParam = make(map[string]string)
	}
	for k, v := range param {
		r.pathParam[k] = url.QueryEscape(v)
	}
	return r
}

func (r *appRequest) ClearAndSetParam(param map[string]string) AppRequest {

	r.ClearParam()
	r.SetParam(param)
	return r
}
func (r *appRequest) ClearParam() AppRequest {
	r.param = nil
	return r
}

// 配置 body // 覆盖配置
func (r *appRequest) SetBody(param interface{}) AppRequest {
	r.body = param
	return r
}
func (r *appRequest) ClearAndSetBody(param map[string]interface{}) AppRequest {
	r.ClearBody()
	r.SetBody(param)
	return r
}
func (r *appRequest) ClearBody() AppRequest {
	r.body = nil
	return r
}

// 将 param拼到url后
func (r *appRequest) buildParam(url string) string {
	if r.param == nil || len(r.param) == 0 {
		return url
	}
	query := []string{}
	for k, v := range r.param {
		if len(k) == 0 {
			continue
		}
		query = append(query, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("%s?%s", url, strings.Join(query, "&"))
}

// pathParam参数替换到url里
func (r *appRequest) buildPathParam(url string) string {
	if r.pathParam == nil || len(r.pathParam) == 0 {
		return url
	}
	for k, v := range r.pathParam {
		url = strings.Replace(url, fmt.Sprintf("{%s}", k), v, -1)
	}
	return url
}
