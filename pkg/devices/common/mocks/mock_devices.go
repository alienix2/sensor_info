package mocks

import "fmt"

type MockDevice struct {
	ID     string
	Name   string
	Status string
	Min    float64
	Max    float64
}

func (m *MockDevice) GetID() string {
	return m.ID
}

func (m *MockDevice) GetName() string {
	return m.Name
}

func (m *MockDevice) GetRange() (float64, float64) {
	return m.Min, m.Max
}

func (m *MockDevice) GetStatus() string {
	return m.Status
}

func (m *MockDevice) SetStatus(status string) {
	m.Status = status
}

func (m *MockDevice) FormatData() (string, error) {
	return fmt.Sprintf("Device ID: %s, Name: %s, Status: %s", m.ID, m.Name, m.Status), nil
}
