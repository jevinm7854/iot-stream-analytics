package sensors

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"producer-service/utils"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type waterSoilSensorMessage struct {
	ControllerID int       `json:"controllerid"`
	Timestamp    time.Time `json:"timestamp"`
	Ph           float64   `json:"ph"`
	Turbidity    float64   `json:"turbidity"`
	SoilMoisture float64   `json:"soilmoisture"`
}

func WaterSoilSensors() {

	broker := os.Getenv("KAFKA_BROKER")
	topic := "waterSoilSensors"

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

	log.Printf("%s producer started!! ", topic)

	time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))

	msg := &waterSoilSensorMessage{}

	for {

		// msg := SensorMessage{
		// 	Timestamp: time.Now(),
		// 	Value:     val,
		// }

		ControllerID := rand.Intn(5) + 5 // 0-4 for environ sensors. 5 - 9 for water and soil sensors
		msg.ControllerID = ControllerID
		msg.Timestamp = time.Now()
		msg.Ph = utils.RoundToTwoDecimalPlaces(rand.Float64() * 14)            // 0 - 14
		msg.SoilMoisture = utils.RoundToTwoDecimalPlaces(rand.Float64() * 100) // 0 -100
		msg.Turbidity = utils.RoundToTwoDecimalPlaces(rand.Float64() * 1000)   // 0 - 1000

		msgbytes, err := json.Marshal(msg)

		if err != nil {
			log.Printf("Error marshalling message: %v", err)
			continue
		}

		err = writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   []byte(strconv.Itoa(ControllerID)),
			Value: msgbytes,
		})

		if err != nil {
			log.Printf("Error producing %s message", topic)
		} else {
			log.Printf("%s Sent: %s", topic, msgbytes)
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))
	}

}
