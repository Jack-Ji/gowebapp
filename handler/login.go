package handler

import (
	"log"
	"net/http"

	"gowebapp/handler/auth"
	"gowebapp/handler/common"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	type postData struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	param := &postData{}
	err := c.ShouldBindJSON(param)
	if err != nil {
		log.Printf("非法消息体: %s", err)
		resp := &common.Response{
			Code: common.ErrCodeParamInvalid,
			Msg:  "参数非法: " + err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 转由jwt中间件处理登录流程
	login := &auth.Login{}
	login.UserName = param.Account
	login.Password = param.Password
	c.Set(auth.LoginKey, login)
	auth.JwtWrapper.LoginHandler(c)
}
