package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/analytics-service/application"
)

func RegisterRoutes(e *echo.Echo, service *application.AnalyticsService) {
	handler := NewAnalyticsHandler(service)
	handler.RegisterRoutes(e)
}
