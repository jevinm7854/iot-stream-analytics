package main

import (
	"producer-service/sensors"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		sensors.SensorsGeneral("temperature")

	}()

	go func() {
		defer wg.Done()
		sensors.SensorsGeneral("humidity")

	}()

	go func() {
		defer wg.Done()
		sensors.SensorsGeneral("ph")

	}()

	wg.Wait()

}
