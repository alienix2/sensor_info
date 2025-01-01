package mqtt_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	server "github.com/mochi-mqtt/server/v2"
	"github.com/stretchr/testify/assert"
	storage "mattemoni.sensor_info/pkg/storage/devices_database"
)

var (
	mqttServer *server.Server
	port       string
)

func TestMain(m *testing.M) {
	var err error

	port, err = getAvailablePort()
	if err != nil {
		log.Fatalf("Error finding available port: %v", err)
	}

	mqttServer, err = StartMockMQTTServer(port)
	if err != nil {
		log.Fatalf("Error starting mock MQTT server: %v", err)
	}

	code := m.Run()

	StopMockMQTTServer(mqttServer)

	fmt.Printf("Exiting with code %d\n", code)
}

func createMQTTClient(t *testing.T) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker("tcp://" + port)
	opts.SetClientID("testClient")

	client := mqtt.NewClient(opts)
	token := client.Connect()
	assert.True(t, token.Wait())
	assert.NoError(t, token.Error())

	t.Cleanup(func() { client.Disconnect(250) })
	return client
}

func MockDatabaseSaveFunc[T any](data T) error {
	log.Printf("Mock save called with data: %+v", data)
	return nil
}

func TestPrinterMessageHandler(t *testing.T) {
	client := createMQTTClient(t)

	printerHandler := &PrinterMessageHandler{}
	topicPrinter := "test/print_topic"

	client.Subscribe(topicPrinter, 0, printerHandler.HandleReceive)

	printPayload := []byte("Hello, World!")
	token := client.Publish(topicPrinter, 0, false, printPayload)
	assert.True(t, token.Wait())
	assert.NoError(t, token.Error())
}

func TestDatabaseMessageHandler(t *testing.T) {
	client := createMQTTClient(t)

	dbHandler := DatabaseMessageHandler[storage.DeviceData]{
		SaveFunc: MockDatabaseSaveFunc[storage.DeviceData],
	}
	topicDB := "test/device_data"

	client.Subscribe(topicDB, 0, dbHandler.HandleReceive)

	testData := storage.DeviceData{
		Timestamp:  time.Now(),
		Name:       "Device1",
		Unit:       "Celsius",
		DeviceID:   "1234",
		DeviceData: 23.5,
	}
	payload, err := json.Marshal(testData)
	assert.NoError(t, err)

	token := client.Publish(topicDB, 0, false, payload)
	assert.True(t, token.Wait())
	assert.NoError(t, token.Error())
}
