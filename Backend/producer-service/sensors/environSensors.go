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

type environSensorMessage struct {
	ControllerID int       `json:"controllerid"`
	Timestamp    time.Time `json:"timestamp"`
	Temperature  float64   `json:"temperature"`
	Humidity     float64   `json:"humidity"`
	Pressure     float64   `json:"pressure"`
	CO2          float64   `json:"co2"`
}

func EnvironSensors() {

	broker := os.Getenv("KAFKA_BROKER")
	topic := "environSensors"

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

	msg := &environSensorMessage{}

	for {

		// msg := SensorMessage{
		// 	Timestamp: time.Now(),
		// 	Value:     val,
		// }

		controllerID := rand.Intn(5) // 0-4 for environ sensors. 5 - 9 for water and soil sensors
		msg.ControllerID = controllerID
		msg.Timestamp = time.Now()
		msg.Temperature = utils.RoundToTwoDecimalPlaces(rand.Float64()*50 + 5) // 5- 55
		msg.Humidity = utils.RoundToTwoDecimalPlaces(rand.Float64() * 100)     // 0- 100
		msg.Pressure = utils.RoundToTwoDecimalPlaces(rand.Float64()*70 + 980)  //980-1050.0
		msg.CO2 = utils.RoundToTwoDecimalPlaces(rand.Float64()*550 + 350)      //350-900.0

		msgbytes, err := json.Marshal(msg)

		if err != nil {
			log.Printf("Error marshalling message: %v", err)
			continue
		}

		err = writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   []byte(strconv.Itoa(controllerID)),
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
