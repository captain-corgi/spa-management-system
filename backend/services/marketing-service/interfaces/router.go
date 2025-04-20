package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/marketing-service/application"
)

func RegisterRoutes(e *echo.Echo, service *application.MarketingService) {
	handler := NewMarketingHandler(service)
	handler.RegisterRoutes(e)
}
