package controllers

import (
	"encoding/json"
	"sensorapi/src/domain"

	"net/http"

	"github.com/deltegui/phoenix"
	"github.com/gorilla/mux"
)

func registerSensorRoutes(app phoenix.App) {
	app.MapGroup("/sensor", func(mapper phoenix.Mapper) {
		mapper.Post("", SaveSensor)
		mapper.Get("/{name}", GetSensorByName)
		mapper.Delete("/{name}", DeleteSensorByName)
		mapper.Post("/update", UpdateSensor)
		mapper.Get("/{name}/now", SensorNow)
	})
}

func SaveSensor(saveSensorCase domain.SaveSensorCase, builder domain.SensorBuilder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		var viewModel domain.SensorViewModel
		if err := json.NewDecoder(req.Body).Decode(&viewModel); err != nil {
			presenter.PresentError(domain.MalformedRequestErr)
			return
		}
		viewModel.SensorBuilder = builder
		saveSensorCase.Exec(presenter, viewModel)
	}
}

func GetSensorByName(getSensorCase domain.GetSensorCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		sensorName := mux.Vars(req)["name"]
		getSensorCase.Exec(presenter, sensorName)
	}
}

func DeleteSensorByName(deleteSensorCase domain.DeleteSensorCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		sensorName := mux.Vars(req)["name"]
		deleteSensorCase.Exec(presenter, sensorName)
	}
}

func UpdateSensor(updateSensorCase domain.UpdateSensorCase, builder domain.SensorBuilder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		var viewModel domain.SensorViewModel
		if err := json.NewDecoder(req.Body).Decode(&viewModel); err != nil {
			presenter.PresentError(domain.MalformedRequestErr)
			return
		}
		viewModel.SensorBuilder = builder
		updateSensorCase.Exec(presenter, viewModel)
	}
}

func SensorNow(sensorNowCase domain.SensorNowCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		sensorName := mux.Vars(req)["name"]
		sensorNowCase.Exec(presenter, sensorName)
	}
}
