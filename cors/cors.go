package cors

import "github.com/gin-gonic/gin"

//Init CORS
func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

//Options CORS
func Options(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
