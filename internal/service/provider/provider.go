package provider

import (
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/service"
	accountv1svc "git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/service/v1/account"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var Set = wire.NewSet(
	service.NewAppService,
	accountv1svc.NewAccountService,
)
