package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAppService,
)

type ServHTTPHandler func(serv *http.Server)
type ServGRPCHandler func(serv *grpc.Server)
type AppService struct {
	HS []ServHTTPHandler
	GS []ServGRPCHandler
}

//NewAppService 依赖参数&grpc && http 注册统一处理
func NewAppService() *AppService {
	app := AppService{
		//http 接口注册处
		HS: []ServHTTPHandler{},
		//gpc 接口注册处s
		GS: []ServGRPCHandler{},
	}
	return &app
}
