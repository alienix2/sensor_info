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
	SensorName string    `gorm:"size:255" json:"name"`
	SensorUnit string    `gorm:"size:255" json:"unit"`
	SensorID   string    `gorm:"not null" json:"id"`
	SensorData float64   `gorm:"not null" json:"sensor_data"`
	ID         int       `gorm:"primaryKey" json:"-"`
}

var db *gorm.DB

func InitMySQLCentralDatabase(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = db.AutoMigrate(&MessageData{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}

func SaveMessageToMySQL(data MessageData) error {
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Message data saved: %+v\n", data)
	return nil
}
