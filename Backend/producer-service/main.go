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
		sensors.Temperature()

	}()

	go func() {
		defer wg.Done()
		sensors.Humidity()

	}()

	go func() {
		defer wg.Done()
		sensors.Ph()

	}()

	wg.Wait()

}
