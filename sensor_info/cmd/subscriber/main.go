package main

import (
	"flag"
	"log"

	"mattemoni.sensor_info/internal/mqtt_utils"
	storage "mattemoni.sensor_info/internal/storage/sensors"
)

func main() {
	// Create and start the client
	handler := &mqtt_utils.DatabaseMessageHandler{}
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	topic := flag.String("topic", "home/temperature", "MQTT topic to subscribe to")
	username := flag.String("username", "", "MQTT username")
	password := flag.String("password", "", "MQTT password")
	flag.Parse()

	subscriber, err := mqtt_utils.NewSubscriber(
		*brokerURL,
		"certifications/subscriber.crt",
		"certifications/subscriber.key",
		"certifications/ca.crt",
		*topic,
		handler,
		*username,
		*password)

	storage.InitSQLiteDatabase("./sqlite/subscribers")

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
