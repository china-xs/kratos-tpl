package provider

import (
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/account"
	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/data/good"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	data.NewData,
	account.NewAccountRepo,
	good.NewGoodRepo,
)
