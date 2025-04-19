package cmd

import (
	"Oasyce-backend/util/conf"
	"Oasyce-backend/util/logs"
	"go.uber.org/fx"

	jwtutil "Oasyce-backend/util/jwt"
)

var jwtModule = fx.Provide(
	func() *jwtutil.Authenticator {
		if conf.GlobalConfig().JWT.PublicKey == "" || conf.GlobalConfig().JWT.PrivateKey == "" {
			return nil
		}
		jwtAuth, err := jwtutil.NewAuthenticator(conf.GlobalConfig().JWT.PublicKey, conf.GlobalConfig().JWT.PrivateKey)
		if err != nil {
			logs.Logger.Errorw("failed to new jwt auth", logs.Error(err))
			return nil
		}
		return jwtAuth
	},
)
