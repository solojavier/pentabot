package pentabot

/*
import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
)

leapAdaptor := leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
leapDriver := leap.NewLeapMotionDriver(leapAdaptor, "leap")

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

func Scale(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 1), 0, 255))
}
*/
