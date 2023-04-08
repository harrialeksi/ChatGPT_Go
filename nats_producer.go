/*
shows nats producer example in go language

In this example, we first connect to a NATS server using the nats.Connect function and passing in the server address as a string. We defer the close of 
the connection to ensure it is properly closed when the function returns.

We then define a subject and a message to send. The subject is a string that identifies the type of message being sent, and the message is a string 
containing the data to be sent.

We publish the message using the nc.Publish method of the connection object, passing in the subject and message as parameters. If an error occurs 
during publishing, we panic with the error message.

Finally, we wait for the message to be delivered by calling nc.Flush and sleeping for 100 milliseconds. Note that this is not always necessary, 
as NATS uses an asynchronous delivery model. However, it is useful in some cases to ensure that the message has been delivered before the program exits.
*/

package main

import (
    "fmt"
    "time"

    "github.com/nats-io/nats.go"
)

func main() {
    // Connect to NATS server
    nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        panic(err)
    }
    defer nc.Close()

    // Define subject and message to send
    subject := "my.subject"
    message := "Hello, world!"

    // Publish message
    err = nc.Publish(subject, []byte(message))
    if err != nil {
        panic(err)
    }

    fmt.Printf("Published message '%s' to subject '%s'\n", message, subject)

    // Wait for the message to be delivered
    nc.Flush()
    time.Sleep(100 * time.Millisecond)
}
