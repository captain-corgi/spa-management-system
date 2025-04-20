package interfaces

import (
	"github.com/labstack/echo/v4"
	"../application"
)

func RegisterRoutes(e *echo.Echo, service *application.StaffService) {
	handler := NewStaffHandler(service)
	handler.RegisterRoutes(e)
}
