/*
show me an nsq consumer that saves data to postgres database

In this example, we create a PostgreSQL database connection using the database/sql package and the pq driver. We then create an NSQ consumer and add a 
message handler that parses the message body into a Message struct and inserts it into the users table in the database using a prepared SQL statement. 
We also handle any unique key violation errors by ignoring duplicate messages. Finally, we connect the consumer to the NSQD instance and wait for messages 
to be processed.

Note that in production code, you should handle errors and implement appropriate error handling, such as logging or retrying failed database inserts.
*/

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/nsqio/go-nsq"
)

type Message struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create a new PostgreSQL database connection
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new NSQ consumer
	consumer, err := nsq.NewConsumer("topic", "channel", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Add a handler that saves messages to the database
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		// Parse the message body into a Message struct
		var m Message
		err := json.Unmarshal(message.Body, &m)
		if err != nil {
			log.Println("Error parsing message body:", err)
			return err
		}

		// Insert the message into the database
		_, err = db.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", m.ID, m.Name, m.Email)
		if err != nil {
			// If the error is a unique key violation, it means we already saved this message, so we can ignore it
			if e, ok := err.(*pq.Error); ok && e.Code.Name() == "unique_violation" {
				log.Println("Duplicate message:", err)
				return nil
			}
			log.Println("Error inserting message:", err)
			return err
		}

		log.Println("Message saved:", m)
		return nil
	}))

	// Connect the consumer to the NSQD instance
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for messages to be processed
	select {}
}
