package storage

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ControlData struct {
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
	Command   string    `gorm:"size:100;not null" json:"command"`
	ID        uint      `gorm:"primaryKey" json:"-"`
}

type SensorData struct {
	Timestamp  time.Time `gorm:"not null" json:"timestamp"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Unit       string    `gorm:"size:50;not null" json:"unit"`
	SensorID   string    `gorm:"size:100;not null" json:"id"`
	SensorData float64   `gorm:"not null" json:"sensor_data"`
	ID         uint      `gorm:"primaryKey" json:"-"`
}

var db *gorm.DB

func InitSQLiteDatabase[T any](databasePath string, model T) {
	var err error
	db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(model)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	log.Println("Database initialized successfully.")
}

func SaveJsonToSQLite[T any](data T) error {
	fmt.Println(data)
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Sensor data saved: %+v\n", data)
	return nil
}

func GetAllSensorData() ([]SensorData, error) {
	var sensorData []SensorData
	result := db.Find(&sensorData)
	if result.Error != nil {
		return nil, result.Error
	}
	return sensorData, nil
}
