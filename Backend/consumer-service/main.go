package main

import (
	"consumer-service/consumer/environ"
	"consumer-service/consumer/watersoil"
)

func main() {

	environ.EnvironConsumer()
	watersoil.WaterSoilConsumer()
}
