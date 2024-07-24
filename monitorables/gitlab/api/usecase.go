//go:generate mockery -name Usecase

package api

import (
	uiConfigModels "github.com/Vaelatern/monitoror/api/config/models"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/gitlab/api/models"
)

const (
	GitlabCountIssuesTileType  coreModels.TileType = "GITLAB-COUNT-ISSUES"
	GitlabPipelineTileType     coreModels.TileType = "GITLAB-PIPELINE"
	GitlabMergeRequestTileType coreModels.TileType = "GITLAB-MERGEREQUEST"
)

type (
	Usecase interface {
		CountIssues(params *models.IssuesParams) (*coreModels.Tile, error)
		Pipeline(params *models.PipelineParams) (*coreModels.Tile, error)
		MergeRequest(params *models.MergeRequestParams) (*coreModels.Tile, error)

		MergeRequestsGenerator(params interface{}) ([]uiConfigModels.GeneratedTile, error)
	}
)
