package mocks

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MockMessageHandler struct {
	ReceivedClient    mqtt.Client
	ReceivedMessage   mqtt.Message
	HandleReceiveFunc func(client mqtt.Client, msg mqtt.Message)
}

func (m *MockMessageHandler) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	m.ReceivedClient = client
	m.ReceivedMessage = msg

	if m.HandleReceiveFunc != nil {
		m.HandleReceiveFunc(client, msg)
	}
}
