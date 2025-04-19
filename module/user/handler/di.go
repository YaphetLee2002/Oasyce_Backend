package handler

import (
	"go.uber.org/fx"

	"Oasyce-backend/api/kitex/kitex_gen/oasyce"
	"Oasyce-backend/util/di"
)

var Module = fx.Options(
	fx.Provide(di.StructConstructor(new(Handler))),
	fx.Provide(di.Bind(new(Handler), new(oasyce.UserService))),
)
