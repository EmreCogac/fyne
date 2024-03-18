package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/data/left", test)
	r.GET("/data/right", test)
	r.GET("/test", test)

	return r

}
func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"json data": "json data",
	})
}
