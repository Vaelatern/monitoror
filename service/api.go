package service

import (
	configDelivery "github.com/Vaelatern/monitoror/api/config/delivery/http"
	configRepository "github.com/Vaelatern/monitoror/api/config/repository"
	configUsecase "github.com/Vaelatern/monitoror/api/config/usecase"
	"github.com/Vaelatern/monitoror/api/info"
	"github.com/Vaelatern/monitoror/monitorables"
	"github.com/Vaelatern/monitoror/service/router"

	"github.com/jsdidierlaurent/echo-middleware/cache"
)

func InitApis(s *Server) {
	// API group
	apiGroup := s.Group("/api/v1")

	// ------------- INFO ------------- //
	infoDelivery := info.NewHTTPInfoDelivery()
	apiGroup.GET("/info", s.CacheMiddleware.UpstreamCacheHandlerWithExpiration(cache.NEVER, infoDelivery.GetInfo))

	// ------------- CONFIG ------------- //
	confRepository := configRepository.NewConfigRepository()
	confUsecase := configUsecase.NewConfigUsecase(confRepository, s.store)
	confDelivery := configDelivery.NewConfigDelivery(confUsecase)
	apiGroup.GET("/configs", s.CacheMiddleware.UpstreamCacheHandler(confDelivery.GetConfigList))
	apiGroup.GET("/configs/:config", s.CacheMiddleware.UpstreamCacheHandler(confDelivery.GetConfig))

	// ---------------------------------- //
	s.store.MonitorableRouter = router.NewMonitorableRouter(apiGroup, s.CacheMiddleware)
	// ---------------------------------- //

	// ------------- MONITORABLES ------------- //
	monitorables.RegisterMonitorables(s.store)

	for _, mm := range s.store.Registry.GetMonitorables() {
		for _, vm := range mm.VariantsMetadata {
			if vm.Enabled {
				mm.Monitorable.Enable(vm.VariantName)
			}
		}
	}
}
