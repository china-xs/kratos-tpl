package provider

import (
	"github.com/china-xs/kratos-tpl/internal/service"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var Set = wire.NewSet(
	service.NewAppService,
	service.NewDemoService,
)
