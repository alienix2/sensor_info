package devices

import (
	"fmt"
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
		return false, err // Return false and the error if the value cannot be parsed
	}

	minValue, maxValue := s.GetRange()
	thresholdMargin := (maxValue - minValue) * 0.1 // 10% margin

	if sensorValue <= minValue+thresholdMargin {
		fmt.Printf("Warning: Sensor value %.2f is close to the minimum (%.2f)!\n", sensorValue, minValue)
		return true, nil
	} else if sensorValue >= maxValue-thresholdMargin {
		fmt.Printf("Warning: Sensor value %.2f is close to the maximum (%.2f)!\n", sensorValue, maxValue)
		return true, nil
	}

	return false, nil
}
