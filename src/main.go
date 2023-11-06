package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var donations []donation

var total float64 = 0

var channelCount uint64 = 0

var streamChannels = make(map[uint64]chan donation)

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

		for _, channel := range streamChannels {
			channel <- newDonation
		}

	})

	router.GET("/event-stream", func(c *gin.Context) {
		var id = channelCount
		channelCount++

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

	// router.POST("/test-stream", func(c *gin.Context) {

	// 	for index, channel := range streamChannels {

	// 		fmt.Printf(fmt.Sprintf("sending event to id:%d\n", index))
	// 		channel <- donations
	// 	}
	// })

	log.Fatalf("error running HTTP server: %s\n", router.Run(":8080"))
}
