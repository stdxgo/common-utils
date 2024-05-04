package httprpc

import (
	"context"
	"crypto/tls"
	"github.com/stdxgo/common-utils/excontext"
	"github.com/stdxgo/common-utils/logging"
	"net/http"
	"sync"
	"time"
)

var (
	timeoutMap     map[int]*http.Client
	timeoutMapLock sync.Mutex
)

// traceId
func (r *appRequest) addTraceId2Header(ctx context.Context, req *http.Request) {
	if tid := excontext.GetTraceID(ctx); tid != "" {
		if r.traceKey == "" {
			r.traceKey = RPC_TRACE_KEY
		}
		req.Header.Add(r.traceKey, tid)
	}
}

func getTimeoutHTTPClient(seconds int) *http.Client {
	if seconds <= 5 {
		return cliFiveSeconds
	}
	if seconds == 60 {
		return cli
	}
	timeoutMapLock.Lock()
	defer timeoutMapLock.Unlock()
	if timeoutMap[seconds] != nil {
		return timeoutMap[seconds]
	}
	c := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
			TLSClientConfig: &tls.Config{
				RootCAs:            certPool,
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * time.Duration(seconds),
	}
	timeoutMap[seconds] = c
	return c
}

func init() {
	timeoutMap = make(map[int]*http.Client)
	timeoutMapLock.Lock()
	defer timeoutMapLock.Unlock()
}

var (
	errorf = logging.Errorf
)
