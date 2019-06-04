package main

import (
	"log"

	"github.com/nats-io/nats.go"
	server "github.com/nats-io/nats-server/v2/test"
)

func main() {
	server.RunServer(nil)

	nc, err := nats.Connect("127.0.0.1", nats.Name("NATS Client"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected!")
	nc.Subscribe("hi", func(m *nats.Msg) {
		log.Println("[Received] ", string(m.Data))
	})
	nc.Publish("hi", []byte("Hello NATS!"))

	select {}
}
