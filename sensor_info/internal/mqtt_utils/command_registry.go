package mqtt_utils

import devices "mattemoni.sensor_info/internal/devices/common"

type CommandRegistry struct {
	commands map[string]devices.Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]devices.Command),
	}
}

func (r *CommandRegistry) RegisterCommand(name string, command devices.Command) {
	r.commands[name] = command
}

func (r *CommandRegistry) UnregisterCommand(name string) {
	delete(r.commands, name)
}

func (r *CommandRegistry) getCommand(name string) (devices.Command, bool) {
	command, exists := r.commands[name]
	return command, exists
}

func (r *CommandRegistry) GetCommands() map[string]devices.Command {
	return r.commands
}
