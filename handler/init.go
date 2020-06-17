package handler

import (
	"gowebapp/handler/auth"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	api := e.Group("/api")
	api.POST("/user/login", LoginHandler)
	api.POST("/user/logout", LogoutHandler)

	auth.Init()
	jwtSecured := api.Group("")
	jwtSecured.Use(auth.JwtWrapper.MiddlewareFunc())
	jwtSecured.Use(auth.JwtRefresher())
	{
		// TODO
	}
}
