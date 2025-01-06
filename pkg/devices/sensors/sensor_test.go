package devices

import (
	"testing"

	"github.com/alienix2/sensor_info/pkg/devices/common/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSensor(t *testing.T) {
	mockReader := &MockReader{
		Value: 42.0,
		Err:   nil,
	}
	mockFormatter := &mocks.MockFormatter{
		FormattedData: "formatted_data",
		FormatErr:     nil,
		ParsedValue:   42.0,
		ParseErr:      nil,
	}

	sensor := NewSensor(
		WithSensorName("Test Sensor"),
		WithSensorRange(0, 100),
		WithReaderStrategy(mockReader),
		WithSensorFormatterStrategy(mockFormatter),
	)

	assert.Equal(t, "Test Sensor", sensor.GetName(), "Sensor name should be 'Test Sensor'")
	assert.Equal(t, "on", sensor.GetStatus(), "Sensor default status should be 'on'")

	minValue, maxValue := sensor.GetRange()
	assert.Equal(t, 0.0, minValue, "Sensor min value should be 0.0")
	assert.Equal(t, 100.0, maxValue, "Sensor max value should be 100.0")

	sensor.SetStatus("off")
	assert.Equal(t, "off", sensor.GetStatus(), "Sensor status should be 'off' after setting")

	formattedData, err := sensor.FormatData()
	assert.NoError(t, err, "Expected no error while formatting data")
	assert.Equal(t, "formatted_data", formattedData, "Formatted data should be 'formatted_data'")

	parsedValue, err := sensor.ParseDeviceValue("formatted_data")
	assert.NoError(t, err, "Expected no error while parsing device value")
	assert.Equal(t, 42.0, parsedValue, "Parsed value should be 42.0")
}

func TestSensor_GenerateID(t *testing.T) {
	sensor1 := NewSensor()
	sensor2 := NewSensor()

	assert.NotEqual(t, sensor1.GetID(), sensor2.GetID(), "Sensors should have different IDs")
}
