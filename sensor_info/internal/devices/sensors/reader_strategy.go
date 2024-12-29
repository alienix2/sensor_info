package devices

import (
	"fmt"
	"math/rand"
)

type ReaderStrategy interface {
	Read(s *Sensor) (float64, error)
}

type DefaultReader struct{}

func (d *DefaultReader) Read(s *Sensor) (float64, error) {
	randomValue := rand.Float64()
	scaledValue := s.minValue + (randomValue * (s.maxValue - s.minValue))

	fmt.Printf("Stub implementation of data reading for sensor [%s]: random value in range %.2f-%.2f %.2f\n", s.name, s.minValue, s.maxValue, scaledValue)
	return scaledValue, nil
}
