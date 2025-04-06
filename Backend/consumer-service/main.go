package main

import (
	"consumer-service/consumer/environ"
	"consumer-service/consumer/watersoil"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		environ.EnvironConsumer()

	}()

	go func() {
		defer wg.Done()

		watersoil.WaterSoilConsumer()

	}()

	wg.Wait()

}
