package models

import (
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable/params"
)

type (
	CheckGeneratorParams struct {
		params.Default

		Tags   string `json:"tags,omitempty" query:"tags"`
		SortBy string `json:"sortBy,omitempty" query:"sortBy" validate:"omitempty,oneof=name"`
	}
)
