package devices

import (
	"fmt"
)

type Command interface {
	Execute() error
}

type TurnOnCommand struct {
	Device Device
}

func (c *TurnOnCommand) Execute() error {
	fmt.Println("Stub for turnOnCommand executed, Implement actual low level interaction!")
	c.Device.SetStatus("on")
	return nil
}

type TurnOffCommand struct {
	Device Device
}

func (c *TurnOffCommand) Execute() error {
	fmt.Println("Stub for turnOffCommand executed, Implement actual low level interaction!")
	c.Device.SetStatus("off")
	return nil
}
