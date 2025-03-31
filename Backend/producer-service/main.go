package main

import (
	"producer-service/sensors"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		sensors.EnvironSensors()

	}()

	go func() {
		defer wg.Done()
		sensors.WaterSoilSensors()

	}()

	wg.Wait()

}
