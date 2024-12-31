package storage_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	mariadb "github.com/testcontainers/testcontainers-go/modules/mariadb"
	"github.com/testcontainers/testcontainers-go/wait"
	storage "mattemoni.sensor_info/pkg/storage/central_database"
)

var (
	mariadbContainer testcontainers.Container
	mariadbHost      string
	mariadbPort      nat.Port
	ctx              = context.Background()
)

func setupContainer(ctx context.Context) (string, error) {
	log.Println("Starting container setup...")
	var err error
	mariadbContainer, err := mariadb.Run(ctx,
		"mariadb:11.0.3",
		mariadb.WithDatabase("test_db"),
		mariadb.WithUsername("root"),
		mariadb.WithPassword("password"),
		testcontainers.WithWaitStrategy(wait.ForLog("ready for connections").WithOccurrence(2)),
	)
	if err != nil {
		return "", err
	}

	mariadbHost, err := mariadbContainer.Host(ctx)
	if err != nil {
		return "", err
	}

	mariadbPort, err := mariadbContainer.MappedPort(ctx, "3306")
	if err != nil {
		return "", err
	}

	dsn := fmt.Sprintf("root:password@tcp(%s:%s)/test_db?charset=utf8mb4&parseTime=True&loc=Local", mariadbHost, mariadbPort.Port())
	log.Printf("MariaDB connection string: %s", dsn)
	return dsn, nil
}

func terminateMariaDBContainer(ctx context.Context, mariadbContainer testcontainers.Container) {
	if mariadbContainer != nil {
		if err := mariadbContainer.Terminate(ctx); err != nil {
			log.Printf("Failed to terminate container: %s", err)
		}
	}
}

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dsn, err := setupContainer(ctx)
	if err != nil {
		log.Fatalf("Error setting up container: %s", err)
	}

	if err := storage.InitMySQLCentralDatabase(dsn); err != nil {
		log.Fatalf("Failed to initialize database: %s", err)
	}

	defer terminateMariaDBContainer(ctx, mariadbContainer)

	exitcode := m.Run()
	os.Exit(exitcode)
}

func TestSaveMessageToMySQL(t *testing.T) {
	message := storage.MessageData{
		SentAt:     time.Now(),
		Topic:      "temperature",
		DeviceName: "Sensor1",
		DeviceUnit: "Celsius",
		DeviceID:   "12345",
		DeviceData: 23.5,
	}

	err := storage.SaveMessageToMySQL(message)
	assert.Nil(t, err, "Failed to save message to MySQL")

	data, err := storage.GetAllData()
	assert.Nil(t, err, "Failed to retrieve all data")

	var retrievedData storage.MessageData
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
	log.Printf("mariadbHost: %s", mariadbHost)
	data, err := storage.GetAllData()
	assert.Nil(t, err, "Error fetching all data")

	assert.Greater(t, len(data), 0, "Expected data, but found none")
}
