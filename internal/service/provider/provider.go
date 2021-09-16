package provider

import (
	"github.com/china-xs/kratos-tpl/internal/service"
	accountv1svc "github.com/china-xs/kratos-tpl/internal/service/v1/account"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var Set = wire.NewSet(
	service.NewAppService,
	accountv1svc.NewAccountService,
)
