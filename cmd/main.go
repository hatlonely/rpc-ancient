package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hatlonely/go-kit/bind"
	"github.com/hatlonely/go-kit/config"
	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	_ "github.com/hatlonely/go-kit/logger/x"
	"github.com/hatlonely/go-kit/micro"
	microx "github.com/hatlonely/go-kit/micro/x"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/hatlonely/go-kit/wrap"

	"github.com/hatlonely/rpc-ancient/api/gen/go/api"
	"github.com/hatlonely/rpc-ancient/internal/service"
)

var Version string

type Options struct {
	flag.Options

	Service            service.Options
	GrpcGateway        rpcx.GrpcGatewayOptions
	Mysql              wrap.GORMDBWrapperOptions
	ESClient           wrap.ESClientWrapperOptions
	RateLimiter        microx.RedisRateLimiterOptions
	ParallelController microx.RedisTimedParallelControllerOptions

	Logger struct {
		Info logger.Options
		Grpc logger.Options
	}
}

func main() {
	var options Options
	refx.Must(flag.Struct(&options, refx.WithCamelName()))
	refx.Must(flag.Parse(flag.WithJsonVal()))
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
	refx.Must(err)
	refx.Must(cfg.Watch())
	defer cfg.Stop()

	refx.Must(bind.Bind(&options, []bind.Getter{
		flag.Instance(), bind.NewEnvGetter(bind.WithEnvPrefix("ANCIENT")), cfg,
	}, refx.WithCamelName()))

	grpcLog, err := logger.NewLoggerWithOptions(&options.Logger.Grpc, refx.WithCamelName())
	refx.Must(err)
	infoLog, err := logger.NewLoggerWithOptions(&options.Logger.Info, refx.WithCamelName())
	refx.Must(err)
	infoLog.With("options", options).Info("init config success")

	rateLimiter, err := microx.NewRedisRateLimiterWithConfig(cfg.Sub("rateLimiter"), refx.WithCamelName())
	refx.Must(err)
	micro.RegisterRateLimiter("RedisRateLimiterInstance", rateLimiter)
	parallelController, err := microx.NewRedisTimedParallelControllerWithOptions(&options.ParallelController)
	refx.Must(err)
	micro.RegisterParallelController("RedisTimedParallelControllerInstance", parallelController)

	mysqlCli, err := wrap.NewGORMDBWrapperWithConfig(cfg.Sub("mysql"), refx.WithCamelName())
	refx.Must(err)
	esCli, err := wrap.NewESClientWrapperWithOptions(&options.ESClient)

	svc, err := service.NewAncientServiceWithOptions(mysqlCli, esCli, &options.Service)
	refx.Must(err)

	grpcGateway, err := rpcx.NewGrpcGatewayWithOptions(&options.GrpcGateway)
	refx.Must(err)
	grpcGateway.SetLogger(infoLog, grpcLog)

	api.RegisterAncientServiceServer(grpcGateway.GRPCServer(), svc)
	refx.Must(grpcGateway.RegisterServiceHandlerFunc(api.RegisterAncientServiceHandlerFromEndpoint))
	grpcGateway.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	grpcGateway.Stop()
	infoLog.Info("server exit properly")
}
