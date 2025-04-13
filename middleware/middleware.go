package middleware

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Token Not present",
		})
		return
	}
	c.Next() // ไป func ที่ call ต่อไป
}

// other way to write middle ware
func AuthenticateOther() gin.HandlerFunc {
	// can write custom login logic applied before my middleware is executed
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") == "auth") {
			c.AbortWithStatusJSON(500, gin.H{
				"Message": "Token Not present",
			})
			return
		}
		c.Next() // ไป func ที่ call ต่อไป
	}
}

func Addheder(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}
