package mqtt_utils

import (
	"testing"

	mocks "mattemoni.sensor_info/pkg/mqtt_utils/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCommandRegistry(t *testing.T) {
	tests := []struct {
		expectedResults map[string]*mocks.MockCommand
		operations      func(registry CommandRegistry)
		name            string
	}{
		{
			name: "Register and retrieve multiple commands",
			operations: func(registry CommandRegistry) {
				command1 := &mocks.MockCommand{Name: "Command1"}
				command2 := &mocks.MockCommand{Name: "Command2"}

				registry.RegisterCommand("command1", command1)
				registry.RegisterCommand("command2", command2)
			},
			expectedResults: map[string]*mocks.MockCommand{
				"command1": {Name: "Command1"},
				"command2": {Name: "Command2"},
			},
		},
		{
			name: "Register, unregister, and verify remaining commands",
			operations: func(registry CommandRegistry) {
				command1 := &mocks.MockCommand{Name: "Command1"}
				command2 := &mocks.MockCommand{Name: "Command2"}

				registry.RegisterCommand("command1", command1)
				registry.RegisterCommand("command2", command2)
				registry.UnregisterCommand("command1")
			},
			expectedResults: map[string]*mocks.MockCommand{
				"command2": {Name: "Command2"},
			},
		},
		{
			name: "Unregister non-existent command",
			operations: func(registry CommandRegistry) {
				command1 := &mocks.MockCommand{Name: "Command1"}

				registry.RegisterCommand("command1", command1)
				registry.UnregisterCommand("nonexistent")
			},
			expectedResults: map[string]*mocks.MockCommand{
				"command1": {Name: "Command1"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			registry := *NewCommandRegistry()

			tt.operations(registry)

			actualCommands := registry.GetCommands()
			assert.Equal(t, len(tt.expectedResults), len(actualCommands))

			for key, expectedCommand := range tt.expectedResults {
				actualCommand, exists := actualCommands[key]
				assert.True(t, exists, "Command %s should exist", key)

				assert.Equal(t, expectedCommand, actualCommand, "Command objects should match")
			}
		})
	}
}
