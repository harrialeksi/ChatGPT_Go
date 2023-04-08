/*
shows nats consumer example in go language

In this example, we first connect to a NATS server using the nats.Connect function and passing in the server address as a string. We defer the 
close of the connection to ensure it is properly closed when the function returns.

We then subscribe to a subject using the nc.QueueSubscribe method of the connection object. We pass in the subject and a queue name as parameters, 
as well as a callback function to be called when a message is received. The callback function takes a single parameter of type *nats.Msg, which 
contains the message data and metadata.

In this example, the callback function simply prints the received message and subject to the console.

We then wait for a SIGINT or SIGTERM signal to exit. When the signal is received, we log a message and exit. Note that this example uses a channel 
to wait for the signal, but you can use any method that works for your use case.
*/

package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/nats-io/nats.go"
)

func main() {
    // Connect to NATS server
    nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        panic(err)
    }
    defer nc.Close()

    // Subscribe to subject
    subject := "my.subject"
    queueName := "my.queue"
    _, err = nc.QueueSubscribe(subject, queueName, func(msg *nats.Msg) {
        fmt.Printf("Received message '%s' from subject '%s'\n", string(msg.Data), msg.Subject)
    })
    if err != nil {
        panic(err)
    }

    // Wait for SIGINT or SIGTERM to exit
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    <-sigCh

    log.Println("Shutting down...")
}
