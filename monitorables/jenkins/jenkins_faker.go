//+build faker

package jenkins

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/jenkins/api"
	jenkinsDelivery "github.com/Vaelatern/monitoror/monitorables/jenkins/api/delivery/http"
	jenkinsModels "github.com/Vaelatern/monitoror/monitorables/jenkins/api/models"
	jenkinsUsecase "github.com/Vaelatern/monitoror/monitorables/jenkins/api/usecase"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	monitorable.DefaultMonitorableFaker

	store *store.Store

	// Config tile settings
	buildTileEnabler registry.TileEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store

	// Register Monitorable Tile in config manager
	m.buildTileEnabler = store.Registry.RegisterTile(api.JenkinsBuildTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string { return "Jenkins" }

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	usecase := jenkinsUsecase.NewJenkinsUsecase()
	delivery := jenkinsDelivery.NewJenkinsDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/jenkins", variantName)
	route := routeGroup.GET("/build", delivery.GetBuild)

	// EnableTile data for config hydration
	m.buildTileEnabler.Enable(variantName, &jenkinsModels.BuildParams{}, route.Path)
}
