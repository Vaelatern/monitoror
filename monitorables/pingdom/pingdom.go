//+build !faker

package pingdom

import (
	"github.com/Vaelatern/monitoror/api/config/versions"
	pkgMonitorable "github.com/Vaelatern/monitoror/internal/pkg/monitorable"
	"github.com/Vaelatern/monitoror/registry"

	coreModels "github.com/Vaelatern/monitoror/models"

	"github.com/Vaelatern/monitoror/monitorables/pingdom/api"
	pingdomDelivery "github.com/Vaelatern/monitoror/monitorables/pingdom/api/delivery/http"
	pingdomModels "github.com/Vaelatern/monitoror/monitorables/pingdom/api/models"
	pingdomRepository "github.com/Vaelatern/monitoror/monitorables/pingdom/api/repository"
	pingdomUsecase "github.com/Vaelatern/monitoror/monitorables/pingdom/api/usecase"
	pingdomConfig "github.com/Vaelatern/monitoror/monitorables/pingdom/config"
	"github.com/Vaelatern/monitoror/store"
)

type Monitorable struct {
	store *store.Store

	config map[coreModels.VariantName]*pingdomConfig.Pingdom

	// Config tile settings
	checkTileEnabler                 registry.TileEnabler
	transactionCheckTileEnabler      registry.TileEnabler
	checkGeneratorEnabler            registry.GeneratorEnabler
	transactionCheckGeneratorEnabler registry.GeneratorEnabler
}

func NewMonitorable(store *store.Store) *Monitorable {
	m := &Monitorable{}
	m.store = store
	m.config = make(map[coreModels.VariantName]*pingdomConfig.Pingdom)

	// Load core config from env
	pkgMonitorable.LoadConfig(&m.config, pingdomConfig.Default)

	// Register Monitorable Tile in config manager
	m.checkTileEnabler = store.Registry.RegisterTile(api.PingdomCheckTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.transactionCheckTileEnabler = store.Registry.RegisterTile(api.PingdomTransactionCheckTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.checkGeneratorEnabler = store.Registry.RegisterGenerator(api.PingdomCheckTileType, versions.MinimalVersion, m.GetVariantsNames())
	m.transactionCheckGeneratorEnabler = store.Registry.RegisterGenerator(api.PingdomTransactionCheckTileType, versions.MinimalVersion, m.GetVariantsNames())

	return m
}

func (m *Monitorable) GetDisplayName() string {
	return "Pingdom"
}

func (m *Monitorable) GetVariantsNames() []coreModels.VariantName {
	return pkgMonitorable.GetVariantsNames(m.config)
}

func (m *Monitorable) Validate(variantName coreModels.VariantName) (bool, []error) {
	conf := m.config[variantName]

	// No configuration set
	if conf.URL == pingdomConfig.Default.URL && conf.Token == "" {
		return false, nil
	}

	// Validate Config
	if errors := pkgMonitorable.ValidateConfig(conf, variantName); errors != nil {
		return false, errors
	}

	return true, nil
}

func (m *Monitorable) Enable(variantName coreModels.VariantName) {
	conf := m.config[variantName]

	repository := pingdomRepository.NewPingdomRepository(conf)
	usecase := pingdomUsecase.NewPingdomUsecase(repository, m.store.CacheStore, conf.CacheExpiration)
	delivery := pingdomDelivery.NewPingdomDelivery(usecase)

	// EnableTile route to echo
	routeGroup := m.store.MonitorableRouter.Group("/pingdom", variantName)
	checkRoute := routeGroup.GET("/check", delivery.GetCheck)
	transactionCheckRoute := routeGroup.GET("/transaction-check", delivery.GetTransactionCheck)

	// EnableTile data for config hydration
	m.checkTileEnabler.Enable(variantName, &pingdomModels.CheckParams{}, checkRoute.Path)
	m.transactionCheckTileEnabler.Enable(variantName, &pingdomModels.CheckParams{}, transactionCheckRoute.Path)
	m.checkGeneratorEnabler.Enable(variantName, &pingdomModels.CheckGeneratorParams{}, usecase.CheckGenerator)
	m.transactionCheckGeneratorEnabler.Enable(variantName, &pingdomModels.TransactionCheckGeneratorParams{}, usecase.TransactionCheckGenerator)
}
