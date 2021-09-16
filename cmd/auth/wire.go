// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	// provider
	pd_biz "git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/biz/provider"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/conf"
	pd_data "git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/provider"
	pd_srv "git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/server"
	pd_svc "git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/service/provider"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		pd_srv.Set,
		pd_data.Set,
		pd_biz.Set,
		pd_svc.Set,
		newApp))
}
