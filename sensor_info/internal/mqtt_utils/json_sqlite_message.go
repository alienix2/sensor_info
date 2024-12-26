package mqtt_utils

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	storage "mattemoni.sensor_info/internal/storage/sensors"
)

type DatabaseMessageHandler struct{}

func (h *DatabaseMessageHandler) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	var data storage.SensorData
	if err := json.Unmarshal(msg.Payload(), &data); err != nil {
		log.Printf("Error parsing JSON data: %v", err)
		return
	}

	if err := storage.SaveJsonToSQLite(data); err != nil {
		log.Printf("Error saving data to SQLite: %v", err)
		return
	}

	fmt.Println("Data stored successfully!")
}
