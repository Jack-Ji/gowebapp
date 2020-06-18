package auth

import (
	"github.com/gin-gonic/gin"
)

// 内部token
type Token struct {
	UserID   int64  // 用户ID
	UserName string // 用户名称
}

// 用户登录数据
type Login struct {
	Token       string `json:"token"`
	TokenExpire string `json:"tokenExpire"`
	UserID      int64  `json:"id"`
	UserName    string `json:"account"`
	Password    string `json:"-"`
}

// 获取登录会话token，失败则panic
func MustGetToken(c *gin.Context) *Token {
	v, ok := c.Get(IdentityKey)
	if !ok {
		panic("获取token失败")
	}
	return v.(*Token)
}
