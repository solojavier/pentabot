package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/joystick"
	"github.com/hybridgroup/gobot/platforms/leap"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"github.com/solojavier/pentabot-game/pentabot"
)

func main() {
	gbot := gobot.NewGobot()
	api.NewAPI(gbot).Start()

	leapAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	spheroAdaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-YBW-RN-SPP")
	joystickAdaptor := joystick.NewJoystickAdaptor("ps3")
	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem14121")

	leapDriver := leap.NewLeapMotionDriver(leapAdaptor, "leap")
	spheroDriver := sphero.NewSpheroDriver(spheroAdaptor, "sphero")
	joystickDriver := joystick.NewJoystickDriver(joystickAdaptor, "ps3", "./dualshock3.json")
	led1 := gpio.NewLedDriver(firmataAdaptor, "led", "4")
	led2 := gpio.NewLedDriver(firmataAdaptor, "led", "5")
	led3 := gpio.NewLedDriver(firmataAdaptor, "led", "6")
	led4 := gpio.NewLedDriver(firmataAdaptor, "led", "7")
	led5 := gpio.NewLedDriver(firmataAdaptor, "led", "8")
	led6 := gpio.NewLedDriver(firmataAdaptor, "led", "9")
	leds := []*gpio.LedDriver{led1, led2} //, led3, led4, led5, led6}

	pentabot.Init()

	work := func() {

		gobot.On(spheroDriver.Event("collision"), func(data interface{}) {
			if pentabot.CurrentStage == "powerUp" {
				pentabot.PowerUp(*spheroDriver, data.(sphero.Collision))

				if pentabot.PowerLevel > 0 && pentabot.PowerLevel < len(leds) {
					leds[pentabot.PowerLevel-1].On()
				} else if pentabot.PowerLevel == len(leds) {
					leds[pentabot.PowerLevel-1].On()
					pentabot.CurrentStage = "joystick"
					spheroDriver.SetRGB(0, 255, 0)
					fmt.Println("joystick")
				}
			}
		})

		gobot.On(leapDriver.Event("message"), func(data interface{}) {
			if pentabot.CurrentStage == "leap" {
				hands := data.(leap.Frame).Hands

				if len(hands) > 0 {
					x := math.Abs(hands[0].Direction[0])
					y := math.Abs(hands[0].Direction[1])
					z := math.Abs(hands[0].Direction[2])
					spheroDriver.SetRGB(pentabot.Scale(x), pentabot.Scale(y), pentabot.Scale(z))
				}
			}
		})

		gobot.On(joystickDriver.Event("right_x"), func(data interface{}) {
			if pentabot.CurrentStage == "joystick" {
				pentabot.X = float64(data.(int16)) - 128
			}
		})
		gobot.On(joystickDriver.Event("right_y"), func(data interface{}) {
			if pentabot.CurrentStage == "joystick" {
				pentabot.Y = float64(data.(int16)) - 128
			}
		})

		gobot.Every(100*time.Millisecond, func() {
			if pentabot.CurrentStage != "powerUp" {
				speed := math.Max(math.Abs(pentabot.X), math.Abs(pentabot.Y))
				heading := 180.0 - (math.Atan2(pentabot.Y, pentabot.X) * (180.0 / math.Pi))

				spheroDriver.Roll(pentabot.ScaleJoystick(speed), uint16(heading))
			}
		})
	}

	robot := gobot.NewRobot("pentabot",
		[]gobot.Connection{firmataAdaptor, leapAdaptor, spheroAdaptor, joystickAdaptor},
		[]gobot.Device{leapDriver, led1, led2, led3, led4, led5, led6, spheroDriver, joystickDriver},
		work,
	)

	robot.AddCommand("updateStage", func(params map[string]interface{}) interface{} {
		pentabot.CurrentStage = params["stage"].(string)
		return pentabot.CurrentStage
	})

	robot.AddCommand("move", func(params map[string]interface{}) interface{} {
		if pentabot.CurrentStage == "commander" {
			message := "Not moving. Expecting direction param with value: up, down, left or right"

			switch params["direction"] {
			case "up":
				spheroDriver.Roll(100, 0)
				message = "Moving up"
			case "down":
				spheroDriver.Roll(100, 90)
				message = "Moving down"
			case "left":
				spheroDriver.Roll(100, 180)
				message = "Moving left"
			case "right":
				spheroDriver.Roll(100, 270)
				message = "Moving right"
			}

			time.Sleep(3 * time.Second)
			spheroDriver.Stop()
			return message
		} else {
			return "Not moving. Commander stage is not active"
		}
	})

	gbot.AddRobot(robot)

	gbot.Start()
}
