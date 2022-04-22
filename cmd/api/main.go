package main

import (
	"go-clean-arch/cmd/api/conf"
	"go-clean-arch/internal/platform/environment"
)

func main() {
	r := conf.SetupRoutes(conf.GetDependencies())
	port := environment.GetEnvOrConfig("PORT")
	r.Run(":" + port)
}
