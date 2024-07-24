//+build faker

package models

import "github.com/Vaelatern/monitoror/internal/pkg/monitorable/params"

type (
	CountParams struct {
		params.Default

		Query string `json:"query" query:"query" validate:"required"`

		ValueValues []string `json:"valueValues" query:"valueValues"`
	}
)
