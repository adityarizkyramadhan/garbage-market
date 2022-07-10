package main

import "github.com/gin-gonic/gin"

func main() {
	// Your code here...
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		data := c.Request.FormValue("nama")
		c.JSON(200, gin.H{
			"message": "Hello World!",
			"data":    data,
		})
	})
	r.Run()
}
