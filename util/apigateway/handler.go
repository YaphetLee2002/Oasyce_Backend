package apigateway

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
)

func Handler(client interface{}) func(ctxRaw context.Context, ctx *app.RequestContext) {
	clientValue := reflect.ValueOf(client)
	return func(ctxRaw context.Context, ctx *app.RequestContext) {
		respMeta := ParseMetadata(ctx)
		action, ok := ctx.GetQuery("Action")
		if !ok {
			ctx.JSON(http.StatusBadRequest, &Response{
				ResponseMetadata: respMeta,
				Result:           "missing action",
			})
			return
		}

		method := clientValue.MethodByName(action)
		if !method.IsValid() {
			ctx.JSON(http.StatusBadRequest, &Response{
				ResponseMetadata: respMeta,
				Result:           "invalid action",
			})
			return
		}

		var params []reflect.Value
		if method.Type().NumIn() == 3 {
			// 创建请求参数对象
			typ := method.Type().In(1)
			v := reflect.New(typ.Elem())
			// 绑定请求参数
			if err := binding.BindAndValidate(ctx.GetRequest(), v.Interface(), ctx.Params); err != nil {
				ctx.JSON(http.StatusBadRequest, &Response{
					ResponseMetadata: respMeta,
					Result:           err.Error(),
				})
				return
			}
			params = []reflect.Value{reflect.ValueOf(ctxRaw), v}
		} else {
			params = []reflect.Value{reflect.ValueOf(ctxRaw)}
		}

		rets := method.Call(params)
		if len(rets) != 2 {
			ctx.JSON(http.StatusInternalServerError, &Response{
				ResponseMetadata: respMeta,
				Result:           "invalid return values",
			})
			return
		}

		if err, _ := rets[1].Interface().(error); err != nil {
			ctx.JSON(http.StatusInternalServerError, &Response{
				ResponseMetadata: respMeta,
				Result:           err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, &Response{
			ResponseMetadata: respMeta,
			Result:           rets[0].Interface(),
		})
	}
}

type ErrorData struct {
	Code    string            `json:"Code"`
	Message string            `json:"Message"`
	Data    map[string]string `json:"Data"`
}

// ResponseMetadata ...
type ResponseMetadata struct {
	RequestID string     `json:"RequestId"`
	Action    string     `json:"Action"`
	Version   string     `json:"Version"`
	Service   string     `json:"Service"`
	Error     *ErrorData `json:"Error,omitempty"`
}

// Response ...
type Response struct {
	ResponseMetadata *ResponseMetadata `json:"ResponseMetadata"`
	Result           interface{}       `json:"Result,omitempty"`
}

// ParseMetadata ...
func ParseMetadata(c *app.RequestContext) *ResponseMetadata {
	return &ResponseMetadata{
		RequestID: c.Request.Header.Get("X-Top-Request-Id"),
		Action:    c.Query("Action"),
		Version:   c.Query("Version"),
		Service:   c.Request.Header.Get("X-Top-Service"),
	}
}
