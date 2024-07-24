//+build faker

package models

import (
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/params"
	coreModels "github.com/Vaelatern/monitoror/models"
)

type (
	TransactionCheckParams struct {
		params.Default

		ID *int `json:"id" query:"id" validate:"required"`

		Status coreModels.TileStatus `json:"status" query:"status"`
	}
)
