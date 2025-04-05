package watersoil

import (
	"consumer-service/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func WaterSoilConsumer() {

	broker := os.Getenv("KAFKA_BROKER")
	topic := "waterSoilSensors"
	groupID := "watersoil-sensors-group"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset,
	})

	defer reader.Close()

	log.Println("Environ sensors consumer started. Waiting for messages...")

	for {

		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		fmt.Printf("Received message: %s (from partition %d, offset %d)\n",
			string(msg.Value), msg.Partition, msg.Offset)

		var sensor models.WaterSoilSensorMessage

		if err := json.Unmarshal(msg.Value, &sensor); err != nil {
			log.Printf("Invalid JSON: %v", err)
			continue
		}

		if !sensor.IsValid() {
			log.Printf("Invalid sensor data: %+v", sensor)
			continue
		}

		fmt.Printf("Valid sensor data: %+v\n", sensor)

	}

}
