package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/your-org/analytics-service/application"
	"github.com/your-org/analytics-service/domain"
)

type AnalyticsHandler struct {
	service *application.AnalyticsService
}

func NewAnalyticsHandler(service *application.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{service: service}
}

func (h *AnalyticsHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/metrics")
	g.POST("", h.CreateMetric)
	g.GET("", h.ListMetrics)
}

func (h *AnalyticsHandler) CreateMetric(c echo.Context) error {
	var metric domain.Metric
	if err := c.Bind(&metric); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateMetric(&metric)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, metric)
}

func (h *AnalyticsHandler) ListMetrics(c echo.Context) error {
	list, err := h.service.ListMetrics()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}
