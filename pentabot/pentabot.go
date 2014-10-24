package pentabot

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

const maxPower = 5000

var (
	CurrentPower int
	CurrentStage string
	PowerLevel   int
	X            float64
	Y            float64
)

func Init() {
	CurrentStage = "powerUp"
	CurrentPower = 0
	PowerLevel = 0
	X = 0
	Y = 0
}

func PowerUp(sphero sphero.SpheroDriver, collision sphero.Collision) {
	rawPower := collision.XMagnitude + collision.YMagnitude
	powerDelta := int(gobot.ToScale(gobot.FromScale(float64(rawPower), 0, 800), 0, 255))

	CurrentPower += powerDelta

	if CurrentPower > 255 {
		CurrentPower = 0
		PowerLevel += 1
	}

	sphero.SetRGB(uint8(255-CurrentPower), uint8(CurrentPower), 0)
}

func Scale(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 1), 0, 255))
}

func ScaleJoystick(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 32768), 0, 255))
}
