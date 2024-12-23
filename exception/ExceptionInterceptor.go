package exception

import "github.com/gin-gonic/gin"

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Tangkap error yang muncul di middleware atau handler
		err := c.Errors.Last()
		if err != nil {
			// Tangani error di sini
			switch errorType := err.Err.(type) {
			case *ClientError:
				c.JSON(errorType.Code, gin.H{
					"error": errorType.Message,
				})
			default:
				// Tangani error lainnya
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
			}

			// Hentikan eksekusi konteks
			c.Abort()
		}
	}
}
