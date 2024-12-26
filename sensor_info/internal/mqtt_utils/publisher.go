package mqtt_utils

import (
	"crypto/tls"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	sensors "mattemoni.sensor_info/internal/sensors"
	tlsconfig "mattemoni.sensor_info/internal/tls_config"
)

type Publisher struct {
	client    mqtt.Client
	clientID  string
	tlsConfig *tls.Config
	topic     string
	sensor    sensors.Sensor
}

func NewPublisher(brokerURL, certPath, keyPath, caPath, topic, clientID string, sensor sensors.Sensor) (*Publisher, error) {
	tlsConfig, err := tlsconfig.LoadTLSConfig(certPath, keyPath, caPath)
	if err != nil {
		return nil, fmt.Errorf("failed to configure TLS: %w", err)
	}

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID(clientID).
		SetTLSConfig(tlsConfig)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	return &Publisher{
		client:    client,
		topic:     topic,
		clientID:  clientID,
		tlsConfig: tlsConfig,
		sensor:    sensor,
	}, nil
}

func (p *Publisher) Publish() error {
	data, err := p.sensor.FormatData()
	if err != nil {
		return fmt.Errorf("failed to read sensor data: %w", err)
	}

	token := p.client.Publish(p.topic, 1, true, data)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish message: %w", token.Error())
	}
	fmt.Printf("Published message: %s\n", data)
	return nil
}

func (p *Publisher) Disconnect() {
	p.client.Disconnect(250)
	fmt.Println("Publisher disconnected.")
}
