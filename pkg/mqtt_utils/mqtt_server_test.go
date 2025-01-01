package mqtt_utils

import (
	"fmt"
	"log"
	"net"

	server "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func getAvailablePort() (string, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", fmt.Errorf("failed to find available port: %v", err)
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	return fmt.Sprintf("localhost:%d", port), nil
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

	return server, nil
}

func StopMockMQTTServer(mqttServer *server.Server) {
	if err := mqttServer.Close(); err != nil {
		log.Fatalf("Failed to stop mock MQTT server: %v", err)
	}
}
