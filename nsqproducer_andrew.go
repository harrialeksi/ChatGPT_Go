//https://andrewrgoss.com/2017/getting-started-with-nsq-for-go/

package main

import (
  "log"

  "github.com/nsqio/go-nsq"
)

func main() {
    config := nsq.NewConfig()
    p, err := nsq.NewProducer("127.0.0.1:4150", config)
    if err != nil {
        log.Panic(err)
    }
    err = p.Publish("My_NSQ_Topic", []byte("sample NSQ message"))
    if err != nil {
        log.Panic(err)
    }
}
