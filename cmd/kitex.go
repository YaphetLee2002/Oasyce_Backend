package cmd

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/endpoint"
	kitexretry "github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"

	"Oasyce-backend/api/kitex/kitex_gen/oasyce"
	"Oasyce-backend/api/kitex/kitex_gen/oasyce/combineservice"
	"Oasyce-backend/util/conf"
	"Oasyce-backend/util/di"
	"Oasyce-backend/util/timeout"

	"go.uber.org/fx"
)

func kitexModule(target buildTarget) fx.Option {
	var opts []fx.Option
	if target == BuildTargetAPI {
		opts = append(opts, fx.Provide(newKitexClient))
	} else if target == BuildTargetRPC {
		opts = append(opts,
			fx.Provide(di.StructConstructor(new(combineService))),
			fx.Provide(di.Bind(new(combineService), new(combineservice.CombineService))),
			fx.Provide(newKitexServer),
			fx.Invoke(startKitexServer))
	}
	return fx.Options(opts...)
}

var _ oasyce.UserService = (*combineService)(nil)

type combineService struct {
	oasyce.UserService
}

func newKitexServer(service combineservice.CombineService) server.Server {
	// get listening addr from config
	kitexAddr, _ := net.ResolveTCPAddr("tcp", conf.GlobalConfig().API.Kitex.Addr)
	opts := []server.Option{
		server.WithServiceAddr(kitexAddr),
	}
	for _, mw := range kitexMiddleWares() {
		opts = append(opts, server.WithMiddleware(mw))
	}
	// Start kitex server.
	return combineservice.NewServer(service, opts...)
}

func newKitexClient() (combineservice.Client, error) {
	opts := []client.Option{
		client.WithRPCTimeout(timeout.DefaultRPCTimeout),
		client.WithLongConnection(connpool.IdleConfig{
			MaxIdlePerAddress: 100,
			MaxIdleGlobal:     100,
			MaxIdleTimeout:    600 * time.Second,
		}),
		client.WithTransportProtocol(transport.TTHeaderFramed),
		client.WithHostPorts(conf.GlobalConfig().API.Kitex.Addr),
		// fall back to default cluster.
		client.WithFailureRetry(kitexretry.NewFailurePolicyWithResultRetry(&kitexretry.ShouldResultRetry{
			NotRetryForTimeout: true, // in order to apply the ErrorRetry above, need to set this to true, ref: github.com/cloudwego/kitex@v0.9.0/pkg/retry/failure_retryer.go, failureRetryer.isRetryErr
		})),
		client.WithMiddleware(timeout.KitexClientMiddleware),
	}
	combineClient, err := combineservice.NewClient("code.app.rpc", opts...)
	if err != nil {
		return nil, err
	}
	return combineClient, nil
}

func startKitexServer(lc fx.Lifecycle, server server.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error, 1)
			go func() {
				errChan <- server.Run()
			}()
			select {
			case err := <-errChan:
				return err
			case <-time.After(5 * time.Second):
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop()
		},
	})
}

func kitexMiddleWares() []endpoint.Middleware {
	return []endpoint.Middleware{}
}
