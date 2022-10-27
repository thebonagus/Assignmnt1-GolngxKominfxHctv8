package middlewares

import (
	"final-project/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if verifyToken, err := helpers.VerifyToken(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		} else {
			c.Set("userData", verifyToken)
			c.Next()
		}
	}
}
