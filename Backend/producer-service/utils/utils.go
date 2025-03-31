package utils

import "math"

func RoundToTwoDecimalPlaces(val float64) float64 {

	// for 2 decimal places
	ratio := math.Pow(10, 2)

	return math.Round(val*ratio) / ratio

}
