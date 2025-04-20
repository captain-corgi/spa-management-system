package interfaces

import (
	"github.com/labstack/echo/v4"
	"../application"
)

func RegisterRoutes(e *echo.Echo, service *application.BranchService) {
	handler := NewBranchHandler(service)
	handler.RegisterRoutes(e)
}
