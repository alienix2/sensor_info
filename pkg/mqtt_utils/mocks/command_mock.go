package mocks

import (
	"sync"
)

type MockCommand struct {
	Name     string
	Executed bool
	Mu       sync.Mutex
}

func (m *MockCommand) Execute() error {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.Executed = true
	return nil
}
