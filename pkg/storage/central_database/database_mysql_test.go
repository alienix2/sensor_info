package storage

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	dsn, err := SetupContainer(ctx)
	if err != nil {
		log.Fatalf("Error setting up container: %s", err)
	}

	if err := InitMySQLCentralDatabase(dsn); err != nil {
		log.Fatalf("Failed to initialize database: %s", err)
	}

	defer TerminateMariaDBContainer(ctx)

	exitcode := m.Run()
	os.Exit(exitcode)
}

func TestSaveMessageToMySQL(t *testing.T) {
	message := MessageData{
		SentAt:     time.Now(),
		Topic:      "temperature",
		DeviceName: "Sensor1",
		DeviceUnit: "Celsius",
		DeviceID:   "12345",
		DeviceData: 23.5,
	}

	err := SaveMessageToMySQL(message)
	assert.Nil(t, err, "Failed to save message to MySQL")

	data, err := GetAllData()
	assert.Nil(t, err, "Failed to retrieve all data")

	var retrievedData MessageData
	for _, d := range data {
		if d.DeviceID == message.DeviceID {
			retrievedData = d
			break
		}
	}

	assert.Equal(t, message.DeviceID, retrievedData.DeviceID, "Device ID mismatch")
	assert.Equal(t, message.DeviceData, retrievedData.DeviceData, "Device data mismatch")
}

func TestGetAllData(t *testing.T) {
	data, err := GetAllData()
	assert.Nil(t, err, "Error fetching all data")

	assert.Greater(t, len(data), 0, "Expected data, but found none")
}
