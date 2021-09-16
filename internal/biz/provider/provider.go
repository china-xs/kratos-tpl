package provider

import (
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/biz/account"
	"github.com/google/wire"
)

// ProviderSet is biz providers
// logic 新建方法需要在这里注册
var Set = wire.NewSet(
	account.NewAccountLogic,
)
