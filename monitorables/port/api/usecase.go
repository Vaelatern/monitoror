//go:generate mockery -name Usecase

package api

import (
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/port/api/models"
)

const (
	PortTileType coreModels.TileType = "PORT"
)

type (
	Usecase interface {
		Port(params *models.PortParams) (*coreModels.Tile, error)
	}
)
