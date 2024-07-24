//+build faker

package pingdom

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	"github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/monitorables/pingdom/api"
	pingdomDelivery "github.com/Vaelatern/monitoror/monitorables/pingdom/api/delivery/http"
	pingdomModels "github.com/Vaelatern/monitoror/monitorables/pingdom/api/models"
	pingdomUsecase "github.com/Vaelatern/monitoror/monitorables/pingdom/api/usecase"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	monitorable.DefaultMonitorableFaker

	store *store.Store

	// Config tile settings
	checkTileEnabler            registry.TileEnabler
	transactionCheckTileEnabler registry.TileEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store

	// Register Monitorable Tile in config manager
	m.checkTileEnabler = store.Registry.RegisterTile(api.PingdomCheckTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.transactionCheckTileEnabler = store.Registry.RegisterTile(api.PingdomTransactionCheckTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string { return "Pingdom" }

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	usecase := pingdomUsecase.NewPingdomUsecase()
	delivery := pingdomDelivery.NewPingdomDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/pingdom", variantName)
	checkRoute := routeGroup.GET("/check", delivery.GetCheck)
	transactionCheckRoute := routeGroup.GET("/transaction-check", delivery.GetTransactionCheck)

	// EnableTile data for config hydration
	m.checkTileEnabler.Enable(variantName, &pingdomModels.CheckParams{}, checkRoute.Path)
	m.transactionCheckTileEnabler.Enable(variantName, &pingdomModels.TransactionCheckParams{}, transactionCheckRoute.Path)
}
