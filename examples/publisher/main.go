package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	common "github.com/alienix2/sensor_info/pkg/devices/common"
	sensors "github.com/alienix2/sensor_info/pkg/devices/sensors"
	mqtt_utils "github.com/alienix2/sensor_info/pkg/mqtt_utils"
	storage "github.com/alienix2/sensor_info/pkg/storage/devices_database"
	tls_config "github.com/alienix2/sensor_info/pkg/tls_config"
)

func main() {
	brokerURL := flag.String("broker", "tls://localhost:8883", "MQTT broker URL")
	interval := flag.Int("interval", 5, "Publish interval in seconds")
	username := flag.String("username", "", "MQTT username")
	password := flag.String("password", "", "MQTT password")
	topic := flag.String("topic", "home/temperature", "MQTT topic to publish to")
	rangeMin := flag.Float64("rangeMin", 0, "Minimum value for the sensor")
	rangeMax := flag.Float64("rangeMax", 20, "Maximum value for the sensor")
	unit := flag.String("unit", "Units", "Unit of the sensor")
	sensorName := flag.String("sensorName", "RandomSensor", "Name of the sensor")
	clientID := flag.String("clientID", "generic_publisher", "Client ID for the publisher")
	flag.Parse()

	sensor := sensors.NewSensor(
		sensors.WithSensorName(*sensorName),
		sensors.WithSensorUnit(*unit),
		sensors.WithSensorRange(*rangeMin, *rangeMax),
		sensors.WithReaderStrategy(&sensors.DefaultReader{}),
		sensors.WithSensorFormatterStrategy(&common.JSONFormatterStrategy{}),
	)

	// actuator := actuators.NewActuator(
	// 	actuators.WithActuatorName("Actuator"),
	// 	actuators.WithActuatorID("actuator_autogenerated-"+uuid.New().String()),
	// 	actuators.WithActuatorRange(0, 100),
	// 	actuators.WithActuatorUnit("unit"),
	// 	actuators.WithActuatorFormatterStrategy(&common.JSONFormatterStrategy{}),
	// )

	tlsConfig, err := tls_config.LoadCertificates("certifications/publisher.crt", "certifications/publisher.key", "certifications/ca.crt")
	publisher, err := mqtt_utils.NewPublisher(
		*brokerURL,
		*topic,
		*clientID,
		sensor,
		*username,
		*password,
		tlsConfig,
	)
	if err != nil {
		log.Fatalf("Failed to initialize publisher: %v", err)
	}
	defer publisher.Disconnect()

	storage.InitSQLiteDatabase("./sqlite/commands.db", &storage.ControlData{})
	commandRegistry := mqtt_utils.NewCommandRegistry()
	commandRegistry.RegisterCommand("turn_on", &common.TurnOnCommand{Device: sensor})
	commandRegistry.RegisterCommand("turn_off", &common.TurnOffCommand{Device: sensor})
	controlHandler := &mqtt_utils.ControlMessageHandler[storage.ControlData]{
		CommandRegistry: commandRegistry,
		CommandKey:      "command",
		DatabaseMessageHandler: &mqtt_utils.DatabaseMessageHandler[storage.ControlData]{
			SaveFunc: storage.SaveJsonToSQLite[storage.ControlData],
		},
	}

	tlsConfig, err = tls_config.LoadCertificates("certifications/subscriber.crt", "certifications/subscriber.key", "certifications/ca.crt")
	subscriber, err := mqtt_utils.NewSubscriber(
		*brokerURL,
		"command/"+*topic,
		*clientID+"_command",
		controlHandler,
		*username,
		*password,
		tlsConfig,
	)
	if err != nil {
		log.Fatalf("Failed to initialize control subscriber: %v", err)
	}
	defer subscriber.Disconnect()

	err = subscriber.Subscribe()
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	go subscriber.Wait()
	go publishWarningMessages(publisher, sensor)
	publishRegularly(publisher, time.Duration(*interval)*time.Second)
}

func publishRegularly(publisher *mqtt_utils.Publisher, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		publishValue(publisher)
	}
}

func publishValue(publisher *mqtt_utils.Publisher) {
	if publisher.GetDevice().GetStatus() == "on" {

		err := publisher.Publish()
		if err != nil {
			log.Printf("Error publishing sensor data: %v", err)
		} else {
			fmt.Println("Sensor data published successfully.")
		}
	}
}

func publishWarningMessages(publisher *mqtt_utils.Publisher, sensor *sensors.Sensor) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		warning, err := sensor.CheckValueInRange()
		if publisher.GetDevice().GetStatus() == "on" {
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
}
