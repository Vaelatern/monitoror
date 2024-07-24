//go:generate mockery -name Usecase

package api

import (
	uiConfigModels "github.com/Vaelatern/monitoror/api/config/models"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/jenkins/api/models"
)

const (
	JenkinsBuildTileType coreModels.TileType = "JENKINS-BUILD"
)

type (
	Usecase interface {
		Build(params *models.BuildParams) (*coreModels.Tile, error)
		BuildGenerator(params interface{}) ([]uiConfigModels.GeneratedTile, error)
	}
)
