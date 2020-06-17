package middle

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeCORS(additionHeaders ...string) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	for _, v := range additionHeaders {
		cfg.AllowHeaders = append(cfg.AllowHeaders, v)
	}
	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}
