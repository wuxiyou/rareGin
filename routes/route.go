package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"rareGin/controllers/admin"
	"rareGin/utils"
)

func Init()  {
	fmt.Println(gin.Version)

	routes := gin.Default()

	webApi := routes.Group("/webApi")

	adminControllerLogin := new(admin.Login)
	webApi.GET("login",adminControllerLogin.LoginInfo)



	routes.Run(":"+utils.ReadConfigInfo("http", "listen_port"))
}