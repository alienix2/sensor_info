package mqtt_utils

import (
	"encoding/json"
	"log"

	central_storage "github.com/alienix2/sensor_info/pkg/storage/central_database"
	sensors_storage "github.com/alienix2/sensor_info/pkg/storage/devices_database"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
		SentAt:      deviceData.Timestamp,
		Topic:       topic,
		DeviceName:  deviceData.Name,
		DeviceUnit:  deviceData.Unit,
		DeviceID:    deviceData.DeviceID,
		DeviceData:  deviceData.DeviceData,
		ControlData: deviceData.ControlData,
	}

	central_storage.SaveMessageToMySQL(messageData)
}
