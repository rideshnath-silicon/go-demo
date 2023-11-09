package main

import (
	studentmodel "gin/Student/StudentModel"
	"gin/config"
	"gin/routes"
	user "os/user"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariable()
	config.ConnectDatabase()

}

func main() {
	err := config.DB.AutoMigrate(&user.User{}, &studentmodel.Student{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	r := gin.Default()
	routes.IndexRouter(r)
	r.Run()
}
