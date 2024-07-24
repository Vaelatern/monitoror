//+build faker

package azuredevops

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/azuredevops/api"
	azuredevopsDelivery "github.com/Vaelatern/monitoror/monitorables/azuredevops/api/delivery/http"
	azuredevopsModels "github.com/Vaelatern/monitoror/monitorables/azuredevops/api/models"
	azuredevopsUsecase "github.com/Vaelatern/monitoror/monitorables/azuredevops/api/usecase"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	monitorable.DefaultMonitorableFaker
	store *store.Store

	// Config tile settings
	buildTileEnabler   registry.TileEnabler
	releaseTileEnabler registry.TileEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store

	// Register Monitorable Tile in config manager
	m.buildTileEnabler = store.Registry.RegisterTile(api.AzureDevOpsBuildTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.releaseTileEnabler = store.Registry.RegisterTile(api.AzureDevOpsReleaseTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string {
	return "Azure DevOps"
}

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	usecase := azuredevopsUsecase.NewAzureDevOpsUsecase()
	delivery := azuredevopsDelivery.NewAzureDevOpsDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/azuredevops", variantName)
	routeBuild := routeGroup.GET("/build", delivery.GetBuild)
	routeRelease := routeGroup.GET("/release", delivery.GetRelease)

	// EnableTile data for config hydration
	m.buildTileEnabler.Enable(variantName, &azuredevopsModels.BuildParams{}, routeBuild.Path)
	m.releaseTileEnabler.Enable(variantName, &azuredevopsModels.ReleaseParams{}, routeRelease.Path)
}
