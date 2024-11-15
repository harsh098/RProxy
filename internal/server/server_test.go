package server

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestReadinessProbeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	ReadinessProbeHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	body := w.Body.String()
	expectedBody := "pong"
	if body != expectedBody {
		t.Errorf("expected body %q; got %q", expectedBody, body)
	}
}

func TestProxyRequestHandler(t *testing.T) {
	targetURL, _ := url.Parse("http://example.com")
	reverseProxy := NewProxyServer(targetURL)

	handler := ProxyRequestHandler(reverseProxy, targetURL, "/api")

	req := httptest.NewRequest(http.MethodGet, "/api/resource", nil)
	req.Host = "localhost:8080"
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Verify URL modifications
	expectedHost := targetURL.Host
	if req.URL.Host != expectedHost {
		t.Errorf("expected host %q; got %q", expectedHost, req.URL.Host)
	}
	if req.URL.Scheme != "http" {
		t.Errorf("expected scheme 'http'; got %q", req.URL.Scheme)
	}

	// Verify Header Modifications
	expectedHeader := "localhost:8080"
	if req.Header.Get("X-Forwarded-Host") != expectedHeader {
		t.Errorf("expected X-Forwarded-Host header %q; got %q", expectedHeader, req.Header.Get("X-Forwarded-Host"))
	}
}
