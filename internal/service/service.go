package service

import (
	"github.com/china-xs/kratos-tpl/api/v1/account"
	svcv1a "github.com/china-xs/kratos-tpl/internal/service/v1/account"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAppService,
	svcv1a.NewAccountService,
)

type ServHTTPHandler func(serv *http.Server)
type ServGRPCHandler func(serv *grpc.Server)
type AppService struct {
	HS []ServHTTPHandler
	GS []ServGRPCHandler
}

//NewAppService 依赖参数&grpc && http 注册统一处理
func NewAppService(
	accSrvV1 *svcv1a.AccountService,
) *AppService {
	app := AppService{
		//http 接口注册处
		HS: []ServHTTPHandler{
			func(srv *http.Server) { account.RegisterAccountHTTPServer(srv, accSrvV1) },
		},
		//gpc 接口注册处s
		GS: []ServGRPCHandler{
			func(srv *grpc.Server) { account.RegisterAccountServer(srv, accSrvV1) },
		},
	}
	return &app
}
