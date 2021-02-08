package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"temsys"

	"github.com/deltegui/phoenix"
	"github.com/go-chi/chi"
)

func GetReportTypesHandler(reportTypeRepo temsys.ReportTypeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		phoenix.NewJSONPresenter(w).Present(reportTypeRepo.GetAll())
	}
}

func SaveReportTypeHandler(reportTypeRepo temsys.ReportTypeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		reportType := chi.URLParam(req, "name")
		if len(reportType) == 0 {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		if err := reportTypeRepo.Save(reportType); err != nil {
			log.Println(err)
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		presenter.Present(struct{ ReportType string }{reportType})
	}
}

func GetReportsBetweenDatesHandler(getReportsByDate temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		name := chi.URLParam(req, "name")
		if len(name) == 0 {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		from, to := readFromTo(req)
		average := readAverage(req)
		trim := readTrim(req)
		typ := readType(req)
		getReportsByDate.Exec(presenter, temsys.FilteredReportsRequest{
			From:       from,
			To:         to,
			Average:    average,
			Trim:       trim,
			Type:       typ,
			SensorName: name,
		})
	}
}

func readType(req *http.Request) string {
	typ, err := getQueryFrom(req, "type")
	if err != nil {
		typ = ""
	}
	log.Println("Type: ", typ)
	return typ
}

func readTrim(req *http.Request) int {
	trimStr, err := getQueryFrom(req, "trim")
	if err != nil {
		trimStr = "100"
	}
	trim, err := strconv.Atoi(trimStr)
	if err != nil {
		trim = 100
	}
	log.Println("Trim: ", trim)
	return trim
}

func readAverage(req *http.Request) bool {
	avgStr, err := getQueryFrom(req, "average")
	average := err == nil && avgStr == "true"
	log.Println("Average: ", average)
	return average
}

func readFromTo(req *http.Request) (from time.Time, to time.Time) {
	from, err := getDateFrom(req, "from")
	if err != nil {
		log.Println(err)
		from = time.Unix(0, 0)
	}
	to, err = getDateFrom(req, "to")
	if err != nil {
		log.Println(err)
		to = time.Now()
	}
	log.Println(from, to)
	return from, to
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

func getQueryFrom(req *http.Request, query string) (string, error) {
	key, ok := req.URL.Query()[query]
	if !ok || len(key[0]) < 1 {
		return "", fmt.Errorf("Key does not exist")
	}
	return key[0], nil
}
