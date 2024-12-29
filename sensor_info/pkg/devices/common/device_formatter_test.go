package devices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawFormatterStrategy_Format(t *testing.T) {
	tests := []struct {
		name       string
		deviceName string
		unit       string
		uid        string
		expected   string
		data       float64
	}{
		{
			name:       "Test format with valid inputs",
			data:       123.456,
			deviceName: "TemperatureSensor",
			unit:       "C",
			uid:        "device-001",
			expected:   "123.46, TemperatureSensor, C, device-001", // .2f formatting for 123.456
		},
		{
			name:       "Test format with zero data",
			data:       0,
			deviceName: "PressureSensor",
			unit:       "Pa",
			uid:        "device-002",
			expected:   "0.00, PressureSensor, Pa, device-002",
		},
	}

	rfs := &RawFormatterStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := rfs.Format(tt.data, tt.deviceName, tt.unit, tt.uid)
			assert.NoError(t, err, "Expected no error during formatting")
			assert.Equal(t, tt.expected, result, "Formatted result should match the expected value")
		})
	}
}

func TestRawFormatterStrategy_Parse(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		expected  float64
		expectErr bool
	}{
		{
			name:      "Test parse valid float",
			data:      "123.45",
			expected:  123.45,
			expectErr: false,
		},
		{
			name:      "Test parse invalid float",
			data:      "abc",
			expected:  0,
			expectErr: true,
		},
	}

	rfs := &RawFormatterStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := rfs.Parse(tt.data)
			if tt.expectErr {
				assert.Error(t, err, "Expected an error for invalid parse")
			} else {
				assert.NoError(t, err, "Expected no error during parse")
				assert.Equal(t, tt.expected, result, "Parsed result should match the expected value")
			}
		})
	}
}
