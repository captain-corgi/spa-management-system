package interfaces

import (
	"github.com/labstack/echo/v4"
	"../application"
)

func RegisterRoutes(e *echo.Echo, service *application.MarketingService) {
	handler := NewMarketingHandler(service)
	handler.RegisterRoutes(e)
}
