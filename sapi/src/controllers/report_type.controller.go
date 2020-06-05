package controllers

import (
	"log"
	"sensorapi/src/domain"

	"net/http"

	"github.com/deltegui/phoenix"
	"github.com/gorilla/mux"
)

func GetReportTypes(reportTypeRepo domain.ReportTypeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		phoenix.NewJSONPresenter(w).Present(reportTypeRepo.GetAll())
	}
}

func SaveReportType(reportTypeRepo domain.ReportTypeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.JSONPresenter{w}
		reportType := domain.ReportType(mux.Vars(req)["name"])
		if err := reportTypeRepo.Save(reportType); err != nil {
			log.Println(err)
			presenter.PresentError(domain.MalformedRequestErr)
			return
		}
		presenter.Present(struct{ ReportType domain.ReportType }{reportType})
	}
}

func registerReportTypesRoutes(app phoenix.App) {
	app.MapGroup("/report/types", func(m phoenix.Mapper) {
		m.Get("/all", GetReportTypes)
		m.Post("/create/{name}", SaveReportType)
	})
}
