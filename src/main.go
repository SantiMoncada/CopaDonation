package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.Static("/assets", "./assets")

	donations := getAllDonations()
	fmt.Printf("%+v\n", donations)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"donations": donations,
		})
	})
	router.Run(":8080")
}
