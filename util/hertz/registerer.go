package hertz

import (
	"github.com/cloudwego/hertz/pkg/route"
	"go.uber.org/fx"
)

// Registerer is an abstract for hertz handlers
type Registerer interface {
	// Register routes to hertz router
	Register(router route.IRouter)
}

// Param declares all Registerers, so fx will resolve the dependency
// ref: https://github.com/uber-go/fx/blob/master/inout.go
type Param struct {
	fx.In
	Registerers []Registerer `group:"hertz_registerer"`
}

// Annotate the constructor function so the returned value will be used as hertz register
func Annotate(t interface{}) interface{} {
	return fx.Annotated{
		Group:  "hertz_registerer",
		Target: t,
	}
}
