//+build faker

package models

import (
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/params"
	coreModels "github.com/Vaelatern/monitoror/models"
)

type (
	PingParams struct {
		params.Default

		Hostname string `json:"hostname" query:"hostname" validate:"required"`

		Status      coreModels.TileStatus `json:"status" query:"status"`
		ValueValues []string              `json:"valueValues" query:"valueValues"`
	}
)
