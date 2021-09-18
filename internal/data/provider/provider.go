package provider

import (
	"github.com/china-xs/kratos-tpl/internal/data"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	data.NewData,
)
