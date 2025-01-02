package devices

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "github.com/alienix2/sensor_info/pkg/devices/common/mocks"
)

// Defined here to avoid import cycle
type MockReader struct {
	Err   error
	Value float64
}

func (m *MockReader) Read(s *Sensor) (float64, error) {
	return m.Value, m.Err
}

func TestSensor_CheckValueInRange(t *testing.T) {
	tests := []struct {
		mockReader      *MockReader
		mockFormatter   *mocks.MockFormatter
		name            string
		minValue        float64
		maxValue        float64
		expectedInRange bool
		expectErr       bool
	}{
		{
			name: "value within range",
			mockReader: &MockReader{
				Value: 46.0,
				Err:   nil,
			},
			mockFormatter: &mocks.MockFormatter{
				FormattedData: "formatted_data",
				FormatErr:     nil,
				ParsedValue:   46.0,
				ParseErr:      nil,
			},
			minValue:        10.0,
			maxValue:        50.0,
			expectedInRange: true,
			expectErr:       false,
		},
		{
			name: "value out of range",
			mockReader: &MockReader{
				Value: 55.0,
				Err:   nil,
			},
			mockFormatter: &mocks.MockFormatter{
				FormattedData: "formatted_data",
				FormatErr:     nil,
				ParsedValue:   55.0,
				ParseErr:      nil,
			},
			minValue:        10.0,
			maxValue:        50.0,
			expectedInRange: false,
			expectErr:       false,
		},
		{
			name: "formatter error",
			mockReader: &MockReader{
				Value: 45.0,
				Err:   nil,
			},
			mockFormatter: &mocks.MockFormatter{
				FormatErr: errors.New("formatter error"),
			},
			minValue:        10.0,
			maxValue:        50.0,
			expectedInRange: false,
			expectErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the sensor with mocks
			sensor := NewSensor(
				WithSensorRange(tt.minValue, tt.maxValue),
				WithReaderStrategy(tt.mockReader),
				WithSensorFormatterStrategy(tt.mockFormatter),
			)

			inRange, err := sensor.CheckValueInRange()

			// Use assertions
			assert.Equal(t, tt.expectErr, err != nil)
			assert.Equal(t, tt.expectedInRange, inRange)
		})
	}
}
