package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/solojavier/pentabot"
)

func main() {
	gbot := gobot.NewGobot()
	api := api.NewAPI(gbot)
	api.Port = "8080"
	api.Start()

	pentabot.Init()

	robot := gobot.NewRobot("pentabot",
		pentabot.Connections(),
		pentabot.Devices(),
		pentabot.Work,
	)

	robot.AddCommand("update_stage", func(params map[string]interface{}) interface{} {
		return pentabot.UpdateStage(params["stage"].(string))
	})

	robot.AddCommand("move", func(params map[string]interface{}) interface{} {
		return pentabot.Move(params["direction"].(string))
	})

	gbot.AddRobot(robot)

	gbot.Start()
}
