//go:generate mockery -name Usecase

package api

import (
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/ping/api/models"
)

const (
	PingTileType coreModels.TileType = "PING"
)

type (
	Usecase interface {
		Ping(params *models.PingParams) (*coreModels.Tile, error)
	}
)
