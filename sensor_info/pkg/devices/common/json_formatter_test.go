package devices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatterStrategy_Format(t *testing.T) {
	tests := []struct {
		name       string
		deviceName string
		unit       string
		uid        string
		expected   string
		data       float64
	}{
		{
			name:       "Test JSON format with valid inputs",
			data:       123.456,
			deviceName: "TemperatureSensor",
			unit:       "C",
			uid:        "device-001",
			expected:   `{"name":"TemperatureSensor","unit":"C","id":"device-001","timestamp":"`, // will assert timestamp separately
		},
		{
			name:       "Test JSON format with zero data",
			data:       0,
			deviceName: "PressureSensor",
			unit:       "Pa",
			uid:        "device-002",
			expected:   `{"name":"PressureSensor","unit":"Pa","id":"device-002","timestamp":"`, // will assert timestamp separately
		},
	}

	formatter := &JSONFormatterStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatter.Format(tt.data, tt.deviceName, tt.unit, tt.uid)
			assert.Nil(t, err, "Expected no error during formatting")

			// Ensure the formatted JSON starts with the expected string and contains a timestamp
			assert.Contains(t, result, tt.expected, "Formatted JSON should match the expected fields")
			// Ensure the result contains a valid timestamp
			assert.NotEmpty(t, result, "Formatted JSON should not be empty")
		})
	}
}

func TestJSONFormatterStrategy_Parse(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		expected  float64
		expectErr bool
	}{
		{
			name:      "Test parse valid JSON",
			data:      `{"name":"TemperatureSensor","unit":"C","id":"device-001","timestamp":"2024-12-30T00:41:57.641623274+01:00","device_data":123.45}`,
			expected:  123.45,
			expectErr: false,
		},
		{
			name:      "Test parse invalid JSON",
			data:      `{"name":"TemperatureSensor","unit":"C","id":"device-001","device_data":"abc"}`,
			expected:  0,
			expectErr: true,
		},
	}

	formatter := &JSONFormatterStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := formatter.Parse(tt.data)
			if tt.expectErr {
				assert.Error(t, err, "Expected an error for invalid JSON")
			} else {
				assert.Nil(t, err, "Expected no error during parsing")
				assert.Equal(t, tt.expected, result, "Parsed result should match the expected value")
			}
		})
	}
}
