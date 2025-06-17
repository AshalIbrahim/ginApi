package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//a:=[5]int{}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "get request received",
		})
	})



	r.Run(":8080")
}
