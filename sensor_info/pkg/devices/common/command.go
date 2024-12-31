package devices

import "log"

type Command interface {
	Execute() error
}

type TurnOnCommand struct {
	Device Device
}

func (c *TurnOnCommand) Execute() error {
	log.Println("Stub for turnOnCommand executed, Implement actual low level interaction!")
	c.Device.SetStatus("on")
	return nil
}

type TurnOffCommand struct {
	Device Device
}

func (c *TurnOffCommand) Execute() error {
	log.Println("Stub for turnOffCommand executed, Implement actual low level interaction!")
	c.Device.SetStatus("off")
	return nil
}
