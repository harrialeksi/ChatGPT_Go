/*
show me an nsq middleware in go language with rate limiting

In this example, we create a rate limiter with a limit of 10 requests per second using the rate package from the Go standard library. We then create an
NSQ consumer and add a middleware to implement rate limiting using the limiterMiddleware function. The limiterMiddleware function waits for the rate limiter 
to allow the request before calling the original message handler. The middleware is added to the consumer using the AddConcurrentHandlers method. 
Finally, we connect the consumer to the NSQD instance and wait for messages to be processed.
*/

package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
	"golang.org/x/time/rate"
)

func main() {
	// Create a new rate limiter with a limit of 10 requests per second
	limiter := rate.NewLimiter(10, 1)

	// Create a new NSQ consumer
	consumer, err := nsq.NewConsumer("topic", "channel", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Add a middleware that implements rate limiting
	consumer.AddConcurrentHandlers(
		limiterMiddleware(limiter, handle),
		1,
	)

	// Connect the consumer to the NSQD instance
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for messages to be processed
	select {}
}

func handle(message *nsq.Message) error {
	log.Printf("Received message: %s\n", message.Body)
	return nil
}

// Middleware that implements rate limiting
func limiterMiddleware(limiter *rate.Limiter, handler nsq.Handler) nsq.Handler {
	return nsq.HandlerFunc(func(message *nsq.Message) error {
		// Wait for the rate limiter to allow the request
		err := limiter.Wait(message.Timestamp)
		if err != nil {
			return err
		}

		// Call the original handler
		return handler.HandleMessage(message)
	})
}
