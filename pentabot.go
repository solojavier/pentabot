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
	Work()
	Devices() []gobot.Device
	Connection() gobot.Connection
}

var (
	currentStage string
	devices      []gobot.Device
	connections  []gobot.Connection
	bots         map[string]Bot
)

func Init() {
	currentStage = "sphero"
	bots = make(map[string]Bot)

	bots["arduino"] = arduinobot.New()
	bots["sphero"] = spherobot.New(bots["arduino"].(*arduinobot.Bot))
	bots["pebble"] = pebblebot.New(bots["sphero"].(*spherobot.Bot))
	bots["leap"] = leapbot.New(bots["sphero"].(*spherobot.Bot))
	bots["joystick"] = joystickbot.New(bots["sphero"].(*spherobot.Bot))

	for _, bot := range bots {
		devices = append(devices, bot.Devices()...)
		connections = append(connections, bot.Connection())
	}
}

func Work() {
	bots[currentStage].Work()
}

func Devices() []gobot.Device {
	return devices
}

func Connections() []gobot.Connection {
	return connections
}

func UpdateStage(stage string) string {
	currentStage = stage
	if currentStage != "commander" {
		Work()
	}
	return currentStage
}

func Move(direction string) string {
	if currentStage == "commander" {
		return bots["sphero"].(*spherobot.Bot).Move(direction)
	} else {
		return "Not moving. Commander stage not active."
	}
}
