package server

import (
	"fmt"
	"github.com/harsh098/RPServer/internal/configs"
	"net/http"
	"net/url"
	"time"
)

func RunServer() error {
	config, err := configs.NewConfig()
	if err != nil {
		fmt.Errorf("[FATAL] %s: Failed to Read Config, Error Trace: %v", time.Now().UTC(), err)
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ReadinessProbeHandler)

	for _, resource := range config.Resources {
		url, _ := url.Parse(resource.Upstream_URL)
		reverseProxy := NewProxyServer(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(reverseProxy, url, resource.Endpoint))
	}

	socketAddress := fmt.Sprintf("%s:%s", config.Gateway.Host, config.Gateway.Listen_port)

	if err := http.ListenAndServe(socketAddress, mux); err != nil {
		fmt.Errorf("[FATAL] %s: Failed to Listen: %v", time.Now().UTC(), err)
	}
	return nil
}
