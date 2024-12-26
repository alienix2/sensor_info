package sensors

import (
	"fmt"
	"strconv"
)

type FormatterStrategy interface {
	Format(data float64, name, unit, uid string) (string, error)
	Parse(data string) (float64, error)
}

type RawFormatterStrategy struct{}

func (r *RawFormatterStrategy) Format(data float64, name, unit, uid string) (string, error) {
	formattedData := fmt.Sprintf("%.2f, %s, %s, %s", data, name, unit, uid)
	return formattedData, nil
}

func (r *RawFormatterStrategy) Parse(data string) (float64, error) {
	return strconv.ParseFloat(data, 64)
}
