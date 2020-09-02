package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HistoryAPIFallback Return the index.html and enable the history API fallback functionality
func HistoryAPIFallback() gin.HandlerFunc {
	// Serve Statics
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		}
		var contentType = c.Request.Header.Get("Accept")
		if strings.Contains(contentType, "text/html") {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		}
		c.Next()
	}
}
