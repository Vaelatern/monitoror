package store

import (
	"github.com/jsdidierlaurent/echo-middleware/cache"

	coreConfig "github.com/Vaelatern/monitoror/config"
	"github.com/Vaelatern/monitoror/registry"
	"github.com/Vaelatern/monitoror/service/router"
)

type (
	// Store is used to share Data in every monitorable
	Store struct {
		// Global CoreConfig
		CoreConfig *coreConfig.CoreConfig

		// CacheStore for every memory persistent data
		CacheStore cache.Store

		// Registry used to register Tile for verify / hydrate
		Registry registry.Registry

		// MonitorableRouter helper wrapping echo Router monitorable
		MonitorableRouter router.MonitorableRouter
	}
)
