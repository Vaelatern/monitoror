//+build faker

package ping

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/ping/api"
	pingDelivery "github.com/Vaelatern/monitoror/monitorables/ping/api/delivery/http"
	pingModels "github.com/Vaelatern/monitoror/monitorables/ping/api/models"
	pingUsecase "github.com/Vaelatern/monitoror/monitorables/ping/api/usecase"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	monitorable.DefaultMonitorableFaker

	store *store.Store

	// Config tile settings
	pingTileEnabler registry.TileEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store

	// Register Monitorable Tile in config manager
	m.pingTileEnabler = store.Registry.RegisterTile(api.PingTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string { return "Ping" }

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	usecase := pingUsecase.NewPingUsecase()
	delivery := pingDelivery.NewPingDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/ping", variantName)
	route := routeGroup.GET("/ping", delivery.GetPing)

	// EnableTile data for config hydration
	m.pingTileEnabler.Enable(variantName, &pingModels.PingParams{}, route.Path)
}
