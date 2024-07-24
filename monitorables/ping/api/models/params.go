//+build !faker

package models

import (
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/params"
)

type (
	PingParams struct {
		params.Default

		Hostname string `json:"hostname" query:"hostname" validate:"required"`
	}
)
