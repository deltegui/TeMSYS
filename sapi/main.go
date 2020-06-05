package main

import (
	"reflect"
	"sensorapi/src/configuration"
	"sensorapi/src/controllers"
	"sensorapi/src/domain"

	"github.com/deltegui/phoenix"
)

func main() {
	app := phoenix.NewApp()
	app.Configure().
		SetProjectInfo("sensorapi", "0.1.0").
		EnableLogoFile()
	config := configuration.Load()
	controllers.Register(app, config)
	app.Injector.ShowAvailableBuilders()
	reporterType := reflect.TypeOf((*domain.Reporter)(nil)).Elem()
	reporter := app.Injector.GetByType(reporterType).(domain.Reporter)
	go reporter.Start()
	app.Run(config.ListenURL)
}
