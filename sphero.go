package pentabot

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

const maxPower = 5000

var (
	currentPower  int
	powerLevel    int
	spheroDriver  *sphero.SpheroDriver
	spheroAdaptor *sphero.SpheroAdaptor
)

func InitSphero() {
	currentPower = 0
	powerLevel = 0
	spheroAdaptor = sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-YBW-RN-SPP")
	spheroDriver = sphero.NewSpheroDriver(spheroAdaptor, "sphero")

	handleSpheroCollision()
}

func Move(direction string) string {
	if CurrentStage() == "commander" {
		message := "Not moving. Expecting direction param with value: up, down, left or right"

		switch direction {
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
}

func powerUp(collision sphero.Collision) {
	rawPower := collision.XMagnitude + collision.YMagnitude
	powerDelta := int(gobot.ToScale(gobot.FromScale(float64(rawPower), 0, 800), 0, 255))

	currentPower += powerDelta

	if currentPower > 255 {
		currentPower = 0
		powerLevel += 1
	}

	spheroDriver.SetRGB(uint8(255-currentPower), uint8(currentPower), 0)
}

func handleSpheroCollision() {
	gobot.On(spheroDriver.Event("collision"), func(data interface{}) {
		if CurrentStage() == "powerUp" {
			powerUp(data.(sphero.Collision))

			if powerLevel > 0 && powerLevel < Leds() {
				LedOn(powerLevel - 1)
			} else if powerLevel == Leds() {
				LedOn(powerLevel - 1)
				UpdateStage("joystick")
				spheroDriver.SetRGB(0, 0, 255)
			}
		}
	})
}
