package mqtt_utils

import (
	"testing"

	"github.com/alienix2/sensor_info/pkg/devices/common/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewPublisher(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	port, err := getAvailablePort()
	if err != nil {
		t.Fatalf("Error finding available port: %v", err)
	}

	mqttServer, err := StartMockMQTTServer(port)
	if err != nil {
		t.Fatalf("Failed to start mock MQTT server: %v", err)
	}
	defer StopMockMQTTServer(mqttServer)

	clientID := "myTestClientID"
	topic := "test/topic"
	username := ""
	password := ""
	device := &mocks.MockDevice{}

	publisher, err := NewPublisher("tcp://"+port, topic, clientID, device, username, password, nil)

	assert.Nil(t, err, "Expected no error while creating publisher")
	assert.NotNil(t, publisher, "Expected publisher to be non-nil")
	assert.Equal(t, topic, publisher.topic, "Expected topic to match the given topic")

	err = publisher.Publish()
	assert.Nil(t, err, "Expected no error while publishing message")

	newTopic := "test/new_topic"
	publisher.SetTopic(newTopic)
	assert.Equal(t, newTopic, publisher.topic, "Expected topic to be updated to the new topic")

	err = publisher.Publish()
	assert.Nil(t, err, "Expected no error while publishing message to the new topic")

	publisher.Disconnect()

	err = publisher.Publish()
	assert.NotNil(t, err, "Expected error while publishing after disconnecting")
}
