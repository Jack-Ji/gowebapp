package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gowebapp/handler/common"
	"gowebapp/model"
	"gowebapp/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	JwtWrapper  *jwt.GinJWTMiddleware // jwt中间件
	IdentityKey = "LM_IDENTITY"       // token存取键
	LoginKey    = "LM_LOGIN"          // login请求存取键
	CookieName  = "LM_JWT"            // 存放令牌的cookie名称
)

func Init() {
	var err error
	JwtWrapper, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm: "LM auth",

		// 用于在会话上下文中获取Token
		// handler可通过ctx.Get(IdentityKey)获取Token
		IdentityKey: IdentityKey,

		// jwt签名私有串
		Key: []byte("LMSoftwareStudio"),

		// 会话刷新最短间隔
		Timeout: time.Hour,

		// 会话最长有效期
		MaxRefresh: time.Hour * 12,

		// 校验用户名和密码并生成Token
		Authenticator: func(c *gin.Context) (interface{}, error) {
			v, ok := c.Get(LoginKey)
			if !ok {
				return nil, fmt.Errorf("获取登录数据失败")
			}
			login, ok := v.(*Login)
			if !ok {
				return nil, fmt.Errorf("登录数据类型错误")
			}

			// 获取用户信息并校验密码
			user := model.User{}
			err := user.GetByName(login.UserName)
			if err != nil {
				if err == model.ErrNotExist {
					return nil, fmt.Errorf("用户不存在")
				}
				return nil, fmt.Errorf("内部错误：%s", err)
			}
			if !utils.VerifySaltedPasswd(login.Password, *user.Password, *user.Salt) {
				log.Printf("用户'%s'密码错误\n", login.UserName)
				return nil, fmt.Errorf("用户或密码错误")
			}
			login.UserID = user.ID

			return &Token{
				UserID:   user.ID,
				UserName: user.Name,
			}, nil
		},

		// 返回登录结果，
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			v, _ := c.Get(LoginKey)
			login := v.(*Login)
			login.Token = token
			login.TokenExpire = expire.Format(time.RFC3339)
			c.JSON(http.StatusOK, common.Response{Code: common.ErrCodeSuccess, Msg: "成功", Data: login})
		},

		// 返回登出结果
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, common.Response{Code: common.ErrCodeSuccess, Msg: "成功"})
		},

		// 转换Token为键值对，交由jwt一起携带
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Token); ok {
				return jwt.MapClaims{
					"user_id":   fmt.Sprintf("%d", v.UserID),
					"user_name": v.UserName,
				}
			}
			return jwt.MapClaims{}
		},

		// 从http请求中提取Token并返回给上下文
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			token := &Token{
				UserName: claims["user_name"].(string),
			}
			token.UserID, _ = strconv.ParseInt(claims["user_id"].(string), 0, 64)
			return token
		},

		// TODO 根据Token.Permission检查api调用是否合法
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},

		// 用户名密码错误，会话数据非法、超时，调用未授权接口等情况的处理
		Unauthorized: func(c *gin.Context, code int, message string) {
			ec := common.ErrCodePriviledge
			if strings.Contains(message, "expired") {
				ec = common.ErrCodeSessionGone
			}
			c.JSON(http.StatusOK, common.Response{
				Code: ec,
				Msg:  message,
			})
		},

		// token读取位置：cookie优先，其次为http头的Authorization
		TokenLookup: fmt.Sprintf("cookie: %s, header: Authorization", CookieName),

		// http头放置token时附带的前缀
		TokenHeadName: "LM_TOKEN",

		// 时间获取函数
		TimeFunc: time.Now,

		// 阻止客户端js获取cookie
		CookieHTTPOnly: true,

		// 存储jwt令牌的cookie名称
		CookieName: CookieName,
		SendCookie: true,
	})
	if err != nil {
		panic(fmt.Errorf("create jwt middleware failed: %s", err))
	}
}

// gin中间件，用于自动刷新token
func JwtRefresher() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		now := JwtWrapper.TimeFunc().Unix()
		expireAt := int64(claims["exp"].(float64))

		//token := MustGetToken(c)
		//res, _ := time.ParseDuration(fmt.Sprintf("%ds", expireAt-now))
		//log.Printf("会话超期检测：用户<%s@%s> 剩余时间<%s>\n", token.UserName, token.CorpName, res)

		if expireAt-now < 600 {
			// 会话即将失效，尝试刷新
			// 注意：如果会话已经持续保活超过JwtWrapper.MaxRefresh指定时间，则必须重新登录
			newToken, _, err := JwtWrapper.RefreshToken(c)
			if err != nil {
				if err == jwt.ErrExpiredToken {
					c.JSON(http.StatusOK, common.Response{
						Code: common.ErrCodeSessionGone,
						Msg:  "登录会话已失效，请重新登录",
					})
				} else {
					c.JSON(http.StatusOK, common.Response{
						Code: common.ErrCodeInternal,
						Msg:  fmt.Sprintf("会话刷新失败：%s", err),
					})
				}
				return
			}

			//log.Printf("会话刷新：用户<%s@%s> 超期时间<%s>\n", token.UserName, token.CorpName, future.Sub(time.Now()))

			// 通过头部返回更新后的token
			c.Header("Authorization", JwtWrapper.TokenHeadName+" "+newToken)
		} else {
			// 通过头部返回当前token
			c.Header("Authorization", JwtWrapper.TokenHeadName+" "+jwt.GetToken(c))
		}

		c.Next()
	}
}
