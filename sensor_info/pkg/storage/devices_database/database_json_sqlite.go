package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitSQLiteDatabase(dsn string, models ...interface{}) {
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database initialized successfully.")
}

func SaveJsonToSQLite[T any](data T) error {
	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Device data saved: %+v\n", data)
	return nil
}

func GetAllData[T any](out *[]T) error {
	result := db.Find(out)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
