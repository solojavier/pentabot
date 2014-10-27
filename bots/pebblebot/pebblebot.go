package pebblebot

import (
	"math"
	"strconv"
	"strings"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/pebble"
	"github.com/solojavier/pentabot/bots/spherobot"
)

type Bot struct {
	pebbleAdaptor *pebble.PebbleAdaptor
	pebbleDriver  *pebble.PebbleDriver
	spheroBot     *spherobot.Bot
}

func New(spheroBot *spherobot.Bot) *Bot {
	var b Bot
	b.pebbleAdaptor = pebble.NewPebbleAdaptor("pebble")
	b.pebbleDriver = pebble.NewPebbleDriver(b.pebbleAdaptor, "pebble")
	b.spheroBot = spheroBot
	return &b
}

func (b *Bot) Work() {
	gobot.On(b.pebbleDriver.Event("accel"), func(data interface{}) {
		values := strings.Split(data.(string), ",")
		x, _ := strconv.ParseFloat(values[0], 64)
		y, _ := strconv.ParseFloat(values[1], 64)

		speed := math.Max(math.Abs(x), math.Abs(y))
		heading := 180.0 - (math.Atan2(y, x) * (180.0 / math.Pi))

		b.spheroBot.Roll(scalePebble(speed), uint16(heading))
	})
}

func (b *Bot) Devices() []gobot.Device {
	return []gobot.Device{b.pebbleDriver}
}

func (b *Bot) Connection() gobot.Connection {
	return b.pebbleAdaptor
}

func scalePebble(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 1000), 0, 255))
}
