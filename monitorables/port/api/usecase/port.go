//+build !faker

package usecase

import (
	"fmt"

	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/port/api"
	"github.com/Vaelatern/monitoror/monitorables/port/api/models"
)

type (
	portUsecase struct {
		repository api.Repository
	}
)

func NewPortUsecase(repository api.Repository) api.Usecase {
	return &portUsecase{repository}
}

func (pu *portUsecase) Port(params *models.PortParams) (tile *coreModels.Tile, err error) {
	tile = coreModels.NewTile(api.PortTileType)
	tile.Label = fmt.Sprintf("%s:%d", params.Hostname, params.Port)

	err = pu.repository.OpenSocket(params.Hostname, params.Port)
	if err == nil {
		tile.Status = coreModels.SuccessStatus
	} else {
		tile.Status = coreModels.FailedStatus
		err = nil
	}

	return
}
