package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Middleware: Logging, Recovery, CORS
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Middleware: Rate Limiting
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	// Middleware: Authentication (JWT)
	e.Use(AuthMiddleware)

	// API Docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Service Routing
	registerServiceRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

// AuthMiddleware is a stub for JWT auth
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Example: check Authorization header, validate JWT, set user context
		return next(c)
	}
}

// registerServiceRoutes configures proxy routes to backend microservices
func registerServiceRoutes(e *echo.Echo) {
	// Example: proxy /customers/* to customer-service
	e.Any("/customers/*", proxyTo("http://customer-service:8081"))
	e.Any("/appointments/*", proxyTo("http://appointment-service:8082"))
	e.Any("/staff/*", proxyTo("http://staff-service:8083"))
	e.Any("/branches/*", proxyTo("http://branch-service:8084"))
	e.Any("/marketing/*", proxyTo("http://marketing-service:8085"))
	e.Any("/finance/*", proxyTo("http://finance-service:8086"))
	e.Any("/analytics/*", proxyTo("http://analytics-service:8087"))
}

// proxyTo returns a handler that proxies the request to the given service URL
func proxyTo(serviceURL string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Simple reverse proxy (for demo); use a robust proxy in production
		resp, err := http.DefaultClient.Get(serviceURL + c.Request().URL.Path)
		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]string{"error": "service unavailable"})
		}
		defer resp.Body.Close()
		return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
	}
}
