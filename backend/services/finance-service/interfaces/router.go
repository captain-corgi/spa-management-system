package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/finance-service/application"
)

func RegisterRoutes(e *echo.Echo, service *application.FinanceService) {
	handler := NewFinanceHandler(service)
	handler.RegisterRoutes(e)
}
