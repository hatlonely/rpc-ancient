package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hatlonely/go-kit/bind"
	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"

	"github.com/hatlonely/rpc-ancient/api/gen/go/api"
	"github.com/hatlonely/rpc-ancient/internal/service"
)

var Version string

type Options struct {
	flag.Options

	Http struct {
		Port int
	}
	Grpc struct {
		Port int
	}

	ExitTimeout time.Duration `dft:"10s"`

	Mysql         wrap.GORMDBWrapperOptions
	Elasticsearch cli.ElasticSearchOptions
	Service       service.Options
	Jaeger        jaegercfg.Configuration

	Logger struct {
		Info logger.Options
		Grpc logger.Options
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var options Options
	Must(flag.Struct(&options, refx.WithCamelName()))
	Must(flag.Parse(flag.WithJsonVal()))
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}

	if options.ConfigPath == "" {
		options.ConfigPath = "config/app.json"
	}
	cfg, err := config.NewConfigWithSimpleFile(options.ConfigPath)
	Must(err)
	Must(cfg.Watch())
	defer cfg.Stop()

	Must(bind.Bind(&options, []bind.Getter{
		flag.Instance(), bind.NewEnvGetter(bind.WithEnvPrefix("ANCIENT")), cfg,
	}, refx.WithCamelName()))

	grpcLog, err := logger.NewLoggerWithOptions(&options.Logger.Grpc, refx.WithCamelName())
	Must(err)
	infoLog, err := logger.NewLoggerWithOptions(&options.Logger.Info, refx.WithCamelName())
	Must(err)

	tracer, closer, err := options.Jaeger.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	Must(err)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	mysqlCli, err := wrap.NewGORMDBWrapperWithConfig(cfg.Sub("mysql"), refx.WithCamelName())
	Must(err)
	esCli, err := cli.NewElasticSearchWithOptions(&options.Elasticsearch)

	svc, err := service.NewAncientServiceWithOptions(mysqlCli, esCli, &options.Service)
	Must(err)

	grpcServer := grpc.NewServer(
		rpcx.GRPCUnaryInterceptor(
			grpcLog,
			rpcx.WithGRPCUnaryInterceptorDefaultValidator(),
			rpcx.WithGRPCUnaryInterceptorEnableTracing(),
		),
	)
	api.RegisterAncientServiceServer(grpcServer, svc)

	go func() {
		address, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port))
		Must(err)
		Must(grpcServer.Serve(address))
	}()

	mux := runtime.NewServeMux(
		rpcx.MuxWithMetadata(),
		rpcx.MuxWithIncomingHeaderMatcher(),
		rpcx.MuxWithOutgoingHeaderMatcher(),
		rpcx.MuxWithProtoErrorHandler(),
	)
	Must(api.RegisterAncientServiceHandlerFromEndpoint(
		context.Background(), mux, fmt.Sprintf("0.0.0.0:%v", options.Grpc.Port), []grpc.DialOption{
			grpc.WithInsecure(), grpc.WithUnaryInterceptor(
				grpc_opentracing.UnaryClientInterceptor(
					grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
				),
			),
		},
	))
	infoLog.Info(options)

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%v", options.Http.Port),
		Handler: rpcx.MetricWrapper(rpcx.TraceWrapper(mux)),
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			infoLog.Warnf("httpServer.ListenAndServe, err: [%v]", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	infoLog.Info("receive exit signal")
	ctx, cancel := context.WithTimeout(context.Background(), options.ExitTimeout)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		infoLog.Warnf("httServer.Shutdown failed, err: [%v]", err)
	}
	grpcServer.Stop()
	infoLog.Info("server exit properly")
}
