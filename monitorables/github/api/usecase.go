//go:generate mockery -name Usecase

package api

import (
	uiConfigModels "github.com/Vaelatern/monitoror/api/config/models"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/github/api/models"
)

const (
	GithubCountTileType       coreModels.TileType = "GITHUB-COUNT"
	GithubChecksTileType      coreModels.TileType = "GITHUB-CHECKS"
	GithubPullRequestTileType coreModels.TileType = "GITHUB-PULLREQUEST"
)

type (
	Usecase interface {
		Count(params *models.CountParams) (*coreModels.Tile, error)
		Checks(params *models.ChecksParams) (*coreModels.Tile, error)
		PullRequest(params *models.PullRequestParams) (*coreModels.Tile, error)

		PullRequestsGenerator(params interface{}) ([]uiConfigModels.GeneratedTile, error)
	}
)
