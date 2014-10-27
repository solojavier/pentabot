package leapbot

import (
	"math"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
	"github.com/solojavier/pentabot/bots/spherobot"
)

var (
	leapAdaptor *leap.LeapMotionAdaptor
	leapDriver  *leap.LeapMotionDriver
)

func Init() {
	leapAdaptor = leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	leapDriver = leap.NewLeapMotionDriver(leapAdaptor, "leap")
}

func Work() {
	gobot.On(leapDriver.Event("message"), func(data interface{}) {
		hands := data.(leap.Frame).Hands

		if len(hands) > 0 {
			x := hands[0].X()
			z := hands[0].Z()
			speed := math.Max(math.Abs(x), math.Abs(z))
			heading := 180.0 - (math.Atan2(z, x) * (180.0 / math.Pi))

			spherobot.Roll(scaleLeap(speed), uint16(heading))
		} else {
			spherobot.Stop()
		}
	})
}

func Devices() []gobot.Device {
	return []gobot.Device{leapDriver}
}

func Connection() gobot.Connection {
	return leapAdaptor
}

func scaleLeap(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 200), 0, 255))
}
