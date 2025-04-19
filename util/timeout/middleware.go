package timeout

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

const (
	DefaultRPCTimeout = 10 * time.Second
	defaultAPITimeout = 30 * time.Second
)

const (
	metainfoDDLKey = "Code-Request-Deadline"
)

func KitexClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		if clientDDL, ok := ctx.Deadline(); ok {
			// Set Kitex request's ddl to min(now+DefaultRPCTimeout, API ddl)
			var requestDDL time.Time
			serverDDL := time.Now().Add(DefaultRPCTimeout)
			if clientDDL.Before(serverDDL) {
				requestDDL = clientDDL
			} else {
				requestDDL = serverDDL
			}
			var cancel context.CancelFunc
			ctx, cancel = context.WithDeadline(ctx, requestDDL)
			defer cancel()
			ctx = metainfo.WithPersistentValue(ctx, metainfoDDLKey, requestDDL.Format(time.RFC3339Nano))
		}
		return next(ctx, req, resp)
	}
}
