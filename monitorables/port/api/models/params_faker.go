//+build faker

package models

import (
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/params"
	coreModels "github.com/Vaelatern/monitoror/models"
)

type (
	PortParams struct {
		params.Default

		Hostname string `json:"hostname" query:"hostname"`
		Port     int    `json:"port" query:"port"`

		Status coreModels.TileStatus `json:"status" query:"status"`
	}
)
