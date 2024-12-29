package mqtt_utils

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ControlMessageHandler[T any] struct {
	DatabaseMessageHandler *DatabaseMessageHandler[T]
	CommandRegistry        *CommandRegistry
	CommandKey             string
}

func (h *ControlMessageHandler[T]) HandleReceive(client mqtt.Client, msg mqtt.Message) {
	if h.DatabaseMessageHandler != nil {
		h.DatabaseMessageHandler.HandleReceive(client, msg)
	}

	commandKey := h.CommandKey
	if commandKey == "" {
		commandKey = "command"
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
		log.Printf("Error parsing payload: %v", err)
		return
	}

	commandName, ok := payload[commandKey].(string)
	if !ok {
		log.Printf("Invalid or missing '%s' field in payload", commandKey)
		return
	}

	if command, exists := h.CommandRegistry.getCommand(commandName); exists {
		if err := command.Execute(); err != nil {
			log.Printf("Error executing command '%s': %v", commandName, err)
		} else {
			log.Printf("Command '%s' executed successfully", commandName)
		}
	} else {
		log.Printf("Unknown command '%s'", commandName)
	}
}
