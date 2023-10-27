package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var donations []donation

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.Static("/assets", "./assets")

	donations = getAllDonations()

	fmt.Printf("%v\n", donations)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"donations": donations,
		})
	})

	router.POST("/webhook", func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Printf("Error reading webhook")
		}

		var stripeWebhookData stripeWebhookResponse

		json.Unmarshal(jsonData, &stripeWebhookData)

		var newDonation = sessionToDonation(stripeWebhookData.Data.Object)

		donations = append([]donation{newDonation}, donations...)
	})

	router.Run(":8080")
}
