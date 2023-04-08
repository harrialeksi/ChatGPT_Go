//https://levelup.gitconnected.com/implementing-messaging-queue-nsq-in-golang-using-docker-99b402293b12

package main
import (
 "encoding/json"
 "github.com/nsqio/go-nsq"
 "log"
 "time"
)
type Message struct {
 Name      string
 Content   string
 Timestamp string
}
func main() {
//The only valid way to instantiate the Config
 config := nsq.NewConfig()
//Creating the Producer using NSQD Address
 producer, err := nsq.NewProducer("127.0.0.1:4150", config)
 if err != nil {
  log.Fatal(err)
 }
//Init topic name and message
 topic := "Topic_Example"
 msg := Message{
  Name:      "Message Name Example",
  Content:   "Message Content Example",
  Timestamp: time.Now().String(),
 }
//Convert message as []byte
 payload, err := json.Marshal(msg)
 if err != nil {
  log.Println(err)
 }
//Publish the Message
 err = producer.Publish(topic, payload)
 if err != nil {
  log.Println(err)
 }
}
