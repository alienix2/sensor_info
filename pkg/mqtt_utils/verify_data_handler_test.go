package mqtt_utils

import (
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
)

type MockData struct {
	SomeField string `json:"someField"`
}

func TestVerifyDataMessageHandler(t *testing.T) {
	client := createMQTTClient(t)

	callbackCalled := false
	actionCallback := func(data MockData) {
		callbackCalled = true
		assert.Equal(t, "test_data", data.SomeField)
	}

	handler := VerifyDataMessageHandler[MockData]{}
	handler.SetActionCallback(actionCallback)

	topic := "test/topic"
	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		handler.HandleReceive(client, msg)
	}); token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to subscribe: %v", token.Error())
	}

	client.Publish(topic, 0, false, `{"someField":"test_data"}`)

	time.Sleep(2 * time.Second)

	assert.True(t, callbackCalled)
}
