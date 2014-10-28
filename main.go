package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/solojavier/pentabot/pentabot"
)

func main() {
	gbot := gobot.NewGobot()
	api := api.NewAPI(gbot)
	api.Port = "8080"
	api.Start()

	gbot.AddCommand("next_stage", func(params map[string]interface{}) interface{} {
		return pentabot.NextStage()
	})

	gbot.AddRobot(pentabot.CreateArduino())
	gbot.AddRobot(pentabot.CreateSphero())
	gbot.AddRobot(pentabot.CreateJoystick())
	gbot.AddRobot(pentabot.CreatePebble())
	gbot.AddRobot(pentabot.CreateLeap())

	gbot.Start()
}
