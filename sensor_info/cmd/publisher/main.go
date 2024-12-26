package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"mattemoni.sensor_info/internal/mqtt_utils"
	"mattemoni.sensor_info/internal/sensors"
)

func main() {
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	interval := flag.Int("interval", 5, "Publish interval in seconds")
	flag.Parse()

	sensor := sensors.NewSensor(
		sensors.WithName("RandomSensor"),
		sensors.WithUnit("Units"),
		sensors.WithRange(0, 20),
		sensors.WithReaderStrategy(&sensors.DefaultReader{}),
		sensors.WithFormatterStrategy(&sensors.JSONFormatterStrategy{}),
	)

	publisher, err := mqtt_utils.NewPublisher(
		*brokerURL,
		"certifications/subscriber.crt",
		"certifications/subscriber.key",
		"certifications/ca.crt",
		"home/temperature",
		"generic_publisher",
		*sensor,
	)
	if err != nil {
		log.Fatalf("Failed to initialize publisher: %v", err)
	}
	defer publisher.Disconnect()

	go publishWarningMessages(publisher, sensor)
	publishRegularly(publisher, time.Duration(*interval)*time.Second) // Publish every 'interval' seconds}
}

func publishRegularly(publisher *mqtt_utils.Publisher, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		publishValue(publisher)
	}
}

func publishValue(publisher *mqtt_utils.Publisher) {
	err := publisher.Publish()
	if err != nil {
		log.Printf("Error publishing sensor data: %v", err)
	} else {
		fmt.Println("Sensor data published successfully.")
	}
}

func publishWarningMessages(publisher *mqtt_utils.Publisher, sensor *sensors.Sensor) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		warning, err := sensor.CheckValueInRange()
		if err != nil {
			log.Printf("Error monitoring sensor: %v", err)
		}
		if warning {
			err := publisher.Publish()
			if err != nil {
				log.Printf("Error publishing warning message: %v", err)
			} else {
				fmt.Println("Warning message published successfully.")
			}
		}
	}
}
