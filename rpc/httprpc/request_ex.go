package httprpc

import (
	"crypto/tls"
	"github.com/stdxgo/common-utils/logging"
	"net/http"
	"sync"
	"time"
)

var (
	timeoutMap     map[int]*http.Client
	timeoutMapLock sync.Mutex
)

func getTimeoutHTTPClient(seconds int) *http.Client {
	if seconds < 5 {
		seconds = 5
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
