package provider

import (
	"github.com/china-xs/kratos-tpl/internal/data"
	"github.com/china-xs/kratos-tpl/internal/data/account"
	"github.com/china-xs/kratos-tpl/internal/data/good"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	data.NewData,
	account.NewAccountRepo,
	good.NewGoodRepo,
)
