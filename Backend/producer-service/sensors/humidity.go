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

func Humidity() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := "humidity"

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

	log.Println("Humidity producer started!! ")

	for {
		hum := rand.Float64() * 100 // Random humidity between 0 and 100
		humBytes := []byte(strconv.FormatFloat(hum, 'f', 2, 64))

		err := writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   []byte(strconv.Itoa(rand.Intn(10))),
			Value: humBytes,
		})

		if err != nil {
			log.Printf("Error producing HUMIDITY message")
		} else {
			log.Printf("Hum Sent: %s", humBytes)
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))

	}

}
