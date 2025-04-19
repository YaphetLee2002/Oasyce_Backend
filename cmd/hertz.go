package cmd

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"go.uber.org/fx"

	"Oasyce-backend/api/kitex/kitex_gen/oasyce/combineservice"
	"Oasyce-backend/util/apigateway"
	"Oasyce-backend/util/conf"
	hertzutil "Oasyce-backend/util/hertz"
)

func hertzModule(target buildTarget) fx.Option {
	opts := []fx.Option{
		fx.Provide(newHertzServer),
		fx.Provide(provideServerConfig),
	}
	if target == BuildTargetAPI {
		opts = append(opts, fx.Invoke(startHertzServer))
	}
	return fx.Options(opts...)
}

// ServerConfig 定义配置参数
type ServerConfig struct {
	Addr string
	Port int
}

// 提供配置
func provideServerConfig() ServerConfig {
	return ServerConfig{
		Addr: "0.0.0.0",
		Port: 8080,
	}
}

// 创建Hertz服务器
func newHertzServer(combineClient combineservice.Client, hertzParam hertzutil.Param) *server.Hertz {
	var h *server.Hertz
	opt := config.Option{F: func(o *config.Options) {
		o.Addr = conf.GlobalConfig().API.Hertz.Addr
		o.MaxRequestBodySize = 20 << 20 // Set a higher request body size limit for multipart forms, 20MB
		o.UseRawPath = true
		o.StreamRequestBody = true
	}}
	switch conf.GlobalConfig().API.ENV {
	case "local", "dev", "prod":
		h = server.Default(opt)
	default:
		panic("Invalid Env config")
	}

	h.Any("/", apigateway.Handler(combineClient))
	group := h.Group("/")
	for _, registerer := range hertzParam.Registerers {
		registerer.Register(group)
	}
	return h
}

func startHertzServer(lc fx.Lifecycle, hertz *server.Hertz) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error, 1)
			go func() {
				errChan <- hertz.Run()
			}()
			select {
			case err := <-errChan:
				return err
			case <-time.After(5 * time.Second):
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			return hertz.Shutdown(ctx)
		},
	})
	return nil
}
