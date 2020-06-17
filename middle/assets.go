package middle

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ServeAssets(pathPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") &&
			!strings.HasPrefix(c.Request.URL.Path, pathPrefix) {
			c.Redirect(http.StatusPermanentRedirect, pathPrefix+c.Request.URL.Path)
			return
		}

		c.Next()
	}
}
