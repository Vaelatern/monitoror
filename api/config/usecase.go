//go:generate mockery -name Usecase

package config

import (
	"github.com/Vaelatern/monitoror/api/config/models"
)

type (
	Usecase interface {
		GetConfigList() []models.ConfigMetadata
		GetConfig(params *models.ConfigParams) *models.ConfigBag
		Verify(config *models.ConfigBag)
		Hydrate(config *models.ConfigBag)
	}
)
