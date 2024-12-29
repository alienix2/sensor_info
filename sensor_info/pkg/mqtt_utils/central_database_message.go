package mqtt_utils

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	central_storage "mattemoni.sensor_info/pkg/storage/central_database"
	sensors_storage "mattemoni.sensor_info/pkg/storage/devices_database"
)

type CentralDatabaseMessageHandler struct{}

func (c *CentralDatabaseMessageHandler) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	var deviceData sensors_storage.DeviceData
	if err := json.Unmarshal(msg.Payload(), &deviceData); err != nil {
		log.Printf("Error parsing JSON data: %v", err)
		return
	}
	topic := msg.Topic()

	messageData := central_storage.MessageData{
		SentAt:     deviceData.Timestamp,
		Topic:      topic,
		DeviceName: deviceData.Name,
		DeviceUnit: deviceData.Unit,
		DeviceID:   deviceData.DeviceID,
		DeviceData: deviceData.DeviceData,
	}

	central_storage.SaveMessageToMySQL(messageData)
}
