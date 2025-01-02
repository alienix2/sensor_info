package mqtt_utils

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/alienix2/sensor_info/pkg/mqtt_utils/mocks"
)

func TestControlMessageHandlerWithRegisteredCommand(t *testing.T) {
	client := createMQTTClient(t)

	commandRegistry := NewCommandRegistry()

	mockCommandExecuted := &mocks.MockCommand{
		Name: "test_command_1",
	}

	commandRegistry.RegisterCommand(mockCommandExecuted.Name, mockCommandExecuted)

	handler := &ControlMessageHandler[map[string]interface{}]{
		CommandRegistry: commandRegistry,
		CommandKey:      "command",
	}

	topic := "test/topic"
	client.Subscribe(topic, 0, handler.HandleReceive)

	payload := map[string]interface{}{
		"command": "test_command_1",
	}
	payloadBytes, err := json.Marshal(payload)
	assert.Nil(t, err)

	token := client.Publish(topic, 0, false, payloadBytes)

	time.Sleep(2 * time.Second)
	mockCommandExecuted.Mu.Lock()
	assert.True(t, token.Wait())
	assert.Nil(t, token.Error())
	assert.True(t, mockCommandExecuted.Executed)
	mockCommandExecuted.Mu.Unlock()

	client.Disconnect(250)
}

func TestControlMessageHandlerWithUnregisteredCommand(t *testing.T) {
	client := createMQTTClient(t)

	commandRegistry := NewCommandRegistry()

	mockCommandNotExecuted := &mocks.MockCommand{}

	handler := &ControlMessageHandler[map[string]interface{}]{
		CommandRegistry: commandRegistry,
		CommandKey:      "command",
	}

	topic := "test/topic"
	client.Subscribe(topic, 0, handler.HandleReceive)

	payload := map[string]interface{}{
		"command": "test_command_2",
	}
	payloadBytes, err := json.Marshal(payload)
	assert.Nil(t, err)

	token := client.Publish(topic, 0, false, payloadBytes)
	time.Sleep(2 * time.Second)
	assert.True(t, token.Wait())
	assert.Nil(t, token.Error())
	assert.False(t, mockCommandNotExecuted.Executed)

	client.Disconnect(250)
}
