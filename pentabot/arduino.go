package pentabot

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

var (
	leds []*gpio.LedDriver
)

func CreateArduino() *gobot.Robot {
	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem14111")
	led1 := gpio.NewLedDriver(firmataAdaptor, "led", "4")
	led2 := gpio.NewLedDriver(firmataAdaptor, "led", "5")
	led3 := gpio.NewLedDriver(firmataAdaptor, "led", "6")
	led4 := gpio.NewLedDriver(firmataAdaptor, "led", "7")
	led5 := gpio.NewLedDriver(firmataAdaptor, "led", "8")
	led6 := gpio.NewLedDriver(firmataAdaptor, "led", "9")
	leds = []*gpio.LedDriver{led1, led2, led3, led4, led5, led6}

	return gobot.NewRobot("arduino",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2, led3, led4, led5, led6},
		func() {},
	)
}
