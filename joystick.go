package pentabot

import (
	"math"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/joystick"
)

var (
	x               float64
	y               float64
	joystickAdaptor *joystick.JoystickAdaptor
	joystickDriver  *joystick.JoystickDriver
)

func InitJoystick() {
	x = 0
	y = 0

	joystickAdaptor = joystick.NewJoystickAdaptor("ps3")
	joystickDriver = joystick.NewJoystickDriver(joystickAdaptor, "ps3", "./config/dualshock3.json")

	gobot.Every(100*time.Millisecond, func() {
		if CurrentStage() == "joystick" {
			speed := math.Max(math.Abs(x), math.Abs(y))
			heading := 180.0 - (math.Atan2(y, x) * (180.0 / math.Pi))

			spheroDriver.Roll(scaleJoystick(speed), uint16(heading))
		}
	})

	gobot.On(joystickDriver.Event("right_x"), func(data interface{}) {
		if CurrentStage() == "joystick" {
			x = float64(data.(int16)) - 128
		}
	})

	gobot.On(joystickDriver.Event("right_y"), func(data interface{}) {
		if CurrentStage() == "joystick" {
			y = float64(data.(int16)) - 128
		}
	})
}

func scaleJoystick(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 32768), 0, 255))
}
