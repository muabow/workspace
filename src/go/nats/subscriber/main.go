package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nats/model"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

var pid int = 0
var r_msg_cnt int = 0

func processMsg(m *nats.Msg) {
	r_msg_cnt++
	pl := &model.Payload{}
	json.Unmarshal(m.Data, pl)
	fmt.Printf("Receive Message = [%v/%v] on subject =[%v], total receive count = %v\n", pl.Count, pl.Data, m.Subject, r_msg_cnt)
	replayData := fmt.Sprintf("ack message # pid(%v)", pid)
	m.Respond([]byte(replayData))
}

func main() {

	pid = os.Getpid()

	fmt.Println("##############")
	fmt.Printf(" [Subscriber] (%v)\n", pid)
	fmt.Println("##############")

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Drain()

	sub, err := nc.QueueSubscribe("intros", "zip1", processMsg)
	if err != nil {
		log.Fatalf("can't subscribe to NATS queue 'zip1': %v", err)
	}
	defer sub.Unsubscribe()

	time.Sleep(1 * time.Hour)

	fmt.Println("Process completed!!\n")
}

