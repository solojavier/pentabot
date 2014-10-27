package joystickbot

import (
	"math"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/joystick"
	"github.com/solojavier/pentabot/bots/spherobot"
)

var (
	x               float64
	y               float64
	joystickAdaptor *joystick.JoystickAdaptor
	joystickDriver  *joystick.JoystickDriver
)

func Init() {
	x = 0
	y = 0

	joystickAdaptor = joystick.NewJoystickAdaptor("ps3")
	joystickDriver = joystick.NewJoystickDriver(joystickAdaptor, "ps3", "./config/dualshock3.json")

}

func Work() {
	gobot.Every(100*time.Millisecond, func() {
		speed := math.Max(math.Abs(x), math.Abs(y))
		heading := 180.0 - (math.Atan2(y, x) * (180.0 / math.Pi))

		spherobot.Roll(scaleJoystick(speed), uint16(heading))
	})

	gobot.On(joystickDriver.Event("right_x"), func(data interface{}) {
		x = float64(data.(int16)) - 128
	})

	gobot.On(joystickDriver.Event("right_y"), func(data interface{}) {
		y = float64(data.(int16)) - 128
	})
}

func Devices() []gobot.Device {
	return []gobot.Device{joystickDriver}
}

func Connection() gobot.Connection {
	return joystickAdaptor
}

func scaleJoystick(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 32768), 0, 255))
}
