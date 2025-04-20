package interfaces

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/your-org/customer-service/application"
	"github.com/your-org/customer-service/domain"
)

type CustomerHandler struct {
	service *application.CustomerService
}

func NewCustomerHandler(service *application.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/customers")
	g.POST("", h.CreateCustomer)
	g.GET("", h.ListCustomers)
	g.GET(":id", h.GetCustomer)
	g.PUT(":id", h.UpdateCustomer)
	g.DELETE(":id", h.DeleteCustomer)
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateCustomer(&customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) ListCustomers(c echo.Context) error {
	list, err := h.service.ListCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := h.service.GetCustomerByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	customer.ID = uint(id)
	err := h.service.UpdateCustomer(&customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteCustomer(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
