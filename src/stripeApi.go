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

type paymentIntent struct {
	Id      string `json:"id"`
	Amount  int    `json:"amount"`
	Created int    `json:"created"`
}

type stripePaymentsIntentsResponse struct {
	Object  string          `json:"object"`
	Data    []paymentIntent `json:"data"`
	HasMore bool            `json:"has_more"`
	Url     string          `json:"url"`
}

func getPaymentIntents() []paymentIntent {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment_intents", baseApi), nil)

	if err != nil {
		log.Fatal("Error creating request:")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", stripeKey)) // Add any other headers you need

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:")
	}

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Could not parse starting payment intents")
	}

	var jsonResponse stripePaymentsIntentsResponse

	json.Unmarshal(responseData, &jsonResponse)

	return jsonResponse.Data
}
