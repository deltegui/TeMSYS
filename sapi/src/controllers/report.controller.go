package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sensorapi/src/domain"
	"time"

	"github.com/deltegui/phoenix"
)

func GetAllReports(reportRepo domain.ReportRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.JSONPresenter{w}
		presenter.Present(reportRepo.GetAll())
	}
}

func GetReportsBetweenDates(getReportsByDate domain.GetReportsByDates) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.JSONPresenter{w}
		from, err := getDateFrom(req, "from")
		if err != nil {
			log.Println(err)
			presenter.PresentError(domain.MalformedRequestErr)
			return
		}
		to, err := getDateFrom(req, "to")
		if err != nil {
			log.Println(err)
			presenter.PresentError(domain.MalformedRequestErr)
			return
		}
		log.Println(from, to)
		getReportsByDate.Exec(presenter, domain.ReportsByDatesRequest{
			From: from,
			To:   to,
		})
	}
}

func getQueryFrom(req *http.Request, query string) (string, error) {
	key, ok := req.URL.Query()[query]
	if !ok || len(key[0]) < 1 {
		return "", fmt.Errorf("Key does not exist")
	}
	return key[0], nil
}

func getDateFrom(req *http.Request, query string) (time.Time, error) {
	dateStr, err := getQueryFrom(req, query)
	if err != nil {
		return time.Now(), err
	}
	layout := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Now(), err
	}
	return date, nil
}

func registerReportRoutes(app phoenix.App) {
	app.MapGroup("/routes", func(m phoenix.Mapper) {
		m.MapRoot(GetAllReports)
		m.Get("/dates", GetReportsBetweenDates)
	})
}
