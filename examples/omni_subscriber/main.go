package main

import (
	"flag"
	"log"

	"github.com/alienix2/sensor_info/pkg/mqtt_utils"
	storage "github.com/alienix2/sensor_info/pkg/storage/central_database"
	tls_config "github.com/alienix2/sensor_info/pkg/tls_config"
	"github.com/google/uuid"
)

func main() {
	defaultUUID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}
	handler := &mqtt_utils.CentralDatabaseMessageHandler{}
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	topic := flag.String("topic", "#", "MQTT topic to subscribe to")
	username := flag.String("username", "omnisub", "Omnisub username")
	password := flag.String("password", "password", "Omnisub password")
	dsn := flag.String("dsn", "mqtt_admin:Panzerotto@tcp(localhost:3306)/mqtt_users?parseTime=true", "Path to MySQL database")
	clientID := flag.String("clientID", defaultUUID.String(), "Client ID for the subscriber")
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

	storage.InitMySQLCentralDatabase(*dsn)

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
