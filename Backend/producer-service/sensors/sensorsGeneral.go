package sensors

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"math"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/segmentio/kafka-go"
// )

// type sensorMessage struct {
// 	SensorID  int       `json:"sensorid"`
// 	Timestamp time.Time `json:"timestamp"`
// 	Value     float64   `json:"value"`
// }

// func roundToTwoDecimalPlaces(val float64) float64 {

// 	// for 2 decimal places
// 	ratio := math.Pow(10, 2)

// 	return math.Round(val*ratio) / ratio

// }

// func SensorsGeneral(topic string) {

// 	broker := os.Getenv("KAFKA_BROKER")

// 	_, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	writer := &kafka.Writer{

// 		Addr:     kafka.TCP(broker),
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	}

// 	defer writer.Close()

// 	log.Printf("%s producer started!! ", topic)

// 	time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))

// 	msg := &sensorMessage{}

// 	for {

// 		var val float64
// 		// Random humidity between 0 and 100
// 		switch topic {
// 		case "humidity":
// 			val = roundToTwoDecimalPlaces(rand.Float64() * 100)
// 		case "temperature":
// 			val = roundToTwoDecimalPlaces(rand.Float64()*50 + 5)
// 		case "ph":
// 			val = roundToTwoDecimalPlaces(rand.Float64()*14 + 1)
// 		}

// 		// msg := SensorMessage{
// 		// 	Timestamp: time.Now(),
// 		// 	Value:     val,
// 		// }

// 		sensorID := rand.Intn(10)

// 		msg.Timestamp = time.Now()
// 		msg.Value = val
// 		msg.SensorID = sensorID

// 		msgbytes, err := json.Marshal(msg)

// 		if err != nil {
// 			log.Printf("Error marshalling message: %v", err)
// 			continue
// 		}

// 		err = writer.WriteMessages(context.TODO(), kafka.Message{
// 			Key:   []byte(strconv.Itoa(sensorID)),
// 			Value: msgbytes,
// 		})

// 		if err != nil {
// 			log.Printf("Error producing %s message", topic)
// 		} else {
// 			log.Printf("%s Sent: %s", topic, msgbytes)
// 		}

// 		time.Sleep(time.Second * time.Duration(rand.Intn(15)+1))
// 	}

// }
