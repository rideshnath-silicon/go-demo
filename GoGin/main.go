package main

import (
	"gin/config"
	"gin/routes"

	"github.com/gin-gonic/gin"
)



func init() {
	config.LoadEnvVariable()
	config.ConnectDatabase()

}

func main() {
	r := gin.Default()
	routes.UserRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
