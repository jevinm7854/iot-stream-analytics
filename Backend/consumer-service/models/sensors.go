package models

import "time"

type EnvironSensorMessage struct {
	ControllerID int       `json:"controllerid"`
	Timestamp    time.Time `json:"timestamp"`
	Temperature  float64   `json:"temperature"`
	Humidity     float64   `json:"humidity"`
	Pressure     float64   `json:"pressure"`
	CO2          float64   `json:"co2"`
}

type WaterSoilSensorMessage struct {
	ControllerID int       `json:"controllerid"`
	Timestamp    time.Time `json:"timestamp"`
	Ph           float64   `json:"ph"`
	Turbidity    float64   `json:"turbidity"`
	SoilMoisture float64   `json:"soilmoisture"`
}
