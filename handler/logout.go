package handler

import (
	"gowebapp/handler/auth"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	// 转由jwt中间件处理登出流程
	auth.JwtWrapper.LogoutHandler(c)
}
