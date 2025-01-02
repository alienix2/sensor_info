package main

import (
	"flag"
	"log"

	mqtt_utils "github.com/alienix2/sensor_info/pkg/mqtt_utils"
	storage "github.com/alienix2/sensor_info/pkg/storage/devices_database"
	tls_config "github.com/alienix2/sensor_info/pkg/tls_config"
)

func main() {
	handler := &mqtt_utils.DatabaseMessageHandler[storage.DeviceData]{
		SaveFunc: storage.SaveJsonToSQLite[storage.DeviceData],
	}
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	topic := flag.String("topic", "home/temperature", "MQTT topic to subscribe to")
	username := flag.String("username", "", "MQTT username")
	password := flag.String("password", "", "MQTT password")
	clientID := flag.String("clientID", "generic_subscriber", "Client ID for the subscriber")
	database_path := flag.String("database_path", "./sqlite/subscribers.db", "Path to the SQLite database")
	flag.Parse()

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
