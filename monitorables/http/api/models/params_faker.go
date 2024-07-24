//+build faker

package models

import coreModels "github.com/Vaelatern/monitoror/models"

type (
	FakerParamsProvider interface {
		GetStatus() coreModels.TileStatus
		GetMessage() string
		GetValueValues() []string
		GetValueUnit() coreModels.TileValuesUnit
	}
)
