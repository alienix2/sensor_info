package mocks

import "errors"

type ManualMockCommand struct {
	ShouldFail bool
}

func (m *ManualMockCommand) Execute() error {
	if m.ShouldFail {
		return errors.New("command execution failed")
	}
	return nil
}
