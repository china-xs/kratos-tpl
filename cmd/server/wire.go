// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	// provider
	"git.dev.enbrands.com/scrm/bed/scrm/pkg/registry"
	pd_biz "github.com/china-xs/kratos-tpl/internal/biz/provider"
	"github.com/china-xs/kratos-tpl/internal/conf"
	pd_data "github.com/china-xs/kratos-tpl/internal/data/provider"
	pd_srv "github.com/china-xs/kratos-tpl/internal/server"
	pd_svc "github.com/china-xs/kratos-tpl/internal/service/provider"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *registry.RegistryConf, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		pd_srv.Set,
		pd_data.Set,
		pd_biz.Set,
		pd_svc.Set,
		newApp))
}
