package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var donations []donation

var total float64 = 0

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "release" {
			gin.SetMode(gin.ReleaseMode)
		}
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.Static("/public", "./public")

	donations = getAllDonations()

	fmt.Printf("%v\n", donations)

	for _, donation := range donations {
		total += donation.AmountNumber
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"donations": donations,
			"total":     fmt.Sprintf("%.2f", total),
		})
	})

	router.POST("/webhook", func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Printf("Error reading webhook")
			return
		}

		var stripeWebhookData stripeWebhookResponse

		json.Unmarshal(jsonData, &stripeWebhookData)

		var newDonation = stripeWebhookData.Data.Object.ToDonation()

		donations = append([]donation{newDonation}, donations...)
		total += newDonation.AmountNumber

		c.HTML(http.StatusCreated, "donate.tmpl", gin.H{
			"donations": donations,
			"total":     fmt.Sprintf("%.2f", total),
		})
	})

	router.Run(":8080")
}
