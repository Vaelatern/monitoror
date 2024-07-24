//go:generate mockery -name Usecase

package api

import (
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/travisci/api/models"
)

const (
	TravisCIBuildTileType coreModels.TileType = "TRAVISCI-BUILD"
)

type (
	Usecase interface {
		Build(params *models.BuildParams) (*coreModels.Tile, error)
	}
)
