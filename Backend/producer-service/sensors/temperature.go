package sensors

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func Temperature() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := "temperature"

	_, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
	if err != nil {
		panic(err.Error())
	}

	writer := &kafka.Writer{

		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	log.Println("Temperature producer started!! ")

	for {
		temp := rand.Float64()*50 + 5 // Random temperature between 5 and 55
		tempBytes := []byte(strconv.FormatFloat(temp, 'f', 2, 64))

		err := writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   []byte(strconv.Itoa(rand.Intn(10))),
			Value: tempBytes,
		})

		if err != nil {
			log.Printf("Error producing TEMPERATURE message %s", err)
		} else {
			log.Printf("temp Sent: %s", tempBytes)
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))

	}

}
