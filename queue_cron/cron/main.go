package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	// library to interact with cron
	"github.com/go-co-op/gocron"
)

type UserData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Scheduled function
func SendRequest() {

	// Setup http client
	client := &http.Client{}

	// Setup request body
	body := map[string]interface{}{
		"name": "Hansen Andy",
		"age":  25,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP request
	// Change the target url with your own url
	// from webhook.site
	req, err := http.NewRequest("POST", "https://webhook.site/21d8bad7-1af5-40cc-ad04-e34ad1072059", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Send request
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	// Setup cron
	sched := gocron.NewScheduler(time.Local)

	// Setup scheduled request sending operation
	// similar to crontab -e
	sched.Cron("32 16 27 12 *").Do(SendRequest)

	// start cron
	sched.StartBlocking()
}
