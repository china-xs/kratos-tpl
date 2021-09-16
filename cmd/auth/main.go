package main

import (
	"flag"
	"fmt"
	"git.dev.enbrands.com/scrm/bed/scrm/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"

	"git.dev.enbrands.com/scrm/bed/scrm/app/auth/internal/conf"
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
	Name = "scrm-auth-权限管理"
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

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
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
		"caller", log.DefaultCaller,
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

	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
