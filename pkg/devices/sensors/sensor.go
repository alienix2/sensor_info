package devices

import (
	"log"

	common "github.com/alienix2/sensor_info/pkg/devices/common"
	"github.com/google/uuid"
)

type Sensor struct {
	reader    ReaderStrategy
	formatter common.DeviceFormatterStrategy
	name      string
	id        string
	unit      string
	status    string
	minValue  float64
	maxValue  float64
}

func (s *Sensor) FormatData() (string, error) {
	data, err := s.reader.Read(s)
	if err != nil {
		return "", err
	}
	return s.formatter.Format(data, s.name, s.unit, s.id)
}

func (s *Sensor) GetRange() (float64, float64) {
	return s.minValue, s.maxValue
}

func (s *Sensor) GetID() string {
	return s.id
}

func (s *Sensor) GetName() string {
	return s.name
}

func (s *Sensor) GetStatus() string {
	return s.status
}

func (s *Sensor) ParseDeviceValue(data string) (float64, error) {
	return s.formatter.Parse(data)
}

func (s *Sensor) SetStatus(status string) {
	s.status = status
}

type Option func(*Sensor)

func WithSensorName(name string) Option {
	return func(s *Sensor) {
		s.name = name
	}
}

func WithSensorUnit(unit string) Option {
	return func(s *Sensor) {
		s.unit = unit
	}
}

func WithSensorRange(min, max float64) Option {
	return func(s *Sensor) {
		s.minValue = min
		s.maxValue = max
	}
}

func WithSensorFormatterStrategy(formatter common.DeviceFormatterStrategy) Option {
	return func(s *Sensor) {
		s.formatter = formatter
	}
}

func WithReaderStrategy(reader ReaderStrategy) Option {
	return func(s *Sensor) {
		s.reader = reader
	}
}

func WithSensorID(id string) Option {
	return func(s *Sensor) {
		s.id = id
	}
}

func NewSensor(opts ...Option) *Sensor {
	sensor := &Sensor{
		name:      "DefaultSensor",
		id:        "sensor_autogenerated-" + uuid.New().String(),
		unit:      "unit",
		minValue:  0,
		maxValue:  100,
		formatter: &common.RawFormatterStrategy{},
		reader:    &DefaultReader{},
		status:    "on",
	}

	for _, opt := range opts {
		opt(sensor)
	}

	log.Printf("Sensor initialized: %+v\n", sensor)

	return sensor
}
