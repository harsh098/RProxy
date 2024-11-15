package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func ReadinessProbeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func NewProxyServer(target *url.URL) *httputil.ReverseProxy {
	rp := httputil.NewSingleHostReverseProxy(target)
	return rp
}

func ProxyRequestHandler(rp *httputil.ReverseProxy, url *url.URL, endpoint string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[INFO] %s: Recv req for URI:%s\n", time.Now().UTC(), r.URL)

		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Host = url.Host
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.URL.Path = strings.TrimPrefix(r.URL.Path, endpoint)

		fmt.Printf(" [INFO] %s: Proxying Request to Upstream URI:%s\n", time.Now().UTC(), r.URL)
	}
}
