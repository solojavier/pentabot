package pentabot

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"time"
)

var (
	spheroDriver *sphero.SpheroDriver
)

func CreateSphero() *gobot.Robot {
	spheroAdaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-YBW-RN-SPP")
	spheroDriver = sphero.NewSpheroDriver(spheroAdaptor, "sphero")

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{spheroAdaptor},
		[]gobot.Device{spheroDriver},
		func() {
			gobot.On(spheroDriver.Event("collision"), func(data interface{}) {
				if currentStage() == "powerUp" {
					powerUp(data.(sphero.Collision))
					spheroDriver.SetRGB(uint8(255-currentPower), uint8(currentPower), 0)

					if powerLevel > 0 && powerLevel < len(leds) {
						ledOn()
					}

					if powerLevel == len(leds) {
						ledOn()
						spheroDriver.SetRGB(0, 50, 150)
						NextStage()
					}
				}
			})
		},
	)

	robot.AddCommand("move", func(params map[string]interface{}) interface{} {
		if currentStage() == "commander" {

			direction := params["direction"].(string)
			message := "Not moving. Expecting direction param with value: up, down, left or right"

			switch direction {
			case "up":
				spheroDriver.Roll(100, 0)
				message = "Moving up"
			case "down":
				spheroDriver.Roll(100, 180)
				message = "Moving down"
			case "left":
				spheroDriver.Roll(100, 270)
				message = "Moving left"
			case "right":
				spheroDriver.Roll(100, 90)
				message = "Moving right"
			}

			time.Sleep(2 * time.Second)
			spheroDriver.Stop()
			return message

		} else {
			return "Not moving. Current stage is not commander"
		}
	})

	return robot
}

func powerUp(collision sphero.Collision) {
	rawPower := collision.XMagnitude + collision.YMagnitude
	powerDelta := int(gobot.ToScale(gobot.FromScale(float64(rawPower), 0, 800), 0, 255))

	currentPower += powerDelta

	if currentPower > 255 {
		currentPower = 0
		powerLevel += 1
	}
}
