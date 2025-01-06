package mqtt_utils

import (
	"crypto/tls"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	client         mqtt.Client
	messageHandler MessageHandlerStrategy
	topic          string
}

func NewSubscriber(broker, topic, clientID string, handler MessageHandlerStrategy, username, password string, tlsConfig *tls.Config) (*Subscriber, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID)

	if tlsConfig != nil {
		opts.SetTLSConfig(tlsConfig)
	}

	if username != "" && password != "" {
		opts.SetUsername(username)
		opts.SetPassword(password)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	return &Subscriber{
		client:         client,
		topic:          topic,
		messageHandler: handler,
	}, nil
}

func (s *Subscriber) Subscribe() error {
	token := s.client.Subscribe(s.topic, 1, s.messageHandler.HandleReceive)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to subscribe to topic: %v", token.Error())
	}

	log.Println("Subscribed to topic:", s.topic)
	return nil
}

func (s *Subscriber) Wait() {
	select {}
}

func (s *Subscriber) Disconnect() {
	s.client.Disconnect(250)
	log.Println("Subscriber disconnected")
}

func (s *Subscriber) SetMessageHandler(handler MessageHandlerStrategy) {
	s.messageHandler = handler
}
