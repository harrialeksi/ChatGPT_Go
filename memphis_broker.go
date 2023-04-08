/*
show message broker example using memphisdev/memphis

In this example, we import the github.com/memphisdev/memphis package and create a new Memphis broker using memphis.NewBroker().

We then subscribe to a topic using broker.Subscribe() and publish a message to the topic using broker.Publish().

We wait for the message to be received using time.Sleep().

Finally, we consume messages from the subscription using sub.Next(), print the payload to the console, and acknowledge the message using msg.Ack(). 
Note that acknowledging the message is important to ensure that it is not re-delivered to the subscriber.
*/

package main

import (
    "fmt"
    "time"

    "github.com/memphisdev/memphis"
)

func main() {
    // Create a new Memphis broker
    broker, err := memphis.NewBroker()
    if err != nil {
        panic(err)
    }

    // Subscribe to a topic
    sub, err := broker.Subscribe("test-topic")
    if err != nil {
        panic(err)
    }

    // Publish a message to the topic
    err = broker.Publish("test-topic", []byte("Hello, world!"))
    if err != nil {
        panic(err)
    }

    // Wait for the message to be received
    time.Sleep(time.Second)

    // Consume messages from the subscription
    for {
        msg, err := sub.Next()
        if err != nil {
            panic(err)
        }

        fmt.Printf("Received message: %s\n", msg.Payload())

        // Acknowledge the message
        err = msg.Ack()
        if err != nil {
            panic(err)
        }
    }
}
