package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func WithBasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("USERNAME"): os.Getenv("PASSWORD"),
	})
}
