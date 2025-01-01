package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testDB string

func TestMain(m *testing.M) {
	testDB = ":memory:"
	InitSQLiteDatabase(testDB, DeviceData{}, ControlData{})

	m.Run()
}

func TestSaveAndGetControlData(t *testing.T) {
	controlData := ControlData{
		Timestamp: time.Now(),
		Command:   "START",
	}
	err := SaveJsonToSQLite(controlData)
	assert.NoError(t, err, "Failed to save control data")

	var retrievedData []ControlData
	err = GetAllData(&retrievedData)
	assert.NoError(t, err, "Failed to retrieve control data")
	assert.Len(t, retrievedData, 1, "Expected 1 control data record")
	assert.Equal(t, controlData.Command, retrievedData[0].Command, "Command mismatch")
}

func TestSaveAndGetDeviceData(t *testing.T) {
	deviceData := DeviceData{
		Timestamp:  time.Now(),
		Name:       "Device1",
		Unit:       "Celsius",
		DeviceID:   "1234",
		DeviceData: 23.5,
	}
	err := SaveJsonToSQLite(deviceData)
	assert.NoError(t, err, "Failed to save device data")

	var retrievedData []DeviceData
	err = GetAllData(&retrievedData)
	assert.NoError(t, err, "Failed to retrieve device data")
	assert.Len(t, retrievedData, 1, "Expected 1 device data record")
	assert.Equal(t, deviceData.DeviceID, retrievedData[0].DeviceID, "Device ID mismatch")
}
