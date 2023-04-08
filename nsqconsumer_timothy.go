//https://levelup.gitconnected.com/implementing-messaging-queue-nsq-in-golang-using-docker-99b402293b12

package main
import (
 "encoding/json"
 "github.com/nsqio/go-nsq"
 "log"
 "os"
 "os/signal"
 "syscall"
 "time"
)
type messageHandler struct{}
type Message struct {
 Name      string
 Content   string
 Timestamp string
}
func main() {
 //The only valid way to instantiate the Config
 config := nsq.NewConfig()
//Tweak several common setup in config
 // Maximum number of times this consumer will attempt to process a message before giving up
 config.MaxAttempts = 10
// Maximum number of messages to allow in flight
 config.MaxInFlight = 5
// Maximum duration when REQueueing
 config.MaxRequeueDelay = time.Second * 900
 config.DefaultRequeueDelay = time.Second * 0
//Init topic name and channel
 topic := "Topic_Example"
 channel := "Channel_Example"
//Creating the consumer
 consumer, err := nsq.NewConsumer(topic, channel, config)
 if err != nil {
  log.Fatal(err)
 }
// Set the Handler for messages received by this Consumer.
 consumer.AddHandler(&messageHandler{})
//Use nsqlookupd to find nsqd instances
 consumer.ConnectToNSQLookupd("127.0.0.1:4161")
// wait for signal to exit
 sigChan := make(chan os.Signal, 1)
 signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
 <-sigChan
// Gracefully stop the consumer.
 consumer.Stop()
}
// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
//Process the Message
 var request Message
 if err := json.Unmarshal(m.Body, &request); err != nil {
  log.Println("Error when Unmarshaling the message body, Err : ", err)
  // Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
  return err
 }
//Print the Message
 log.Println("Message")
 log.Println("--------------------")
 log.Println("Name : ", request.Name)
 log.Println("Content : ", request.Content)
 log.Println("Timestamp : ", request.Timestamp)
 log.Println("--------------------")
 log.Println("")
// Will automatically set the message as finish
 return nil
}
