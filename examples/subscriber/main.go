package main

import (
	"flag"
	"fmt"
	"log"

	mqtt_utils "github.com/alienix2/sensor_info/pkg/mqtt_utils"
	storage "github.com/alienix2/sensor_info/pkg/storage/devices_database"
	tls_config "github.com/alienix2/sensor_info/pkg/tls_config"
	"github.com/google/uuid"
)

func main() {
	defaultUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	topic := flag.String("topic", "home/temperature", "MQTT topic to subscribe to")
	username := flag.String("username", "", "MQTT username")
	password := flag.String("password", "", "MQTT password")
	clientID := flag.String("clientID", defaultUUID.String(), "Client ID for the subscriber")
	database_path := flag.String("database_path", "./sqlite/subscribers.db", "Path to the SQLite database")
	flag.Parse()

	handler := &mqtt_utils.VerifyDataMessageHandler[storage.DeviceData]{
		DatabaseMessageHandler: &mqtt_utils.DatabaseMessageHandler[storage.DeviceData]{
			SaveFunc: storage.SaveJsonToSQLite[storage.DeviceData],
		},
		ActionCallback: func(data storage.DeviceData) {
			if data.DeviceData > 10 {
				fmt.Println("Be careful, data is greater than 10, maybe you should do something!")
			}
		},
	}

	tlsConfig, err := tls_config.LoadCertificates("certifications/subscriber.crt", "certifications/subscriber.key", "certifications/ca.crt")
	subscriber, err := mqtt_utils.NewSubscriber(
		*brokerURL,
		*topic,
		*clientID,
		handler,
		*username,
		*password,
		tlsConfig,
	)

	storage.InitSQLiteDatabase(*database_path, &storage.DeviceData{})

	if err != nil {
		log.Fatalf("Failed to create MQTT subscriber: %v", err)
	}
	defer subscriber.Disconnect()

	err = subscriber.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	subscriber.Wait()
}
