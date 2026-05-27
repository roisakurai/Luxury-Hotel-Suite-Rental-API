package handler

import (
	"net/http"
	"p2-ip-hotel-rental/models"
	"p2-ip-hotel-rental/service"

	"github.com/labstack/echo/v4"
)

type SuiteHandler struct {
	suiteService service.SuiteService
}

func NewSuiteHandler(suiteService service.SuiteService) *SuiteHandler {
	return &SuiteHandler{suiteService}
}

// GetSuites godoc
// @Summary Get list of suites
// @Description Retrieve all available hotel suites
// @Tags Suites
// @Produce json
// @Success 200 {object} models.SuitesResponse
// @Failure 500 {object} map[string]string
// @Router /suites [get]
func (h *SuiteHandler) GetSuites(c echo.Context) error {

	suites, err := h.suiteService.GetAllSuites()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	var result []models.SuiteResponse
	for _, s := range suites {
		result = append(result, models.ToSuiteResponse(s))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    result,
	})
}
