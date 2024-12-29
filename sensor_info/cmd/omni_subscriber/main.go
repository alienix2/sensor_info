package main

import (
	"flag"
	"log"

	"mattemoni.sensor_info/pkg/mqtt_utils"
	storage "mattemoni.sensor_info/pkg/storage/central_database"
)

func main() {
	handler := &mqtt_utils.CentralDatabaseMessageHandler{}
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	topic := flag.String("topic", "#", "MQTT topic to subscribe to")
	username := flag.String("username", "omnisub", "Omnisub username")
	password := flag.String("password", "password", "Omnisub password")
	database_path := flag.String("database_path", "mqtt_admin:Panzerotto@tcp(localhost:3306)/mqtt_users?parseTime=true", "Path to the SQLite database")
	clientID := flag.String("clientID", "generic_subscriber", "Client ID for the subscriber")
	flag.Parse()

	subscriber, err := mqtt_utils.NewSubscriber(
		*brokerURL,
		"certifications/subscriber.crt",
		"certifications/subscriber.key",
		"certifications/ca.crt",
		*topic,
		*clientID,
		handler,
		*username,
		*password)

	storage.InitMySQLCentralDatabase(*database_path)

	if err != nil {
		log.Fatalf("Failed to create MQTT subscriber: %v", err)
	}
	defer subscriber.Disconnect()

	// Subscribe to the topic
	err = subscriber.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	subscriber.Wait()
}
