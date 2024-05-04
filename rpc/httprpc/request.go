package httprpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/stdxgo/common-utils/crash"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var certPool = &x509.CertPool{}

var cli = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: false,
		TLSClientConfig: &tls.Config{
			RootCAs:            certPool,
			InsecureSkipVerify: true,
		},
	},
	Timeout: time.Second * 60,
}

// 调用外部接口超时时间设置为5s
var cliFiveSeconds = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: false,
		TLSClientConfig: &tls.Config{
			RootCAs:            certPool,
			InsecureSkipVerify: true,
		},
	},
	Timeout: time.Second * 5,
}

// AppRequest do a request
type AppRequest interface {
	AddCert2Pool(cert *x509.Certificate) AppRequest
	SetTimeout(second int) AppRequest
	DisabledKeepAlive(disabledKeepAlive bool) AppRequest

	SetParam(param map[string]string) AppRequest
	SetPathParam(param map[string]string) AppRequest
	ClearAndSetParam(param map[string]string) AppRequest
	ClearParam() AppRequest
	ClearBody() AppRequest
	SetBody(param interface{}) AppRequest
	ClearAndSetBody(param map[string]interface{}) AppRequest

	// SetRawHeaders 使用rawHeader时，会忽略正常设置的header
	SetRawHeaders(param http.Header) AppRequest
	// SetHeaders 设置header
	SetHeaders(param map[string]string) AppRequest
	// SetHeadersIfNotExist 如果header不存在才设置header
	SetHeadersIfNotExist(param map[string]string) AppRequest
	// ClearAndSetHeaders 清除旧header，并设置新header
	ClearAndSetHeaders(param map[string]string) AppRequest
	// ClearHeaders 清理header
	ClearHeaders() AppRequest

	DoGet(ctx context.Context, url string) (Response, error)
	DoPost(ctx context.Context, url string) (Response, error)
	DoPut(ctx context.Context, url string) (Response, error)
	DoDelete(ctx context.Context, url string) (Response, error)
}

type appRequest struct {
	client            *http.Client
	timeout           int
	headers           map[string]string
	isRawHeaders      bool
	rawHeaders        http.Header
	param             map[string]string
	body              interface{}
	disabledKeepAlive bool
	certPool          *x509.CertPool
	pathParam         map[string]string
	traceKey          string
}

func (r *appRequest) DisabledKeepAlive(disabledKeepAlive bool) AppRequest {
	r.disabledKeepAlive = disabledKeepAlive
	return r
}
func (r *appRequest) SetTimeout(second int) AppRequest {
	r.timeout = second
	r.client = getTimeoutHTTPClient(second)
	return r
}

// set cert pool
func (r *appRequest) AddCert2Pool(cert *x509.Certificate) AppRequest {
	if r.certPool == nil {
		r.certPool = certPool
	}
	r.certPool.AddCert(cert)
	return r
}

// build client
func (r *appRequest) buildClient() *http.Client {

	if r.client != nil {
		return r.client
	}
	if r.certPool == nil {
		r.certPool = certPool
	}
	r.client = cli
	return r.client
}

// build body

func buildBody(m interface{}) ([]byte, error) {

	if ret, isByteArr := m.([]byte); isByteArr {
		return ret, nil
	}
	if m == nil || reflect.ValueOf(m).IsZero() {
		return []byte{}, nil
	}
	return json.Marshal(m)
}
func add2Header(r *http.Request, m map[string]string) {
	if r == nil || len(m) == 0 {
		return
	}
	for k, v := range m {
		r.Header.Set(k, v)
	}
}

// build req
func (r *appRequest) buildRequest(ctx context.Context, url string, method string) (*http.Request, error) {

	body, err := buildBody(r.body)
	if err != nil {
		err = fmt.Errorf("构造请求body失败(%s)", err.Error())
		return nil, err
	}

	// body
	bodyReader := strings.NewReader(fmt.Sprintf("%s", body))
	request, err := http.NewRequest(strings.ToUpper(method), url, bodyReader)
	if err != nil {
		return request, fmt.Errorf("构造Request失败(%s)", err.Error())
	}
	if request == nil {
		return nil, fmt.Errorf("构造Request失败")
	}

	// headers
	defaultHeaders := make(map[string]string, 10)

	if len(r.headers) > 0 && !r.isRawHeaders {
		add2Header(request, r.headers)
	}
	if r.isRawHeaders {
		request.Header = r.rawHeaders
	}
	add2Header(request, defaultHeaders)
	r.addTraceId2Header(ctx, request)
	return request, nil
}

func (r *appRequest) DoGet(ctx context.Context, url string) (Response, error) {

	return r.doRequest(ctx, url, "GET")
}

func (r *appRequest) DoPost(ctx context.Context, url string) (Response, error) {

	r.SetHeadersIfNotExist(defaultHeader)
	return r.doRequest(ctx, url, "POST")
}

func (r *appRequest) DoPut(ctx context.Context, url string) (Response, error) {

	r.SetHeadersIfNotExist(defaultHeader)
	return r.doRequest(ctx, url, "PUT")
}

func (r *appRequest) DoDelete(ctx context.Context, url string) (Response, error) {

	r.SetHeadersIfNotExist(defaultHeader)
	return r.doRequest(ctx, url, "DELETE")
}

func (r *appRequest) doRequest(ctx context.Context, url, method string) (resp Response, err error) {
	method = strings.ToUpper(method)
	// build pathParam
	url = r.buildPathParam(url)
	// build param
	url = r.buildParam(url)

	request, err := r.buildRequest(ctx, url, method)
	if err != nil {
		return nil, fmt.Errorf("请求失败(%s:%s):%s", method, url, err.Error())
	}
	defer crash.Handler(ctx, func() {
		bd, _ := buildBody(r.body)
		resp = nil
		err = fmt.Errorf("http crash : %s body(%s)", url, string(bd))
	})
	defer func() {
		tmpErr := request.Body.Close()
		if tmpErr != nil {
			errorf(ctx, "请求结束时Request.Body.Close错误:%s", tmpErr.Error())
		}
	}()

	r.client = r.buildClient()
	res, err := r.client.Do(request)

	if err != nil {
		return nil, err
	}
	defer func() {
		tmpErr := res.Body.Close()
		if tmpErr != nil {
			errorf(ctx, "请求结束时Response.Body.Close错误:%s", tmpErr.Error())
		}
	}()
	var bys []byte

	for {
		var (
			tmp     = make([]byte, 1024)
			bodyLen int
		)

		bodyLen, err = res.Body.Read(tmp)
		if bodyLen == 0 || bodyLen == -1 {
			break
		}
		if err != nil && err.Error() != "EOF" {
			errorf(ctx, "读取请求结果读取失败:%s", err.Error())
			return nil, err
		}
		bys = append(bys, tmp[0:bodyLen]...)
	}

	respX := new(response)
	respX.body = bys
	respX.status = res.Status
	respX.statusCode = res.StatusCode
	respX.header = make(Header)
	respX.request = request
	for k, v := range res.Header {
		respX.header[k] = v
	}
	return respX, nil
}
