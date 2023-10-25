package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseApi = "https://api.stripe.com/v1"
const stripeKey = "rk_test_51O3hRaJQy7oGudPMNpcZFOWzma0AE5zyF1290grVx7u12LvjQAofzO9iwPUS6GXoWuttVqgSyZIC8fPI4zPDd3US00GIXtBJtL"

type stripeResponse[T any] struct {
	Object  string `json:"object"`
	Data    []T    `json:"data"`
	HasMore bool   `json:"has_more"`
	Url     string `json:"url"`
}

type paymentIntent struct {
	Id      string `json:"id"`
	Amount  int    `json:"amount"`
	Created int    `json:"created"`
}

func getPaymentIntents() []paymentIntent {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment_intents", baseApi), nil)

	if err != nil {
		log.Fatal("Error creating request:")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", stripeKey))

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:")
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Could not parse starting payment intents")
	}

	var jsonResponse stripeResponse[paymentIntent]

	json.Unmarshal(responseData, &jsonResponse)

	return jsonResponse.Data
}

type checkoutSession struct {
	Id           string `json:"id"`
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
}

func getSessionData(id string) checkoutSession {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/checkout/sessions?payment_intent=%s", baseApi, id), nil)

	if err != nil {
		log.Fatal("Error creating request:")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", stripeKey))

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:")
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Could not parse starting payment intents")
	}

	var jsonResponse stripeResponse[checkoutSession]

	json.Unmarshal(responseData, &jsonResponse)

	return jsonResponse.Data[0]
}

type donation struct {
	Amount   string
	Message  string
	Bootcamp string
}

func getAllDonations() []donation {
	var sessions []checkoutSession

	intents := getPaymentIntents()

	for _, intent := range intents {
		sessions = append(sessions, getSessionData(intent.Id))
	}

	var output []donation

	for _, session := range sessions {
		var message string
		var bootcamp string
		var ammount string

		for _, custom_field := range session.CustomFields {
			if custom_field.Key == "bootcamp" {
				bootcamp = custom_field.Dropdown.Value
				continue
			}

			if custom_field.Key == "messageforthefeed" {
				message = custom_field.Text.Value
				continue
			}

		}

		ammount = fmt.Sprintf("%d.%d", session.Amount/100, session.Amount%100)

		output = append(output, donation{ammount, message, bootcamp})
	}

	return output
}
