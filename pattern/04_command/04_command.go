package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Свет выключен")
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
	}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
	}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndoButton() {
	r.command.Undo()
}

func main() {
	light := &Light{}
	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remoteControl := &RemoteControl{}
	remoteControl.SetCommand(lightOnCommand)

	remoteControl.PressButton()

	remoteControl.SetCommand(lightOffCommand)

	remoteControl.PressButton()

	remoteControl.PressUndoButton()
}
