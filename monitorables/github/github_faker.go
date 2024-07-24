//+build faker

package github

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/github/api"
	githubDelivery "github.com/Vaelatern/monitoror/monitorables/github/api/delivery/http"
	githubModels "github.com/Vaelatern/monitoror/monitorables/github/api/models"
	githubUsecase "github.com/Vaelatern/monitoror/monitorables/github/api/usecase"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	monitorable.DefaultMonitorableFaker

	store *store.Store

	// Config tile settings
	countTileEnabler       registry.TileEnabler
	checksTileEnabler      registry.TileEnabler
	pullRequestTileEnabler registry.TileEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store

	// Register Monitorable Tile in config manager
	m.countTileEnabler = store.Registry.RegisterTile(api.GithubCountTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.checksTileEnabler = store.Registry.RegisterTile(api.GithubChecksTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.pullRequestTileEnabler = store.Registry.RegisterTile(api.GithubPullRequestTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string {
	return "GitHub"
}

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	usecase := githubUsecase.NewGithubUsecase()
	delivery := githubDelivery.NewGithubDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/github", variantName)
	routeCount := routeGroup.GET("/count", delivery.GetCount)
	routeChecks := routeGroup.GET("/checks", delivery.GetChecks)
	routePullRequest := routeGroup.GET("/pullrequest", delivery.GetPullRequest)

	// EnableTile data for config hydration
	m.countTileEnabler.Enable(variantName, &githubModels.CountParams{}, routeCount.Path)
	m.checksTileEnabler.Enable(variantName, &githubModels.ChecksParams{}, routeChecks.Path)
	m.pullRequestTileEnabler.Enable(variantName, &githubModels.PullRequestParams{}, routePullRequest.Path)
}
