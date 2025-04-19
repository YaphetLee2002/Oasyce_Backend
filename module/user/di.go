package user

import (
	"go.uber.org/fx"

	"Oasyce-backend/module/user/handler"
)

var Module = fx.Options(
	handler.Module,
)
