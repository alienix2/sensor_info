package mqtt_utils

import (
	"testing"

	"github.com/alienix2/sensor_info/pkg/mqtt_utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewSubscriber(t *testing.T) {
	clientID := "myTestClientID"
	topic := "test/topic"
	username := ""
	password := ""
	handler := &mocks.MockMessageHandler{}

	subscriber, err := NewSubscriber("tcp://"+port, topic, clientID, handler, username, password, nil)

	assert.Nil(t, err, "Expected no error while creating subscriber")
	assert.NotNil(t, subscriber, "Expected subscriber to be non-nil")
	assert.Equal(t, topic, subscriber.topic, "Expected topic to match the given topic")

	err = subscriber.Subscribe()
	assert.Nil(t, err, "Expected no error while subscribing")

	subscriber.Disconnect()

	err = subscriber.Subscribe()
	assert.NotNil(t, err, "Expected error while subscribing after disconnecting")
}
