package pebblebot

import (
	"math"
	"strconv"
	"strings"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/pebble"
	"github.com/solojavier/pentabot/bots/spherobot"
)

var (
	pebbleAdaptor *pebble.PebbleAdaptor
	pebbleDriver  *pebble.PebbleDriver
)

func Init() {
	pebbleAdaptor = pebble.NewPebbleAdaptor("pebble")
	pebbleDriver = pebble.NewPebbleDriver(pebbleAdaptor, "pebble")
}

func Work() {
	gobot.On(pebbleDriver.Event("accel"), func(data interface{}) {
		values := strings.Split(data.(string), ",")
		x, _ := strconv.ParseFloat(values[0], 64)
		y, _ := strconv.ParseFloat(values[1], 64)

		speed := math.Max(math.Abs(x), math.Abs(y))
		heading := 180.0 - (math.Atan2(y, x) * (180.0 / math.Pi))

		spherobot.Roll(scalePebble(speed), uint16(heading))
	})
}

func Devices() []gobot.Device {
	return []gobot.Device{pebbleDriver}
}

func Connection() gobot.Connection {
	return pebbleAdaptor
}

func scalePebble(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 1000), 0, 255))
}
