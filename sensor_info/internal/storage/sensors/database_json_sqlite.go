package storage

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SensorData struct {
	Timestamp  time.Time `gorm:"not null"`
	Name       string    `gorm:"size:100;not null"`
	Unit       string    `gorm:"size:50;not null"`
	SensorID   string    `gorm:"size:100;not null" json:"ID"`
	SensorData float64   `gorm:"not null"`
	ID         uint      `gorm:"primaryKey" json:"-"`
}

var db *gorm.DB

func InitSQLiteDatabase(databasePath string) {
	var err error
	db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&SensorData{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	log.Println("Database initialized successfully.")
}

func SaveJsonToSQLite(data SensorData) error {
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
