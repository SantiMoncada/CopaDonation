package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		texto := c.DefaultQuery("texto", "none")
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"test": texto,
		})
	})
	router.Run(":8080")
}
