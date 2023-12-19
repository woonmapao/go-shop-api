package main

import (
	"github.com/gin-gonic/gin"
	"github.com/woonmapao/user-management/controllers"
	"github.com/woonmapao/user-management/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.DBInitializer()
}

func main() {
	r := gin.Default()

	r.POST("/addUser", controllers.AddUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
