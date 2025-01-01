package mocks

type MockCommand struct {
	Name string
}

func (m *MockCommand) Execute() error {
	return nil
}
