package mqtt_utils

import (
	"fmt"
	"log"
	"net"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	server "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/stretchr/testify/assert"
)

var client_port string

func GetAvailablePort() (string, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", fmt.Errorf("failed to find available port: %v", err)
	}
	defer listener.Close()
	int_port := listener.Addr().(*net.TCPAddr).Port
	return fmt.Sprintf("localhost:%d", int_port), nil
}

func StartMockMQTTServer(port string) (*server.Server, error) {
	server := server.New(nil)
	tcp := listeners.NewTCP(listeners.Config{ID: "t1", Address: port})
	_ = server.AddHook(new(auth.AllowHook), nil)
	err := server.AddListener(tcp)
	if err != nil {
		return nil, fmt.Errorf("failed to create mock MQTT server: %v", err)
	}

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("Failed to start mock MQTT server: %v", err)
		}
	}()

	client_port = port

	return server, nil
}

func StopMockMQTTServer(mqttServer *server.Server) {
	if err := mqttServer.Close(); err != nil {
		log.Fatalf("Failed to stop mock MQTT server: %v", err)
	}
}

func createMQTTClient(t *testing.T) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + client_port)
	opts.SetClientID("testClient")

	client := mqtt.NewClient(opts)
	token := client.Connect()
	assert.True(t, token.Wait())
	assert.NoError(t, token.Error())

	t.Cleanup(func() { client.Disconnect(250) })
	return client
}
