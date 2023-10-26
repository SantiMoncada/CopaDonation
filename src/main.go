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

	fmt.Printf("%v/n", donations)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"donations": donations,
		})
	})

	type stripeWebhookResponse struct {
		Created int `json:"created"`
		Data    struct {
			Object struct {
				Amount       int    `json:"amount_total"`
				Created      int    `json:"created"`
				Currency     string `json:"currency"`
				CustomFields []struct {
					Key      string `json:"key"`
					Dropdown struct {
						Value string `json:"value"`
					} `json:"dropdown"`
					Text struct {
						Value string `json:"value"`
					} `json:"text"`
				} `json:"custom_fields"`
			} `json:"object"`
		} `json:"data"`
	}

	router.POST("/webhook", func(c *gin.Context) {
		jsonData, err := io.ReadAll(c.Request.Body)
		if err == nil {
			fmt.Printf("Error reading webhook")
		}

		var stripeWebhookData stripeWebhookResponse
		json.Unmarshal(jsonData, &stripeWebhookData)

		var message string
		var bootcamp string
		var ammount string

		for _, custom_field := range stripeWebhookData.Data.Object.CustomFields {
			if custom_field.Key == "bootcamp" {
				bootcamp = custom_field.Dropdown.Value
				continue
			}

			if custom_field.Key == "messageforthefeed" {
				message = custom_field.Text.Value
				continue
			}
		}

		ammount = fmt.Sprintf("%d.%s", stripeWebhookData.Data.Object.Amount/100, toFixed2(stripeWebhookData.Data.Object.Amount%100))

		donations = append(donations, donation{ammount, message, bootcamp})
	})

	router.Run(":8080")
}
