package interfaces

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"../application"
	"../domain"
)

type BranchHandler struct {
	service *application.BranchService
}

func NewBranchHandler(service *application.BranchService) *BranchHandler {
	return &BranchHandler{service: service}
}

func (h *BranchHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/branches")
	g.POST("", h.CreateBranch)
	g.GET("", h.ListBranches)
	g.GET(":id", h.GetBranch)
	g.PUT(":id", h.UpdateBranch)
	g.DELETE(":id", h.DeleteBranch)
}

func (h *BranchHandler) CreateBranch(c echo.Context) error {
	var branch domain.Branch
	if err := c.Bind(&branch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.service.CreateBranch(&branch)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, branch)
}

func (h *BranchHandler) ListBranches(c echo.Context) error {
	list, err := h.service.ListBranches()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, list)
}

func (h *BranchHandler) GetBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	branch, err := h.service.GetBranchByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, branch)
}

func (h *BranchHandler) UpdateBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var branch domain.Branch
	if err := c.Bind(&branch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	branch.ID = uint(id)
	err := h.service.UpdateBranch(&branch)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, branch)
}

func (h *BranchHandler) DeleteBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteBranch(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
