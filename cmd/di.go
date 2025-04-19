package cmd

import (
	"go.uber.org/fx"

	"Oasyce-backend/module/user"
)

type buildTarget int

const (
	BuildTargetAPI buildTarget = iota
	BuildTargetRPC
)

func Module(target buildTarget) fx.Option {
	return fx.Options(
		hertzModule(target),
		kitexModule(target),
		jwtModule,

		user.Module,
	)
}
