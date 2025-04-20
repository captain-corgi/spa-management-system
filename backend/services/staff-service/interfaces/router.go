package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/staff-service/application"
)

func RegisterRoutes(e *echo.Echo, service *application.StaffService) {
	handler := NewStaffHandler(service)
	handler.RegisterRoutes(e)
}
