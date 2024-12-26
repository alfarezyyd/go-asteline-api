package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Interceptor() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer func() {
			if occurredError := recover(); occurredError != nil {
				// Check if it's our custom
				if clientError, ok := occurredError.(ClientError); ok {
					ginContext.AbortWithStatusJSON(clientError.StatusCode, gin.H{
						"message": clientError.Message,
					})
					return
				}

				// Unknown
				ginContext.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server occuredErroror",
				})
			}
		}()
		ginContext.Next()
	}
}
