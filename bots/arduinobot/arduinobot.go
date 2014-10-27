package arduinobot

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

type Bot struct {
	leds           []*gpio.LedDriver
	firmataAdaptor *firmata.FirmataAdaptor
	led1           *gpio.LedDriver
	led2           *gpio.LedDriver
	led3           *gpio.LedDriver
	led4           *gpio.LedDriver
	led5           *gpio.LedDriver
	led6           *gpio.LedDriver
}

func New() *Bot {
	var b Bot
	b.firmataAdaptor = firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem14141")
	b.led1 = gpio.NewLedDriver(b.firmataAdaptor, "led", "4")
	b.led2 = gpio.NewLedDriver(b.firmataAdaptor, "led", "5")
	b.led3 = gpio.NewLedDriver(b.firmataAdaptor, "led", "6")
	b.led4 = gpio.NewLedDriver(b.firmataAdaptor, "led", "7")
	b.led5 = gpio.NewLedDriver(b.firmataAdaptor, "led", "8")
	b.led6 = gpio.NewLedDriver(b.firmataAdaptor, "led", "9")
	b.leds = []*gpio.LedDriver{b.led1, b.led2, b.led3, b.led4, b.led5, b.led6}
	return &b
}

func (b *Bot) Work() {}

func (b *Bot) Devices() []gobot.Device {
	return []gobot.Device{b.led1, b.led2, b.led3, b.led4, b.led5, b.led6}
}

func (b *Bot) Connection() gobot.Connection {
	return b.firmataAdaptor
}

func (b *Bot) Leds() int {
	return len(b.leds)
}

func (b *Bot) LedOn(index int) {
	b.leds[index].On()
}
