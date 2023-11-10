package main

import (
	classmodel "gin/Class/ClassModel"
	studentmodel "gin/Student/StudentModel"
	submodel "gin/Subject/subModel"
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
	err := config.DB.AutoMigrate(&user.User{}, &studentmodel.Student{}, &submodel.Subject{}, &classmodel.Class{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	r := gin.Default()
	routes.IndexRouter(r)
	r.Run()
}
