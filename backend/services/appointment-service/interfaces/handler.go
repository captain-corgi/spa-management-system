package interfaces

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/your-org/appointment-service/application"
	"github.com/your-org/appointment-service/domain"
)

type AppointmentHandler struct {
	service *application.AppointmentService
}

func NewAppointmentHandler(service *application.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

func (h *AppointmentHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/appointments")
	g.POST("", h.CreateAppointment)
	g.GET("", h.ListAppointments)
	g.GET(":id", h.GetAppointment)
	g.PUT(":id", h.UpdateAppointment)
	g.DELETE(":id", h.DeleteAppointment)
}

func (h *AppointmentHandler) CreateAppointment(c echo.Context) error {
	var appointment domain.Appointment
	if err := c.Bind(&appointment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateAppointment(&appointment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, appointment)
}

func (h *AppointmentHandler) ListAppointments(c echo.Context) error {
	list, err := h.service.ListAppointments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h *AppointmentHandler) GetAppointment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	appointment, err := h.service.GetAppointmentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) UpdateAppointment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var appointment domain.Appointment
	if err := c.Bind(&appointment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	appointment.ID = uint(id)
	err := h.service.UpdateAppointment(&appointment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, appointment)
}

func (h *AppointmentHandler) DeleteAppointment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteAppointment(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
