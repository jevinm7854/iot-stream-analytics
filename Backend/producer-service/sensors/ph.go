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

func Ph() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := "ph"

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

	log.Println("ph producer started!! ")

	for {
		ph := rand.Float64() * 14 // Random humidity between 0 and 100
		phBytes := []byte(strconv.FormatFloat(ph, 'f', 2, 64))

		err := writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   []byte(strconv.Itoa(rand.Intn(10))),
			Value: phBytes,
		})

		if err != nil {
			log.Printf("Error producing PH message")
		} else {
			log.Printf("PH Sent: %s", phBytes)
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))

	}

}
