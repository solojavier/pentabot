package joystickbot

import (
	"math"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/joystick"
	"github.com/solojavier/pentabot/bots/spherobot"
)

type Bot struct {
	x               float64
	y               float64
	joystickAdaptor *joystick.JoystickAdaptor
	joystickDriver  *joystick.JoystickDriver
	spheroBot       *spherobot.Bot
}

func New(spheroBot *spherobot.Bot) *Bot {
	var b Bot
	b.x = 0
	b.y = 0

	b.joystickAdaptor = joystick.NewJoystickAdaptor("ps3")
	b.joystickDriver = joystick.NewJoystickDriver(b.joystickAdaptor, "ps3", "./config/dualshock3.json")
	b.spheroBot = spheroBot
	return &b
}

func (b *Bot) Work() {
	gobot.Every(100*time.Millisecond, func() {
		speed := math.Max(math.Abs(b.x), math.Abs(b.y))
		heading := 180.0 - (math.Atan2(b.y, b.x) * (180.0 / math.Pi))

		b.spheroBot.Roll(scaleJoystick(speed), uint16(heading))
	})

	gobot.On(b.joystickDriver.Event("right_x"), func(data interface{}) {
		b.x = float64(data.(int16)) - 128
	})

	gobot.On(b.joystickDriver.Event("right_y"), func(data interface{}) {
		b.y = float64(data.(int16)) - 128
	})
}

func (b *Bot) Devices() []gobot.Device {
	return []gobot.Device{b.joystickDriver}
}

func (b *Bot) Connection() gobot.Connection {
	return b.joystickAdaptor
}

func scaleJoystick(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 32768), 0, 255))
}
