package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nats/model"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	pl := &model.Payload{}

	fmt.Println("##############")
	fmt.Println(" [Publisher] ")
	fmt.Println("##############")
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Drain()

	count := 0
	for {
		count++

		pl.Count = count
		pl.Data = fmt.Sprintf("(%d)th greeting Hello World!!", count)

		data, _ := json.Marshal(pl)

		reply, err := nc.Request("intros", data, 500*time.Millisecond)
		if err != nil {
			log.Printf("error sending message count = %v, err: %v\n", pl.Count, err)
		} else {
			fmt.Printf("Sending (%d)th message = [%s], ", pl.Count, pl.Data)
			fmt.Printf("Receive reply = [%s]\n", string(reply.Data))
		}
		time.Sleep(1 * time.Millisecond)
	}
}
