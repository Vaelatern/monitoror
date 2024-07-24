//go:generate mockery -name MonitorableRouter|MonitorableRouterGroup  -output ../mocks

package router

import (
	"fmt"

	coreModels "github.com/Vaelatern/monitoror/models"
	"github.com/Vaelatern/monitoror/service/middlewares"
	"github.com/Vaelatern/monitoror/service/options"

	"github.com/labstack/echo/v4"
)

type (
	MonitorableRouter interface {
		Group(path string, variantName coreModels.VariantName) MonitorableRouterGroup
	}
	MonitorableRouterGroup interface {
		GET(path string, handlerFunc echo.HandlerFunc, options ...options.RouterOption) *echo.Route
	}

	router struct {
		apiVersion      *echo.Group
		cacheMiddleware *middlewares.CacheMiddleware
	}

	group struct {
		router *router
		group  *echo.Group
	}
)

func NewMonitorableRouter(apiVersion *echo.Group, cacheMiddleware *middlewares.CacheMiddleware) MonitorableRouter {
	return &router{apiVersion: apiVersion, cacheMiddleware: cacheMiddleware}
}

func (r *router) Group(path string, variantName coreModels.VariantName) MonitorableRouterGroup {
	return &group{router: r, group: r.apiVersion.Group(fmt.Sprintf(`%s/%s`, path, variantName))}
}

func (g *group) GET(path string, handlerFunc echo.HandlerFunc, opts ...options.RouterOption) *echo.Route {
	routerSettings := options.ApplyOptions(opts...)

	handler := handlerFunc
	if !routerSettings.NoCache {
		if routerSettings.CustomCacheExpiration != nil {
			handler = g.router.cacheMiddleware.UpstreamCacheHandlerWithExpiration(*routerSettings.CustomCacheExpiration, handlerFunc)
		} else {
			handler = g.router.cacheMiddleware.UpstreamCacheHandler(handlerFunc)
		}
	}

	return g.group.GET(path, handler, routerSettings.Middlewares...)
}
