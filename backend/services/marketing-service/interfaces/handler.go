package interfaces

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"../application"
	"../domain"
)

type MarketingHandler struct {
	service *application.MarketingService
}

func NewMarketingHandler(service *application.MarketingService) *MarketingHandler {
	return &MarketingHandler{service: service}
}

func (h *MarketingHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/campaigns")
	g.POST("", h.CreateCampaign)
	g.GET("", h.ListCampaigns)
	g.GET(":id", h.GetCampaign)
	g.PUT(":id", h.UpdateCampaign)
	g.DELETE(":id", h.DeleteCampaign)
}

func (h *MarketingHandler) CreateCampaign(c echo.Context) error {
	var campaign domain.Campaign
	if err := c.Bind(&campaign); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateCampaign(&campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, campaign)
}

func (h *MarketingHandler) ListCampaigns(c echo.Context) error {
	list, err := h.service.ListCampaigns()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h *MarketingHandler) GetCampaign(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := h.service.GetCampaignByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, campaign)
}

func (h *MarketingHandler) UpdateCampaign(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var campaign domain.Campaign
	if err := c.Bind(&campaign); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	campaign.ID = uint(id)
	err := h.service.UpdateCampaign(&campaign)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, campaign)
}

func (h *MarketingHandler) DeleteCampaign(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteCampaign(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
