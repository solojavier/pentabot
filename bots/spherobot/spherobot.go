package spherobot

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"github.com/solojavier/pentabot/bots/arduinobot"
)

const maxPower = 5000

type Bot struct {
	currentPower  int
	powerLevel    int
	spheroDriver  *sphero.SpheroDriver
	spheroAdaptor *sphero.SpheroAdaptor
	arduinoBot    *arduinobot.Bot
}

func New(arduinoBot *arduinobot.Bot) *Bot {
	var b Bot
	b.currentPower = 0
	b.powerLevel = 0
	b.spheroAdaptor = sphero.NewSpheroAdaptor("Sphero", "/dev/tty.Sphero-YBW-RN-SPP")
	b.spheroDriver = sphero.NewSpheroDriver(b.spheroAdaptor, "sphero")
	b.arduinoBot = arduinoBot
	return &b
}

func (b *Bot) Work() {
	gobot.On(b.spheroDriver.Event("collision"), func(data interface{}) {
		b.powerUp(data.(sphero.Collision))

		if b.powerLevel > 0 && b.powerLevel < b.arduinoBot.Leds() {
			b.arduinoBot.LedOn(b.powerLevel - 1)
		} else if b.powerLevel == b.arduinoBot.Leds() {
			b.arduinoBot.LedOn(b.powerLevel - 1)
			b.spheroDriver.SetRGB(0, 0, 255)
		}
	})
}

func (b *Bot) Devices() []gobot.Device {
	return []gobot.Device{b.spheroDriver}
}

func (b *Bot) Connection() gobot.Connection {
	return b.spheroAdaptor
}

func (b *Bot) Roll(speed uint8, heading uint16) {
	b.spheroDriver.Roll(speed, heading)
}

func (b *Bot) Stop() {
	b.spheroDriver.Stop()
}

func (b *Bot) Move(direction string) string {
	message := "Not moving. Expecting direction param with value: up, down, left or right"

	switch direction {
	case "up":
		b.spheroDriver.Roll(100, 0)
		message = "Moving up"
	case "down":
		b.spheroDriver.Roll(100, 90)
		message = "Moving down"
	case "left":
		b.spheroDriver.Roll(100, 180)
		message = "Moving left"
	case "right":
		b.spheroDriver.Roll(100, 270)
		message = "Moving right"
	}

	time.Sleep(3 * time.Second)
	b.spheroDriver.Stop()
	return message
}

func (b *Bot) powerUp(collision sphero.Collision) {
	rawPower := collision.XMagnitude + collision.YMagnitude
	powerDelta := int(gobot.ToScale(gobot.FromScale(float64(rawPower), 0, 800), 0, 255))

	b.currentPower += powerDelta

	if b.currentPower > 255 {
		b.currentPower = 0
		b.powerLevel += 1
	}

	b.spheroDriver.SetRGB(uint8(255-b.currentPower), uint8(b.currentPower), 0)
}
