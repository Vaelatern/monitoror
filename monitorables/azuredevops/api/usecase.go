//go:generate mockery -name Usecase

package api

import (
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/azuredevops/api/models"
)

const (
	AzureDevOpsBuildTileType   coreModels.TileType = "AZUREDEVOPS-BUILD"
	AzureDevOpsReleaseTileType coreModels.TileType = "AZUREDEVOPS-RELEASE"
)

type (
	Usecase interface {
		Build(params *models.BuildParams) (*coreModels.Tile, error)
		Release(params *models.ReleaseParams) (*coreModels.Tile, error)
	}
)
