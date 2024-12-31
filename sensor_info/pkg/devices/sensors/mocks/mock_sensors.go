package mocks

import "errors"

type MockSensor struct {
	FormatErr   error
	ParseErr    error
	Data        string
	MinValue    float64
	MaxValue    float64
	ParsedValue float64
}

func (m *MockSensor) FormatData() (string, error) {
	return m.Data, m.FormatErr
}

func (m *MockSensor) ParseDeviceValue(data string) (float64, error) {
	if data == m.Data {
		return m.ParsedValue, m.ParseErr
	}
	return 0, errors.New("unexpected data")
}

func (m *MockSensor) GetRange() (float64, float64) {
	return m.MinValue, m.MaxValue
}
