package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/branch-service/application"
)

func RegisterRoutes(e *echo.Echo, service *application.BranchService) {
	handler := NewBranchHandler(service)
	handler.RegisterRoutes(e)
}
