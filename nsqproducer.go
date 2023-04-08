/*
show me an nsq producer from standard net/http handler function

In this example, we create an NSQ producer using the nsq package from the go-nsq library. We then create an HTTP handler function that parses the request body 
into a Message struct, publishes the message to NSQ using the producer, and returns a success response. Finally, we start an HTTP server to handle incoming requests.

Note that in production code, you should handle errors and implement appropriate error handling, such as logging or retrying failed message publishes.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nsqio/go-nsq"
)

type Message struct {
	Body string `json:"body"`
}

func main() {
	// Create a new NSQ producer
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Stop()

	// Create an HTTP handler function that publishes messages to NSQ
	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body into a Message struct
		var message Message
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Publish the message to NSQ
		err = producer.Publish("topic", []byte(message.Body))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Return a success response
		w.WriteHeader(http.StatusCreated)
	})

	// Start the HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
