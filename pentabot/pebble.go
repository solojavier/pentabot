package pentabot

import (
	"math"
	"strconv"
	"strings"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/pebble"
)

func CreatePebble() *gobot.Robot {
	pebbleAdaptor := pebble.NewPebbleAdaptor("pebble")
	pebbleDriver := pebble.NewPebbleDriver(pebbleAdaptor, "pebble")

	return gobot.NewRobot("pebble",
		[]gobot.Connection{pebbleAdaptor},
		[]gobot.Device{pebbleDriver},
		func() {
			gobot.On(pebbleDriver.Event("accel"), func(data interface{}) {
				if currentStage() == "pebble" {
					values := strings.Split(data.(string), ",")
					x, _ := strconv.ParseFloat(values[0], 64)
					y, _ := strconv.ParseFloat(values[1], 64)

					speed := math.Max(math.Abs(x), math.Abs(y))
					heading := 180.0 - (math.Atan2(x, y) * (180.0 / math.Pi))

					spheroRoll(scalePebble(speed), uint16(heading))
				}
			})
		},
	)
}

func scalePebble(position float64) uint8 {
	return uint8(gobot.ToScale(gobot.FromScale(position, 0, 1000), 0, 255))
}
