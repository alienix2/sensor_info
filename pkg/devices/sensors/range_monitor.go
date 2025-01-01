package devices

import (
	"log"
)

func (s *Sensor) CheckValueInRange() (bool, error) {
	data, err := s.FormatData()
	if err != nil {
		log.Printf("Error reading sensor data: %v", err)
		return false, err
	}

	sensorValue, err := s.ParseDeviceValue(data)
	if err != nil {
		log.Printf("Error parsing sensor value: %v", err)
		return false, err
	}

	minValue, maxValue := s.GetRange()
	if sensorValue < minValue || sensorValue > maxValue {
		return false, nil
	}

	thresholdMargin := (maxValue - minValue) * 0.1

	if sensorValue <= minValue+thresholdMargin {
		log.Printf("Warning: Sensor value %.2f is close to the minimum (%.2f)!\n", sensorValue, minValue)
		return true, nil
	} else if sensorValue >= maxValue-thresholdMargin {
		log.Printf("Warning: Sensor value %.2f is close to the maximum (%.2f)!\n", sensorValue, maxValue)
		return true, nil
	}

	return false, nil
}
