package interfaces

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"../application"
	"../domain"
)

type FinanceHandler struct {
	service *application.FinanceService
}

func NewFinanceHandler(service *application.FinanceService) *FinanceHandler {
	return &FinanceHandler{service: service}
}

func (h *FinanceHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/invoices")
	g.POST("", h.CreateInvoice)
	g.GET("", h.ListInvoices)
	g.GET(":id", h.GetInvoice)
	g.PUT(":id", h.UpdateInvoice)
	g.DELETE(":id", h.DeleteInvoice)
}

func (h *FinanceHandler) CreateInvoice(c echo.Context) error {
	var invoice domain.Invoice
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateInvoice(&invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, invoice)
}

func (h *FinanceHandler) ListInvoices(c echo.Context) error {
	list, err := h.service.ListInvoices()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h *FinanceHandler) GetInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.service.GetInvoiceByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, invoice)
}

func (h *FinanceHandler) UpdateInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var invoice domain.Invoice
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	invoice.ID = uint(id)
	err := h.service.UpdateInvoice(&invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoice)
}

func (h *FinanceHandler) DeleteInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteInvoice(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
