package main

import (
	"flag"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"

	registry2 "git.dev.enbrands.com/scrm/bed/scrm/pkg/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/common/constant"

	"git.dev.enbrands.com/scrm/bed/scrm/pkg/logger"
	"github.com/china-xs/kratos-tpl/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "微服务-单体tpl"
	// Version is the version of the compiled software.
	Version = "0.0.0.1"
	// flagconf is the config flag.
	flagconf string
	//线上根据 oa.
	LogPath = "../../../../runtime/logs/app.log" //项目初期，没有链路日志统一写到app.log
	id, _   = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()
	//zap 实现kratos logs Logger
	lg := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s][%s]", Name, Version, id)),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileRotationP(LogPath),
	)
	logger := log.With(
		lg,
		"caller", log.Caller(4),
		"trace_id", log.TraceID(),
		"span_id", log.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	/*
		//trace jaeger 链路追踪
		exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(bc.Trace.Endpoint)))
		if err != nil {
			panic(err)
		}
	*/
	tp := traceSDK.NewTracerProvider(
		//traceSDK.WithBatcher(exp),
		traceSDK.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
		)),
	)
	//Tracer 全局注册
	otel.SetTracerProvider(tp)

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	//服务发现参数
	rc := &registry2.RegistryConf{
		Sc: []constant.ServerConfig{
			*constant.NewServerConfig(
				bc.Registry.Nacos.Address,
				bc.Registry.Nacos.Port,
				constant.WithScheme(bc.Registry.Nacos.Scheme),
				constant.WithContextPath(bc.Registry.Nacos.Path),
			),
		},
		Cc: &constant.ClientConfig{
			NamespaceId:         bc.Registry.Nacos.Config.NamespaceId, //namespace id
			TimeoutMs:           bc.Registry.Nacos.Config.TimeoutMs,
			NotLoadCacheAtStart: bc.Registry.Nacos.Config.NotLoadCacheAtStart,
			LogDir:              bc.Registry.Nacos.Config.LogDir,
			CacheDir:            bc.Registry.Nacos.Config.CacheDir,
			RotateTime:          bc.Registry.Nacos.Config.RotateTime,
			MaxAge:              bc.Registry.Nacos.Config.MaxAge,
			LogLevel:            bc.Registry.Nacos.Config.LogLevel,
		},
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
