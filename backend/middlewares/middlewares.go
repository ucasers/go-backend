package middlewares

import (
	"github.com/ucasers/go-backend/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ucasers/go-backend/backend/auth"
)

// TokenAuthMiddleware checks if the request contains a valid JWT token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		user, err := dao.Q.User.
			WithContext(c).
			Where(dao.User.ID.Eq(userID)).
			First()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "User not found",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

// CORSMiddleware sets the necessary headers for CORS and handles preflight requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
