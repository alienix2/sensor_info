package mocks

type MockFormatter struct {
	FormatErr     error
	ParseErr      error
	FormattedData string
	ParsedValue   float64
}

func (m *MockFormatter) Format(value float64, name, unit, id string) (string, error) {
	return m.FormattedData, m.FormatErr
}

func (m *MockFormatter) Parse(data string) (float64, error) {
	return m.ParsedValue, m.ParseErr
}
