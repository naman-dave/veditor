package server

import (
	"github.com/gin-gonic/gin"
)

func authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
