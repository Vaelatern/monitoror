package http

import (
	"net/http"

	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/delivery"
	"github.com/Vaelatern/monitoror/monitorables/azuredevops/api"
	"github.com/Vaelatern/monitoror/monitorables/azuredevops/api/models"

	"github.com/labstack/echo/v4"
)

type AzureDevOpsDelivery struct {
	azureDevOpsUsecase api.Usecase
}

func NewAzureDevOpsDelivery(p api.Usecase) *AzureDevOpsDelivery {
	return &AzureDevOpsDelivery{p}
}

func (h *AzureDevOpsDelivery) GetBuild(c echo.Context) error {
	// Bind / check Params
	params := &models.BuildParams{}
	if err := delivery.BindAndValidateParams(c, params); err != nil {
		return err
	}

	tile, err := h.azureDevOpsUsecase.Build(params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tile)
}

func (h *AzureDevOpsDelivery) GetRelease(c echo.Context) error {
	// Bind / check Params
	params := &models.ReleaseParams{}
	if err := delivery.BindAndValidateParams(c, params); err != nil {
		return err
	}

	tile, err := h.azureDevOpsUsecase.Release(params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tile)
}
