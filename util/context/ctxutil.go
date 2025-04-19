package context

import (
	"context"
	"github.com/AlekSi/pointer"
)

// [ request id ]

type ctxKeyTypeRequestID struct{}

var ctxKeyRequestID = ctxKeyTypeRequestID{}

func WithRequestID(ctx context.Context, reqID string) context.Context {
	if reqID == "" {
		return ctx
	}
	return context.WithValue(ctx, ctxKeyRequestID, reqID)
}

func GetRequestID(ctx context.Context) string {
	val := ctx.Value(ctxKeyRequestID)
	if val != nil {
		return val.(string)
	}
	return ""
}

// [ user info ]

type ctxKeyTypeUserInfo struct{}

var ctxKeyUserInfo = ctxKeyTypeUserInfo{}

//func WithUserInfo(ctx context.Context, user *accountmodel.User) context.Context {
//	if user == nil {
//		return ctx
//	}
//	return context.WithValue(ctx, ctxKeyUserInfo, user)
//}
//
//func GetUserInfo(ctx context.Context) *accountmodel.User {
//	val := ctx.Value(ctxKeyUserInfo)
//	if val != nil {
//		return val.(*accountmodel.User)
//	}
//	return nil
//}

// [repository model]

type ctxKeyTypeRepository struct{}

var ctxKeyRepository = ctxKeyTypeRepository{}

func WithRepository(ctx context.Context, repo interface{}) context.Context {
	return context.WithValue(ctx, ctxKeyRepository, repo)
}

// GetRepository extracts repository injected by middleware loader.
//
// We use interface here because we don't want to depend on business model struct.
func GetRepository(ctx context.Context) interface{} {
	return ctx.Value(ctxKeyRepository)
}

// [authctx]

type ctxKeyTypeAuthCtx struct{}

var ctxKeyAuthCtx = ctxKeyTypeAuthCtx{}

func WithAuthCtx(ctx context.Context, authCtx interface{}) context.Context {
	return context.WithValue(ctx, ctxKeyAuthCtx, authCtx)
}

func GetAuthCtx(ctx context.Context) interface{} {
	return ctx.Value(ctxKeyAuthCtx)
}

// [ Bypass TOP request flag]

type ctxKeyTypeBypassTopRequestFlag struct{}

var ctxKeyBypassTopRequestFlag = ctxKeyTypeBypassTopRequestFlag{}

func WithBypassTOPRequestFlag(ctx context.Context, bypassRequestFlag bool) context.Context {
	return context.WithValue(ctx, ctxKeyBypassTopRequestFlag, bypassRequestFlag)
}

func GetBypassTOPRequestFlag(ctx context.Context) bool {
	val := ctx.Value(ctxKeyBypassTopRequestFlag)
	if val != nil {
		return val.(bool)
	}
	return false
}

// [ instance id ]

type ctxKeyTypeInstanceID struct{}

var ctxKeyInstanceID = ctxKeyTypeInstanceID{}

func WithInstanceID(ctx context.Context, instanceID *int64) context.Context {
	return context.WithValue(ctx, ctxKeyInstanceID, instanceID)
}

func GetInstanceID(ctx context.Context) *int64 {
	val := ctx.Value(ctxKeyInstanceID)
	if val != nil {
		return val.(*int64)
	}
	return nil
}

// [shard key]

type ctxKeyTypeShardKey struct{}

var ctxKeyShardKey = ctxKeyTypeShardKey{}

func WithDBShardKey(ctx context.Context, repoID int64) context.Context {
	return context.WithValue(ctx, ctxKeyShardKey, repoID)
}

func GetDBShardKey(ctx context.Context) *int64 {
	val := ctx.Value(ctxKeyShardKey)
	if val != nil {
		return pointer.To(val.(int64))
	}
	return nil
}

type ctxKeyTypeDBMode struct{}

var ctxKeyDBMode = ctxKeyTypeDBMode{}

const (
	DBModeRead  = "read"
	DBModeWrite = "write"
)

func WithWriteDB(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyDBMode, DBModeWrite)
}

func WithReadDB(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyDBMode, DBModeRead)
}

func GetDBMode(ctx context.Context) string {
	val := ctx.Value(ctxKeyDBMode)
	if val != nil {
		return val.(string)
	}
	return ""
}

type ctxKeyTypeUseCache struct{}

var ctxKeyUseCache = ctxKeyTypeUseCache{}

func WithUseCache(ctx context.Context, useCache bool) context.Context {
	return context.WithValue(ctx, ctxKeyUseCache, useCache)
}

func IfUseCache(ctx context.Context) bool {
	val := ctx.Value(ctxKeyUseCache)
	if val != nil {
		return val.(bool)
	}
	return true
}

type ctxKeyTypeTargetID struct{}

var ctxKeyTargetID = ctxKeyTypeTargetID{}

func WithTargetID(ctx context.Context, targetID string) context.Context {
	return context.WithValue(ctx, ctxKeyTargetID, targetID)
}

func GetTargetID(ctx context.Context) string {
	val := ctx.Value(ctxKeyTargetID)
	if val != nil {
		return val.(string)
	}
	return ""
}

type ctxKeyTypeMergeRequestID struct{}

var ctxKeyMergeRequestID = ctxKeyTypeMergeRequestID{}

func WithMergeRequestID(ctx context.Context, mrID string) context.Context {
	return context.WithValue(ctx, ctxKeyMergeRequestID, mrID)
}

func GetMergeRequestID(ctx context.Context) string {
	val := ctx.Value(ctxKeyMergeRequestID)
	if val != nil {
		return val.(string)
	}
	return ""
}

//func EnsureLogID(ctx context.Context) context.Context {
//	logID := logutils.ExtractLogID(ctx)
//	if logID == "" {
//		return context.WithValue(ctx, logutils.LOGIDKEY, logid.GenLogID())
//	}
//	return ctx
//}

type ctxKeyTypeAuthType struct{}

var ctxKeyAuthType = ctxKeyTypeAuthType{}

func WithAuthType(ctx context.Context, authType string) context.Context {
	return context.WithValue(ctx, ctxKeyAuthType, authType)
}

func GetAuthType(ctx context.Context) string {
	val := ctx.Value(ctxKeyAuthType)
	if val != nil {
		return val.(string)
	}
	return "unknown"
}

type ctxKeyTypeRateLimitRule struct{}

var ctxKeyRateLimitRule = ctxKeyTypeRateLimitRule{}

func WithRateLimitRule(ctx context.Context, rule interface{}) context.Context {
	return context.WithValue(ctx, ctxKeyRateLimitRule, rule)
}

func GetRateLimitRule(ctx context.Context) interface{} {
	return ctx.Value(ctxKeyRateLimitRule)
}
