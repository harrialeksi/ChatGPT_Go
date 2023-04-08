//https://andrewrgoss.com/2017/getting-started-with-nsq-for-go/

package main

import (
    "log"
    "sync"

    "github.com/nsqio/go-nsq"
)

func main() {
    wg := &sync.WaitGroup{}
    wg.Add(1)

    decodeConfig := nsq.NewConfig()
    c, err := nsq.NewConsumer("My_NSQ_Topic", "My_NSQ_Channel", decodeConfig)
    if err != nil {
        log.Panic("Could not create consumer")
    }
    //c.MaxInFlight defaults to 1

    c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
        log.Println("NSQ message received:")
        log.Println(string(message.Body))
        return nil
    }))

    err = c.ConnectToNSQD("127.0.0.1:4150")
    if err != nil {
        log.Panic("Could not connect")
    }
    log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
    wg.Wait()
}
