package main

import (
	"anapi/src/configuration"
	"anapi/src/controllers"

	"github.com/deltegui/phoenix"
)

func main() {
	app := phoenix.NewApp()
	app.Configure().
		SetProjectInfo("anapi", "0.1.0").
		EnableLogoFile()
	config := configuration.Load()
	controllers.Register(app)
	app.Run(config.ListenURL)
}
