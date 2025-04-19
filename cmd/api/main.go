package main

import (
	"go.uber.org/fx"

	"Oasyce-backend/cmd"
)

func main() {
	fx.New(cmd.Module(cmd.BuildTargetAPI)).Run()
}
