package devices

type ControlExecution struct {
	command Command
}

func (s *ControlExecution) SetCommand(command Command) {
	s.command = command
}

func (s *ControlExecution) Execute() error {
	return s.command.Execute()
}
