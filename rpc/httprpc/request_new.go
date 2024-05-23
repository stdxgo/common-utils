package httprpc

import "net/http"

const (
	RPC_TRACE_KEY = "stdxgo-x-trace-id"
)

type Option func(ar *appRequest)

func Timeout(sec int) Option {
	return func(ar *appRequest) {
		ar.timeout = sec
		ar.client = getTimeoutHTTPClient(sec)
	}
}

func WithHTTPCli(cli *http.Client) Option {
	return func(ar *appRequest) {
		ar.client = cli
	}
}

func WithTraceKey(traceKey string) Option {
	return func(ar *appRequest) {
		ar.traceKey = traceKey
	}
}

// NewAppRequest 初始化一个默认的Request对象
func NewAppRequest(opts ...Option) AppRequest {

	r := &appRequest{
		timeout:  60,
		certPool: certPool,
		client:   cli,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
