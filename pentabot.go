package pentabot

import (
	"github.com/hybridgroup/gobot"
)

var (
	currentStage string
)

func Init() {
	UpdateStage("powerUp")

	//TODO: Make this better
	InitSphero()
	InitArduino()
	InitJoystick()
	InitLeap()
	InitPebble()
}

// Execute work according to current Stage ... spheroWhatever.Work()
func Work() {}

//TODO: Make this better
func Devices() []gobot.Device {
	return []gobot.Device{spheroDriver, led1, led2, led3, led4, led5, led6, joystickDriver, leapDriver, pebbleDriver}
}

//TODO: Make this better
func Connections() []gobot.Connection {
	return []gobot.Connection{spheroAdaptor, firmataAdaptor, joystickAdaptor, leapAdaptor, pebbleAdaptor}
}

//TODO: Change this to NextStage
func UpdateStage(stage string) string {
	currentStage = stage
	return currentStage
}

func CurrentStage() string {
	return currentStage
}
