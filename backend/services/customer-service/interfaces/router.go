package interfaces

import (
	"github.com/labstack/echo/v4"
	"../application"
)

func RegisterRoutes(e *echo.Echo, service *application.CustomerService) {
	handler := NewCustomerHandler(service)
	handler.RegisterRoutes(e)
}
