package mqtt_utils

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	central_storage "mattemoni.sensor_info/internal/storage/central_database"
	sensors_storage "mattemoni.sensor_info/internal/storage/devices_database"
)

type CentralDatabaseMessageHandler struct{}

func (c *CentralDatabaseMessageHandler) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	var sensorData sensors_storage.SensorData
	if err := json.Unmarshal(msg.Payload(), &sensorData); err != nil {
		log.Printf("Error parsing JSON data: %v", err)
		return
	}
	topic := msg.Topic()

	messageData := central_storage.MessageData{
		SentAt:     sensorData.Timestamp,
		Topic:      topic,
		SensorName: sensorData.Name,
		SensorUnit: sensorData.Unit,
		SensorID:   sensorData.SensorID,
		SensorData: sensorData.SensorData,
	}

	central_storage.SaveMessageToMySQL(messageData)
}
