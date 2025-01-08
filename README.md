# SensorInfo

Sensor info is a library that wraps MQTT and allows to manage simple sensors in a cohesive way.

## Pre-implemented data management

The data management can be handled with custom strategies that can be easily integrated inside the APIs. The pre-implemented data management functions rely on the use of JSON for sending data, SQLite for the local storage (if needed) and MySQL for the central storage.

There are some functions which are generic, in particular the ones that save the data in SQLite format. Other functions instead are specific to data structures that are defined in the library itself, like the ones regarding the storage of the data in the central database

This has been done to show a way to use the library in a more specific way, allowing to build on top of the generic functions. If the structure provided doesn't fit a specific use case another strategy using MySQL could be easily implemented and used instead.

An app to visualize the data stored in the MySQL database is provided at the following repository: <https://github.com/alienix2/sensor_visualization>

## Broker

The examples are setup to dialog easily with the Mosquitto broker, in particular there's the handling of the tls communication. Anyway the library is broker agnostic and is tested using another broker that can be launched directly from inside the test files.

## Testing

The library has many tests cases covered and moreover also offers a few utility files that give easy access to mocks for the MySQL database using Testcontainers (<https://golang.testcontainers.org/modules/mariadb/>) and for the MQTT broker using mochi-mqtt (<https://github.com/mochi-mqtt/server>)

## Examples

The examples folder contains some examples on how to use the library. Each one can take many parameters as input (to check them run the go file with the -h parameter).

In particular there are:

- One example of a publisher, which represents a sensor that sends data to the broker using JSON at specific time intervals and if the value that it reads is not within a specific range. That publisher also automatically subscribes to a topic used for internal communications (turning on/off a sensor etc.)
- One example of a subscriber, which represents a sensor that receives data in JSON format from the broker and stores it in a local SQLite database. The subscriber also has a simple implementation of the handler that verifies the data received
- One omnisubscriber which subscribes to all the topics and stores the message data in a MySQL database. This subscriber is given as it's fundamental to handle the visualization of all the data sent to a broker as today there is no easy way to handle the easy storing of such data directly from the Mosquitto broker.

### Usage of the examples

The examples can be run using the `go run examples/<sepcific_example_folder>` command. By default they look for certifications inside the *certifications* folder **which is empty by default** and therefore the publisher and subscribers will run without any certificate.

The default broker address is the one usually set up for the TLS, if you want to use the default one for the TCP one, pass it as a parameter.

If you need them to be used, just put the inside that folder.
