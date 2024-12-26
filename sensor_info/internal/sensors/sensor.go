package sensors

import "github.com/google/uuid"

type Sensor struct {
	reader    ReaderStrategy
	formatter FormatterStrategy
	name      string
	id        string
	unit      string
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

func (s *Sensor) ParseSensorValue(data string) (float64, error) {
	return s.formatter.Parse(data)
}

type Option func(*Sensor)

func WithName(name string) Option {
	return func(s *Sensor) {
		s.name = name
	}
}

func WithUnit(unit string) Option {
	return func(s *Sensor) {
		s.unit = unit
	}
}

func WithRange(min, max float64) Option {
	return func(s *Sensor) {
		s.minValue = min
		s.maxValue = max
	}
}

func WithFormatterStrategy(formatter FormatterStrategy) Option {
	return func(s *Sensor) {
		s.formatter = formatter
	}
}

func WithReaderStrategy(reader ReaderStrategy) Option {
	return func(s *Sensor) {
		s.reader = reader
	}
}

func NewSensor(opts ...Option) *Sensor {
	sensor := &Sensor{
		name:      "DefaultSensor",
		id:        uuid.New().String(),
		unit:      "unit",
		minValue:  0,
		maxValue:  100,
		formatter: &RawFormatterStrategy{},
		reader:    &DefaultReader{},
	}

	for _, opt := range opts {
		opt(sensor)
	}

	return sensor
}
