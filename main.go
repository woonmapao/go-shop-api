package main

import (
	"github.com/woonmapao/user-management/initializer"
	"github.com/woonmapao/user-management/routes"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.DBInitializer()
}

func main() {
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
