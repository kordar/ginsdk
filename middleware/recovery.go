package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecoveryMiddleware RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusOK, map[string]interface{}{"code": 500, "message": err})
				return
			}
		}()
		c.Next()
	}
}
