package mqtt_utils

import (
	"crypto/tls"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	devices "mattemoni.sensor_info/internal/devices/common"
	tlsconfig "mattemoni.sensor_info/internal/tls_config"
)

type Publisher struct {
	device    devices.Device
	client    mqtt.Client
	tlsConfig *tls.Config
	clientID  string
	username  string
	password  string
	topic     string
}

func NewPublisher(brokerURL, certPath, keyPath, caPath, topic, clientID string, device devices.Device, username, password string) (*Publisher, error) {
	tlsConfig, err := tlsconfig.LoadTLSConfig(certPath, keyPath, caPath)
	if err != nil {
		return nil, fmt.Errorf("failed to configure TLS: %w", err)
	}

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID(clientID).
		SetTLSConfig(tlsConfig)
	if username != "" && password != "" {
		opts.SetUsername(username)
		opts.SetPassword(password)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	return &Publisher{
		client:    client,
		topic:     topic,
		clientID:  clientID,
		tlsConfig: tlsConfig,
		device:    device,
	}, nil
}

func (p *Publisher) Publish() error {
	data, err := p.device.FormatData()
	if err != nil {
		return fmt.Errorf("failed to read sensor data: %w", err)
	}

	token := p.client.Publish(p.topic, 0, true, data)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish message: %w", token.Error())
	}

	fmt.Printf("Published message: %s to topic: %s\n", data, p.topic)
	return nil
}

func (p *Publisher) Disconnect() {
	p.client.Disconnect(250)
	fmt.Println("Publisher disconnected.")
}
