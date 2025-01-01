package devices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mattemoni.sensor_info/pkg/devices/common/mocks"
)

func TestActuator(t *testing.T) {
	mockFormatter := &mocks.MockFormatter{
		FormattedData: "formatted_data",
		FormatErr:     nil,
		ParseErr:      nil,
	}

	actuator := NewActuator(
		WithActuatorName("Test Actuator"),
		WithActuatorRange(0, 100),
		WithActuatorFormatterStrategy(mockFormatter),
	)

	assert.Equal(t, "Test Actuator", actuator.GetName(), "Expected actuator name to be 'Test Actuator'")
	assert.Equal(t, "idle", actuator.GetStatus(), "Expected default status to be 'idle'")

	minPower, maxPower := actuator.GetRange()
	assert.Equal(t, 0.0, minPower, "Expected minPower to be 0")
	assert.Equal(t, 100.0, maxPower, "Expected maxPower to be 100")

	actuator.SetStatus("active")
	assert.Equal(t, "active", actuator.GetStatus(), "Expected status to be 'active' after setting")

	formattedData, err := actuator.FormatData()
	assert.Nil(t, err, "Expected no error from FormatData")
	assert.Equal(t, "formatted_data", formattedData, "Expected formatted data to be 'formatted_data'")
}

func TestActuator_GenerateID(t *testing.T) {
	actuator1 := NewActuator()
	actuator2 := NewActuator()

	assert.NotEqual(t, actuator1.GetID(), actuator2.GetID(), "Expected different IDs for each actuator")
}

func TestActuator_DefaultValues(t *testing.T) {
	actuator := NewActuator()

	assert.Equal(t, "DefaultActuator", actuator.GetName(), "Expected default name to be 'DefaultActuator'")
	assert.Equal(t, "idle", actuator.GetStatus(), "Expected default status to be 'idle'")

	minPower, maxPower := actuator.GetRange()
	assert.Equal(t, 0.0, minPower, "Expected default minPower to be 0.0")
	assert.Equal(t, 100.0, maxPower, "Expected default maxPower to be 100.0")
}
