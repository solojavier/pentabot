package pentabot

import (
	"github.com/hybridgroup/gobot"
	"github.com/solojavier/pentabot/bots/arduinobot"
	"github.com/solojavier/pentabot/bots/joystickbot"
	"github.com/solojavier/pentabot/bots/leapbot"
	"github.com/solojavier/pentabot/bots/pebblebot"
	"github.com/solojavier/pentabot/bots/spherobot"
)

type Bot interface {
	Init()
	Work()
	Devices()
	Connection()
}

var (
	currentStage string
	devices      []gobot.Device
	connections  []gobot.Connection
)

func Init() {
	currentStage = "powerUp"

	//TODO: Remove duplication
	spherobot.Init()
	arduinobot.Init()
	joystickbot.Init()
	leapbot.Init()
	pebblebot.Init()

	devices = append(devices, spherobot.Devices()...)
	devices = append(devices, arduinobot.Devices()...)
	devices = append(devices, joystickbot.Devices()...)
	devices = append(devices, leapbot.Devices()...)
	devices = append(devices, pebblebot.Devices()...)

	connections = append(connections, spherobot.Connection())
	connections = append(connections, arduinobot.Connection())
	connections = append(connections, joystickbot.Connection())
	connections = append(connections, leapbot.Connection())
	connections = append(connections, pebblebot.Connection())
}

func Work() {
	switch currentStage {
	case "powerUp":
		spherobot.Work()
	case "pebble":
		pebblebot.Work()
	case "leap":
		leapbot.Work()
	case "joystick":
		joystickbot.Work()
	}
}

func Devices() []gobot.Device {
	return devices
}

func Connections() []gobot.Connection {
	return connections
}

func UpdateStage(stage string) string {
	currentStage = stage
	Work()
	return currentStage
}

func Move(direction string) string {
	if currentStage == "commander" {
		return spherobot.Move(direction)
	} else {
		return "Not moving. Commander stage not active."
	}
}
