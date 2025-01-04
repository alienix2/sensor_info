package mqtt_utils

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	central_storage "github.com/alienix2/sensor_info/pkg/storage/central_database"
	storage "github.com/alienix2/sensor_info/pkg/storage/devices_database"
	"github.com/stretchr/testify/assert"
)

func TestCentralDatabaseMessageHandler(t *testing.T) {
	ctx := context.Background()

	dsn, err := central_storage.SetupContainer(ctx)
	assert.NoError(t, err)
	defer central_storage.TerminateMariaDBContainer(ctx)

	err = central_storage.InitMySQLCentralDatabase(dsn)
	assert.NoError(t, err)

	client := createMQTTClient(t)

	handler := &CentralDatabaseMessageHandler{}
	topic := "test/central_device_data"

	client.Subscribe(topic, 0, handler.HandleReceive)

	testData := storage.DeviceData{
		Timestamp:  time.Now(),
		Name:       "Device1",
		Unit:       "Celsius",
		DeviceID:   "1234",
		DeviceData: 23.5,
		Notes:      "Temperature sensor",
	}
	payload, err := json.Marshal(testData)
	assert.NoError(t, err)

	token := client.Publish(topic, 0, false, payload)
	assert.True(t, token.Wait())
	assert.NoError(t, token.Error())

	time.Sleep(2 * time.Second)

	savedMessages, err := central_storage.GetAllData()
	assert.NoError(t, err)

	var savedMessage central_storage.MessageData
	for _, msg := range savedMessages {
		if msg.DeviceID == testData.DeviceID && msg.Topic == topic {
			savedMessage = msg
			break
		}
	}

	expectedSentAt := testData.Timestamp.Truncate(time.Millisecond)
	actualSentAt := savedMessage.SentAt.Truncate(time.Millisecond)

	assert.Equal(t, expectedSentAt, actualSentAt)
	assert.Equal(t, testData.Name, savedMessage.DeviceName)
	assert.Equal(t, testData.Unit, savedMessage.DeviceUnit)
	assert.Equal(t, testData.DeviceID, savedMessage.DeviceID)
	assert.Equal(t, testData.DeviceData, savedMessage.DeviceData)
}
