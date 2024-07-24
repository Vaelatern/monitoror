package info

import (
	"net/http"

	"github.com/Vaelatern/monitoror/cli/version"
	"github.com/Vaelatern/monitoror/models"

	"github.com/labstack/echo/v4"
)

type HTTPInfoDelivery struct {
}

func NewHTTPInfoDelivery() *HTTPInfoDelivery {
	return &HTTPInfoDelivery{}
}

func (h *HTTPInfoDelivery) GetInfo(c echo.Context) error {
	response := models.NewInfoResponse(version.Version, version.GitCommit, version.BuildTime, version.BuildTags)
	return c.JSON(http.StatusOK, response)
}
