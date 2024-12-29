package devices

import (
	"testing"

	mocks "mattemoni.sensor_info/pkg/devices/common/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCommands(t *testing.T) {
	tests := []struct {
		name           string
		command        Command
		initialStatus  string
		expectedStatus string
	}{
		{
			name: "TurnOnCommand sets status to 'on'",
			command: &TurnOnCommand{
				Device: newMockDevice("off"),
			},
			initialStatus:  "off",
			expectedStatus: "on",
		},
		{
			name: "TurnOffCommand sets status to 'off'",
			command: &TurnOffCommand{
				Device: newMockDevice("on"),
			},
			initialStatus:  "on",
			expectedStatus: "off",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockDevice *mocks.MockDevice
			if _, ok := tt.command.(*TurnOffCommand); ok {
				mockDevice = tt.command.(*TurnOffCommand).Device.(*mocks.MockDevice)
			} else {
				mockDevice = tt.command.(*TurnOnCommand).Device.(*mocks.MockDevice)
			}

			mockDevice.Status = tt.initialStatus
			err := tt.command.Execute()

			assert.Equal(t, tt.expectedStatus, mockDevice.GetStatus(), "Device status should match the expected value")
			assert.Nil(t, err, "Error should be nil")
		})
	}
}

func newMockDevice(initialStatus string) *mocks.MockDevice {
	return &mocks.MockDevice{
		ID:     "123",
		Name:   "Mock Device",
		Status: initialStatus,
	}
}
