package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/appointment-service/application"
)

func RegisterRoutes(e *echo.Echo, service *application.AppointmentService) {
	handler := NewAppointmentHandler(service)
	handler.RegisterRoutes(e)
}
