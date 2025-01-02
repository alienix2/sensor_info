# SensorInfo

Sensor info is a library that wraps MQTT and allows to manage simple sensors in a cohesive way.

## Pre-implemented data management

The data management can be handled with custom strategies that can be easily integrated inside the APIs. The pre-implemented data management functions rely on the use of JSON for sending data, SQLite for the local storage (if needed) and MySQL for the central storage.

There are some functions which are generic, in particular the ones that save the data in SQLite format. Other functions instead are specific to data structures that are defined in the library itself, like the ones regarding the storage of the data in the central database

This has been done to show a way to use the library in a more specific way, allowing to build on top of the generic functions. If the structure provided doesn't fit a specific use case another strategy using MySQL could be easily implemented and used instead.

## Examples

The examples folder contains some examples on how to use the library. Each one can take many parameters as input (*to check them run the go file with the -h parameter).

In particular there are:

- One example of a publisher, which represents a sensor that sends data to the broker using JSON at specific time intervals and if the value that it reads is not within a specific range. That publisher also automatically subscribes to a topic used for internal communications (turning on/off a sensor etc.)
- One example of a subscriber, which represents a sensor that receives data in JSON format from the broker and stores it in a local SQLite database.
