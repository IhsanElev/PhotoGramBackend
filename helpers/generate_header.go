package helpers

import "github.com/gin-gonic/gin"

func GetHeaderValue(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
