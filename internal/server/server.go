package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var Set = wire.NewSet(NewHTTPServer, NewGRPCServer)
