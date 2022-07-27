package router

import (
	"AnimeCat/controllers"
	"AnimeCat/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/api/account", controllers.LoginHandler)
	api := r.Group("/api", middlewares.JWTAuthMiddleware())
	{
		api.GET("/setting", controllers.GetSettingHandler)
	}
	err := r.Run(":80")
	if err != nil {
		log.Fatalf("InitRouter failed, the err is %s", err)
	}
}
