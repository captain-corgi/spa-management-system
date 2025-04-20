package interfaces

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/your-org/staff-service/application"
	"github.com/your-org/staff-service/domain"
)

type StaffHandler struct {
	service *application.StaffService
}

func NewStaffHandler(service *application.StaffService) *StaffHandler {
	return &StaffHandler{service: service}
}

func (h *StaffHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/staff")
	g.POST("", h.CreateStaff)
	g.GET("", h.ListStaff)
	g.GET(":id", h.GetStaff)
	g.PUT(":id", h.UpdateStaff)
	g.DELETE(":id", h.DeleteStaff)
}

func (h *StaffHandler) CreateStaff(c echo.Context) error {
	var staff domain.Staff
	if err := c.Bind(&staff); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateStaff(&staff)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, staff)
}

func (h *StaffHandler) ListStaff(c echo.Context) error {
	list, err := h.service.ListStaff()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h *StaffHandler) GetStaff(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	staff, err := h.service.GetStaffByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, staff)
}

func (h *StaffHandler) UpdateStaff(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var staff domain.Staff
	if err := c.Bind(&staff); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	staff.ID = uint(id)
	err := h.service.UpdateStaff(&staff)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, staff)
}

func (h *StaffHandler) DeleteStaff(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteStaff(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
