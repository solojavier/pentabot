package pentabot

import "fmt"

const (
	maxPower = 5000
)

var (
	currentPower = 0
	powerLevel   = 0
	stageIndex   = 0
	stages       = []string{"powerUp", "commander", "joystick", "pebble", "leap"}
)

func NextStage() string {
	stageIndex += 1
	stage := currentStage()
	fmt.Println(stage)
	return stage
}

func ledOn() {
	index := powerLevel - 1
	leds[index].On()
}

func spheroRoll(speed uint8, heading uint16) {
	spheroDriver.Roll(speed, heading)
}

func spheroStop() {
	spheroDriver.Stop()
}

func currentStage() string {
	return stages[stageIndex]
}
