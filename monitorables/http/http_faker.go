//+build faker

package http

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/http/api"
	httpDelivery "github.com/Vaelatern/monitoror/monitorables/http/api/delivery/http"
	httpModels "github.com/Vaelatern/monitoror/monitorables/http/api/models"
	httpUsecase "github.com/Vaelatern/monitoror/monitorables/http/api/usecase"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	monitorable.DefaultMonitorableFaker

	store *store.Store

	// Config tile settings
	statusTileEnabler    registry.TileEnabler
	rawTileEnabler       registry.TileEnabler
	formattedTileEnabler registry.TileEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store

	// Register Monitorable Tile in config manager
	m.statusTileEnabler = store.Registry.RegisterTile(api.HTTPStatusTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.rawTileEnabler = store.Registry.RegisterTile(api.HTTPRawTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.formattedTileEnabler = store.Registry.RegisterTile(api.HTTPFormattedTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string { return "HTTP" }

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	usecase := httpUsecase.NewHTTPUsecase()
	delivery := httpDelivery.NewHTTPDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/http", variantName)
	routeStatus := routeGroup.GET("/status", delivery.GetHTTPStatus)
	routeRaw := routeGroup.GET("/raw", delivery.GetHTTPRaw)
	routeJSON := routeGroup.GET("/formatted", delivery.GetHTTPFormatted)

	// EnableTile data for config hydration
	m.statusTileEnabler.Enable(variantName, &httpModels.HTTPStatusParams{}, routeStatus.Path)
	m.rawTileEnabler.Enable(variantName, &httpModels.HTTPRawParams{}, routeRaw.Path)
	m.formattedTileEnabler.Enable(variantName, &httpModels.HTTPFormattedParams{}, routeJSON.Path)
}
