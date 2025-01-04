package storage

import "time"

type ControlData struct {
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
	Command   string    `gorm:"size:100;not null" json:"command"`
	ID        uint      `gorm:"primaryKey" json:"-"`
}

type DeviceData struct {
	Timestamp   time.Time `gorm:"not null" json:"timestamp"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Unit        string    `gorm:"size:50;not null" json:"unit"`
	DeviceID    string    `gorm:"size:100;not null" json:"id"`
	ControlData string    `gorm:"size:100;not null" json:"control_data"`
	DeviceData  float64   `gorm:"not null" json:"device_data"`
	ID          uint      `gorm:"primaryKey" json:"-"`
}
