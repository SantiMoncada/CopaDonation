package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.Static("/assets", "./assets")

	paymentIntents := getPaymentIntents()

	type templateDataType struct {
		Amounts []string
	}

	var templateData templateDataType

	for _, val := range paymentIntents {
		templateData.Amounts = append(templateData.Amounts, fmt.Sprintf("%d.%d", val.Amount/100, int(math.Mod(float64(val.Amount), 100))))
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"payments": paymentIntents,
		})
	})
	router.Run(":8080")
}
