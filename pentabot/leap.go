package pentabot

import (
	"math"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
)

func CreateLeap() *gobot.Robot {
	leapAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	leapDriver := leap.NewLeapMotionDriver(leapAdaptor, "leap")

	return gobot.NewRobot("leap",
		[]gobot.Connection{leapAdaptor},
		[]gobot.Device{leapDriver},
		func() {
			gobot.On(leapDriver.Event("message"), func(data interface{}) {
				if currentStage() == "leap" {
					hands := data.(leap.Frame).Hands

					if len(hands) > 0 {
						x := hands[0].X()
						z := hands[0].Z()
						speed := math.Max(math.Abs(x), math.Abs(z))
						heading := 180.0 - (math.Atan2(z, x) * (180.0 / math.Pi))

						spheroRoll(scaleLeap(speed), uint16(heading))
					} else {
						spheroStop()
					}
				}
			})
		},
	)
}

func scaleLeap(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 200), 0, 255))
}
