package pentabot

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/joystick"
	"math"
	"time"
)

func CreateJoystick() *gobot.Robot {
	x := 0.0
	y := 0.0
	joystickAdaptor := joystick.NewJoystickAdaptor("ps3")
	joystickDriver := joystick.NewJoystickDriver(joystickAdaptor, "ps3", "./config/dualshock3.json")

	return gobot.NewRobot("joystick",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{joystickDriver},
		func() {
			gobot.Every(100*time.Millisecond, func() {
				if currentStage() == "joystick" {
					speed := math.Max(math.Abs(x), math.Abs(y))
					heading := 180.0 - (math.Atan2(y, x) * (180.0 / math.Pi))

					fmt.Println(speed)
					spheroRoll(scaleJoystick(speed), uint16(heading))
				}
			})

			gobot.On(joystickDriver.Event("right_x"), func(data interface{}) {
				if currentStage() == "joystick" {
					x = float64(data.(int16)) - 128
				}
			})

			gobot.On(joystickDriver.Event("right_y"), func(data interface{}) {
				if currentStage() == "joystick" {
					y = float64(data.(int16)) - 128
				}
			})
		},
	)
}

func scaleJoystick(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 32768), 0, 255))
}
