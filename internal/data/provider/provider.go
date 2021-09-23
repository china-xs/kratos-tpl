package provider

import (
	"git.dev.enbrands.com/scrm/bed/scrm/pkg/registry"
	"github.com/china-xs/kratos-tpl/internal/data"
	"github.com/china-xs/kratos-tpl/internal/data/demo"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	registry.NewNacosRegistrar, //服务注册
	registry.NewNacosDiscovery, //服务发现
	data.NewData,
	demo.NewDemoRepo,
)
