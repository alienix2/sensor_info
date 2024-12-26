package mqtt_utils

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	tlsconfig "mattemoni.sensor_info/internal/tls_config"
)

type Subscriber struct {
	client         mqtt.Client
	messageHandler MessageHandlerStrategy
	topic          string
}

func NewSubscriber(broker, certFile, keyFile, caFile, topic string, handler MessageHandlerStrategy, username, password string) (*Subscriber, error) {
	tlsConfig, err := tlsconfig.LoadTLSConfig(certFile, keyFile, caFile)
	if err != nil {
		return nil, fmt.Errorf("failed to configure TLS: %v", err)
	}

	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID("generic_subscriber").
		SetTLSConfig(tlsConfig)
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

	fmt.Println("Subscribed to topic:", s.topic)
	return nil
}

func (s *Subscriber) Wait() {
	select {}
}

func (s *Subscriber) Disconnect() {
	s.client.Disconnect(250)
	fmt.Println("Subscriber disconnected")
}
