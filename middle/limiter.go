package middle

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(maxEventsPerSec int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), 10)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}

		log.Printf("服务器过载！url:%s", c.Request.URL.String())
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}
