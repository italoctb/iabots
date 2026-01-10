package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ContextUserIDKey = "user_id"
)

// DevIdentityMiddleware
// ⚠️ Middleware TEMPORÁRIO para desenvolvimento.
// Injeta user_id no contexto a partir do header X-User-ID.
func DevIdentityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDHeader := c.GetHeader("X-User-ID")
		if userIDHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "missing X-User-ID header",
			})
			return
		}

		userID, err := uuid.Parse(userIDHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid X-User-ID",
			})
			return
		}

		// Injeta no contexto
		c.Set(ContextUserIDKey, userID)

		c.Next()
	}
}
