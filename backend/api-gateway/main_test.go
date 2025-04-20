package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAuthMiddleware_AllowsRequest(t *testing.T) {
	e := echo.New()
	e.Use(AuthMiddleware)
	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp := httptest.NewRecorder()
	e.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.Code)
	}
}

func TestProxyTo_ServiceUnavailable(t *testing.T) {
	e := echo.New()
	e.GET("/customers/*", proxyTo("http://localhost:9999")) // unreachable port

	req := httptest.NewRequest(http.MethodGet, "/customers/1", nil)
	resp := httptest.NewRecorder()
	e.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadGateway {
		t.Errorf("expected status 502, got %d", resp.Code)
	}
}
