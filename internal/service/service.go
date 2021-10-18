package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"git.dev.enbrands.com/scrm/bed/scrm/api/demo"
)

type ServHTTPHandler func(srv *http.Server)
type ServGRPCHandler func(srv *grpc.Server)
type AppService struct {
	HS []ServHTTPHandler
	GS []ServGRPCHandler
}

//NewAppService 依赖参数&grpc && http 注册统一处理
func NewAppService(
	service *DemoService, //每新增一个proto service 需要在这里添加对应服务

) *AppService {
	app := AppService{
		//http 接口注册处
		HS: []ServHTTPHandler{},
		//gpc 接口注册处s
		GS: []ServGRPCHandler{
			func(srv *grpc.Server) { demo.RegisterDemoServer(srv, service) },
		},
	}
	return &app
}
