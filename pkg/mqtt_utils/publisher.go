package mqtt_utils

import (
	"crypto/tls"
	"fmt"
	"log"

	devices "github.com/alienix2/sensor_info/pkg/devices/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	Device   devices.Device
	client   mqtt.Client
	username string
	password string
	topic    string
}

func NewPublisher(broker, topic, clientID string, device devices.Device, username, password string, tlsConfig *tls.Config) (*Publisher, error) {
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

	return &Publisher{
		client: client,
		topic:  topic,
		Device: device,
	}, nil
}

func (p *Publisher) SetTopic(topic string) {
	p.topic = topic
}

func (p *Publisher) Publish() error {
	data, err := p.Device.FormatData()
	if err != nil {
		return fmt.Errorf("failed to read sensor data: %w", err)
	}

	token := p.client.Publish(p.topic, 1, true, data)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed to publish message: %w", token.Error())
	}

	log.Printf("Published message: %s to topic: %s\n", data, p.topic)
	return nil
}

func (p *Publisher) Disconnect() {
	p.client.Disconnect(250)
	log.Println("Publisher disconnected.")
}

func (p *Publisher) GetDevice() devices.Device {
	return p.Device
}
