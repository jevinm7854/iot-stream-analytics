package models

func (s *EnvironSensorMessage) IsValid() bool {

	if s.ControllerID < 0 || s.ControllerID > 5 {
		return false
	}

	if s.Temperature < -10 || s.Temperature > 40 { // limits
		return false
	}
	if s.Humidity < 0 || s.Humidity > 100 {
		return false
	}

	if s.Pressure < 980 || s.Pressure > 1080 {
		return false
	}

	if s.CO2 < 350 || s.CO2 > 900 {
		return false
	}

	return true
}

func (s *WaterSoilSensorMessage) IsValid() bool {

	if s.Ph < 0 || s.Ph > 14 {
		return false
	}

	if s.SoilMoisture < 0 || s.SoilMoisture > 100 {
		return false
	}

	if s.Turbidity < 0 || s.Turbidity > 1000 {
		return false
	}

	return true

}
