package mqtt_utils

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MessageHandlerStrategy interface {
	HandleReceive(client mqtt.Client, msg mqtt.Message)
}

type PrinterMessageHandler struct{}

func (h *PrinterMessageHandler) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
