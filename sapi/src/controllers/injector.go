package controllers

import (
	"sensorapi/src/builders"
	"sensorapi/src/configuration"
	"sensorapi/src/cronscheluder"
	"sensorapi/src/domain"
	"sensorapi/src/persistence"
	"sensorapi/src/validator"

	"github.com/deltegui/phoenix"
)

func registerUseCases(app phoenix.App) {
	app.Injector.Add(domain.NewGetAllSensorsCase)
	app.Injector.Add(domain.NewAllSensorNowCase)
	app.Injector.Add(domain.NewDeleteSensorCase)
	app.Injector.Add(domain.NewGetSensorCase)
	app.Injector.Add(domain.NewSaveSensorCase)
	app.Injector.Add(domain.NewSensorNowCase)
	app.Injector.Add(domain.NewUpdateSensorCase)
	app.Injector.Add(domain.NewGetReportsByDates)
}

func registerDependencies(app phoenix.App) {
	app.Injector.Add(builders.NewHttpSensorBuilder)
	app.Injector.Add(validator.NewPlaygroundValidator)
	app.Injector.Add(persistence.NewSqlxReportTypeRepo)
	app.Injector.Add(persistence.NewSqlxSensorRepo)
	app.Injector.Add(persistence.NewSqlxReportRepo)
	app.Injector.Add(domain.NewReporter)
	app.Injector.Add(cronscheluder.NewCronScheluder)
}

func Register(app phoenix.App, config configuration.Configuration) {
	registerUseCases(app)
	registerDependencies(app)
	conn := persistence.NewSqlxConnection(config)
	app.Injector.Add(func() *persistence.SqlxConnection { return &conn })
	app.Injector.Add(func() configuration.Configuration { return config })

	app.Get("404", NotFound)
	registerReportTypesRoutes(app)
	registerSensorsRoutes(app)
	registerSensorRoutes(app)
	registerReportRoutes(app)
}
