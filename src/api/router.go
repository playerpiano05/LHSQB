package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(gin *gin.Engine) {

}

func setupGET(gin *gin.Engine) {
	gin.GET("/ping", loginRoute)
}

/*func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

*/
