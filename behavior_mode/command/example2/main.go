package main

import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}

// 关闭命令 - start
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// 关闭命令 - end

// 开启命令 - start
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// 关闭命令 - end

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}
	onButton := &button{
		command: &onCommand{
			device: tv,
		},
	}
	onButton.press()

	offButton := &button{
		command: &offCommand{
			device: tv,
		},
	}
	offButton.press()
}
