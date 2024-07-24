//go:generate mockery -name Usecase

package api

import (
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/http/api/models"
)

const (
	HTTPStatusTileType    coreModels.TileType = "HTTP-STATUS"
	HTTPRawTileType       coreModels.TileType = "HTTP-RAW"
	HTTPFormattedTileType coreModels.TileType = "HTTP-FORMATTED"
)

type (
	Usecase interface {
		HTTPStatus(params *models.HTTPStatusParams) (*coreModels.Tile, error)
		HTTPRaw(params *models.HTTPRawParams) (*coreModels.Tile, error)
		HTTPFormatted(params *models.HTTPFormattedParams) (*coreModels.Tile, error)
	}
)
