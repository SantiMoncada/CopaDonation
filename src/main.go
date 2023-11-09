package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var donations []donation

var total float64 = 0

var webTotal float64 = 0

var uxTotal float64 = 0

var dataTotal float64 = 0

var streamChannels = make(map[uuid.UUID]chan donation)
var stripeKey = os.Getenv("STRIPE_API_KEY")
var stripeDonateLink = os.Getenv("STRIPE_DONATION_LINK")

func main() {

	if stripeDonateLink == "" {
		stripeDonateLink = "https://donate.stripe.com/test_eVa9CWfaNcsB7pm28c"
	}
	if stripeKey == "" {
		stripeKey = "rk_test_51O3hRaJQy7oGudPMNpcZFOWzma0AE5zyF1290grVx7u12LvjQAofzO9iwPUS6GXoWuttVqgSyZIC8fPI4zPDd3US00GIXtBJtL"
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "release" {
			gin.SetMode(gin.ReleaseMode)
		}
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.Static("/public", "./public")

	donations = getAllDonations()

	for _, donation := range donations {
		total += donation.AmountNumber

		switch donation.Bootcamp {
		case "web":
			webTotal += donation.AmountNumber
		case "ux":
			uxTotal += donation.AmountNumber
		case "data":
			dataTotal += donation.AmountNumber
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "donate.tmpl", gin.H{
			"donateLink": stripeDonateLink,
			"donations":  donations,
			"total":      fmt.Sprintf("%.2f", total),
			"webTotal":   webTotal,
			"uxTotal":    uxTotal,
			"dataTotal":  dataTotal,
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

		switch newDonation.Bootcamp {
		case "web":
			webTotal += newDonation.AmountNumber
		case "ux":
			uxTotal += newDonation.AmountNumber
		case "data":
			dataTotal += newDonation.AmountNumber
		}

		c.HTML(http.StatusCreated, "donate.tmpl", gin.H{
			"donateLink": stripeDonateLink,
			"donations":  donations,
			"total":      fmt.Sprintf("%.2f", total),
			"webTotal":   webTotal,
			"uxTotal":    uxTotal,
			"dataTotal":  dataTotal,
		})

		for _, channel := range streamChannels {
			channel <- newDonation
		}

	})

	router.GET("/event-stream", func(c *gin.Context) {
		id := uuid.New()

		ch := make(chan donation)

		streamChannels[id] = ch

		c.Stream(func(w io.Writer) bool {
			msg, ok := <-ch
			if ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})

		close(streamChannels[id])
		delete(streamChannels, id)
	})

	router.GET("api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"total":     total,
			"webTotal":  webTotal,
			"uxTotal":   uxTotal,
			"dataTotal": dataTotal,
			"donations": donations,
		})
	})

	log.Fatalf("error running HTTP server: %s\n", router.Run(":3000"))
}
