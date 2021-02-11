package main

import (
	"log"
	"net/http"
	"temsys"
	"temsys/api"
	"temsys/builders"
	"temsys/configuration"
	"temsys/cronscheluder"
	"temsys/hash"
	"temsys/mysql"
	"temsys/token"
	"temsys/validator"
	"time"

	phx "github.com/deltegui/phoenix"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
)

var optionsHandler http.HandlerFunc = func(w http.ResponseWriter, req *http.Request) {
	// Response CORS headers without content. Prevents CORS preflight
	// error in Firefox.
	w.WriteHeader(http.StatusNoContent)
}

type context struct {
	Config    configuration.Configuration
	DB        *sqlx.DB
	Validator temsys.Validator
	Tokenizer temsys.Tokenizer
	Auth      api.JWTAuth
	Reporter  temsys.Reporter
}

func main() {
	phx.PrintLogo("ascii-art.ans")
	config := configuration.Load()
	db := mysql.Connect(config)
	tokenizer := token.New(config.JWTKey)
	reporter := startReportScheluder(db)
	context := context{
		Config:    config,
		DB:        db,
		Validator: validator.New(),
		Tokenizer: tokenizer,
		Auth:      api.NewJWTAuth(tokenizer),
		Reporter:  reporter,
	}
	r := createRouter()
	mountRoutes(r, context)
	phx.FileServerStatic(r, "/")
	log.Printf("Listening on %s with tls? %t\n", config.ListenURL, config.TLSEnabled)
	log.Println("CORS allow origin: ", config.CORS)
	var err error
	if config.TLSEnabled {
		err = http.ListenAndServeTLS(config.ListenURL, config.TLSCRT, config.TLSKEY, r)
	} else {
		err = http.ListenAndServe(config.ListenURL, r)
	}
	if err != nil {
		log.Fatal("Error listening: ", err)
	}
}

func startReportScheluder(db *sqlx.DB) temsys.Reporter {
	sensorRepo := mysql.NewSensorRepo(db)
	reportRepo := mysql.NewReportRepo(db)
	cron := cronscheluder.New()
	reporter := temsys.NewReporter(sensorRepo, reportRepo, cron)
	go reporter.Start()
	return reporter
}

func createRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	return r
}

func mountRoutes(r chi.Router, ctx context) {
	r.Mount("/api", boostrapAPI(ctx))
}

func boostrapAPI(ctx context) chi.Router {
	r := chi.NewRouter()
	r.Use(api.NewCors(ctx.Config.CORS).EnableCors)
	r.Mount("/sensor", bootstrapSensor(ctx))
	r.Mount("/sensors", bootstrapSensors(ctx))
	r.Mount("/user", bootstrapUsers(ctx))
	r.Mount("/reports/types", bootstrapReportTypes(ctx))
	return r
}

func bootstrapSensor(ctx context) chi.Router {
	r := chi.NewRouter()
	sensorRepo := mysql.NewSensorRepo(ctx.DB)
	reportRepo := mysql.NewReportRepo(ctx.DB)

	getByNameCase := temsys.NewGetSensorCase(sensorRepo)
	sensorByNameHandler := api.GetSensorByNameHandler(getByNameCase)
	r.Get("/{name}", sensorByNameHandler)
	r.Options("/{name}", optionsHandler)

	nowCase := temsys.NewSensorNowCase(sensorRepo)
	// nowHandler := ctx.Auth.Authorize(api.SensorNowHandler(nowCase))
	nowHandler := api.SensorNowHandler(nowCase)
	r.Get("/{name}/now", nowHandler)
	r.Options("/{name}/now", optionsHandler)

	reportsBetweenDates := temsys.NewGetFilteredReports(ctx.Validator, reportRepo)
	reportsBetweenDatesHandler := api.GetReportsBetweenDatesHandler(reportsBetweenDates)
	r.Get("/{name}/reports", reportsBetweenDatesHandler)
	r.Options("/{name}/reports", optionsHandler)

	reportTypeRepo := mysql.NewReportTypeRepo(ctx.DB)
	sensorBuilder := builders.NewHttpSensorBuilder()

	saveSensor := temsys.NewSaveSensorCase(ctx.Validator, sensorRepo, ctx.Reporter, reportTypeRepo)
	r.Post("/", ctx.Auth.Admin(api.SaveSensorHandler(saveSensor, sensorBuilder)))

	deleteSensor := temsys.NewDeleteSensorCase(sensorRepo, ctx.Reporter)
	r.Delete("/{name}", ctx.Auth.Admin(api.DeleteSensorByNameHandler(deleteSensor)))

	updateSensor := temsys.NewUpdateSensorCase(ctx.Validator, sensorRepo, ctx.Reporter, reportTypeRepo)
	r.Patch("/", ctx.Auth.Admin(api.UpdateSensorHandler(updateSensor, sensorBuilder)))

	return r
}

func bootstrapSensors(ctx context) chi.Router {
	r := chi.NewRouter()
	sensorRepo := mysql.NewSensorRepo(ctx.DB)

	getAllSensorsCase := temsys.NewGetAllSensorsCase(sensorRepo)
	getAllHandler := api.GetAllSensorsHandler(getAllSensorsCase)
	r.Get("/", getAllHandler)
	r.Options("/", optionsHandler)

	allAverage := temsys.NewAllSensorNowAverageCase(sensorRepo)
	// allAverageHandler := ctx.Auth.Authorize(api.AllSensorsAverageHandler(allAverage))
	allAverageHandler := api.AllSensorsAverageHandler(allAverage)
	r.Get("/now/average", allAverageHandler)
	r.Options("/now/average", optionsHandler)

	return r
}

func bootstrapUsers(ctx context) chi.Router {
	r := chi.NewRouter()
	userRepo := mysql.NewUserRepo(ctx.DB)
	hasher := hash.BcryptPasswordHasher{}
	loginCase := temsys.NewLoginCase(ctx.Validator, userRepo, hasher, ctx.Tokenizer)
	r.Post("/login", api.LoginHandler(loginCase))
	r.Options("/login", optionsHandler)
	return r
}

func bootstrapReportTypes(ctx context) chi.Router {
	r := chi.NewRouter()
	reportTypeRepo := mysql.NewReportTypeRepo(ctx.DB)
	r.Get("/", api.GetReportTypesHandler(reportTypeRepo))
	r.Post("/{name}", ctx.Auth.Admin(api.SaveReportTypeHandler(reportTypeRepo)))
	return r
}
