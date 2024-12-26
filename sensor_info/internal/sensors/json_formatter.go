package sensors

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONFormatterStrategy struct{}

type SensorData struct {
	Name       string  `json:"name"`
	Unit       string  `json:"unit"`
	ID         string  `json:"id"`
	Timestamp  string  `json:"timestamp"`
	SensorData float64 `json:"sensor_data"`
}

func (j *JSONFormatterStrategy) Format(data float64, name, unit, id string) (string, error) {
	sensorData := SensorData{
		SensorData: data,
		Name:       name,
		Unit:       unit,
		ID:         id,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(sensorData)
	if err != nil {
		return "", fmt.Errorf("failed to format data as JSON: %w", err)
	}
	return string(jsonData), nil
}

func (j *JSONFormatterStrategy) Parse(data string) (float64, error) {
	var sensorData SensorData
	err := json.Unmarshal([]byte(data), &sensorData)
	if err != nil {
		return 0, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return sensorData.SensorData, nil
}
