package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	api := r.Group("/api")

	api.GET("/neo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Follow the white rabbit",
		})
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run(":3000")
}
