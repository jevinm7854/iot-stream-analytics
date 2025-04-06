package watersoil

import (
	"consumer-service/models"
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func waterSoilProducer(s *models.WaterSoilSensorMessage) {

	broker := os.Getenv("KAFKA_BROKER")
	topic := "waterSoilSensorsClean"

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

	msgbytes, err := json.Marshal(s)

	if err != nil {
		log.Printf("Error clean marshalling message: %v", err)

	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(strconv.Itoa(s.ControllerID)),
		Value: msgbytes,
	})

	if err != nil {
		log.Printf("Error producing %s message", topic)
	} else {
		log.Printf("%s Sent: %s", topic, msgbytes)
	}

}
