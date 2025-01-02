package mqtt_utils

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DataActionCallback[T any] func(data T)

type VerifyDataMessageHandler[T any] struct {
	DatabaseMessageHandler *DatabaseMessageHandler[T]
	ActionCallback         DataActionCallback[T]
}

func (v *VerifyDataMessageHandler[T]) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	var data T

	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		return
	}

	if v.ActionCallback != nil {
		v.ActionCallback(data)
	}

	if v.DatabaseMessageHandler != nil {
		v.DatabaseMessageHandler.HandleReceive(client, msg)
	}
}

func (v *VerifyDataMessageHandler[T]) SetActionCallback(callback DataActionCallback[T]) {
	v.ActionCallback = callback
}
