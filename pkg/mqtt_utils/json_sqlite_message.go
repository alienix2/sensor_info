package mqtt_utils

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DatabaseSaveFunc[T any] func(data T) error

type DatabaseMessageHandler[T any] struct {
	SaveFunc DatabaseSaveFunc[T]
}

func (h *DatabaseMessageHandler[T]) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	var data T
	if err := json.Unmarshal(msg.Payload(), &data); err != nil {
		log.Printf("Error parsing JSON data: %v", err)
		return
	}

	if err := h.SaveFunc(data); err != nil {
		log.Printf("Error saving data to SQLite: %v", err)
		return
	}

	log.Println("Data stored successfully!")
}
