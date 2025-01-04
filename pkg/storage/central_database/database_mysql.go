package storage

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MessageData struct {
	SentAt     time.Time `gorm:"not null" json:"timestamp"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
	Topic      string    `gorm:"size:255" json:"topic"`
	DeviceName string    `gorm:"size:255" json:"name"`
	DeviceUnit string    `gorm:"size:255" json:"unit"`
	DeviceID   string    `gorm:"not null" json:"id"`
	Notes      string    `gorm:"size:255" json:"notes"`
	DeviceData float64   `gorm:"not null" json:"device_data"`
	ID         int       `gorm:"primaryKey" json:"-"`
}

var db *gorm.DB

func InitMySQLCentralDatabase(dsn string) error {
	var err error
	log.Println("Connecting to database...")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return err
	}

	log.Println("Starting AutoMigrate...")
	err = db.AutoMigrate(&MessageData{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
		return err
	}
	log.Println("Mysql database initialized successfully.")
	return nil
}

func SaveMessageToMySQL(data MessageData) error {
	log.Printf("Message data to save: %+v\n", data)
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Message data saved: %+v\n", data)
	return nil
}

func GetAllData() ([]MessageData, error) {
	var messages []MessageData
	result := db.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}
