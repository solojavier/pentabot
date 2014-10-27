package leapbot

import (
	"math"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/leap"
	"github.com/solojavier/pentabot/bots/spherobot"
)

type Bot struct {
	leapAdaptor *leap.LeapMotionAdaptor
	leapDriver  *leap.LeapMotionDriver
	spheroBot   *spherobot.Bot
}

func New(spheroBot *spherobot.Bot) *Bot {
	var b Bot
	b.leapAdaptor = leap.NewLeapMotionAdaptor("leap", "127.0.0.1:6437")
	b.leapDriver = leap.NewLeapMotionDriver(b.leapAdaptor, "leap")
	b.spheroBot = spheroBot
	return &b
}

func (b *Bot) Work() {
	gobot.On(b.leapDriver.Event("message"), func(data interface{}) {
		hands := data.(leap.Frame).Hands

		if len(hands) > 0 {
			x := hands[0].X()
			z := hands[0].Z()
			speed := math.Max(math.Abs(x), math.Abs(z))
			heading := 180.0 - (math.Atan2(z, x) * (180.0 / math.Pi))

			b.spheroBot.Roll(scaleLeap(speed), uint16(heading))
		} else {
			b.spheroBot.Stop()
		}
	})
}

func (b *Bot) Devices() []gobot.Device {
	return []gobot.Device{b.leapDriver}
}

func (b *Bot) Connection() gobot.Connection {
	return b.leapAdaptor
}

func scaleLeap(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 200), 0, 255))
}
