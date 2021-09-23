package provider

import (
	"github.com/china-xs/kratos-tpl/internal/biz/demologic"
	"github.com/google/wire"
)

// ProviderSet is biz providers
// logic 新建方法需要在这里注册
var Set = wire.NewSet(
	demologic.NewDemoLogic, //案例逻辑层注册
)
