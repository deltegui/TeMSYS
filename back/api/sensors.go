package api

import (
	"encoding/json"
	"net/http"

	"temsys"

	"github.com/deltegui/phoenix"
	"github.com/go-chi/chi"
)

func SaveSensorHandler(saveSensorCase temsys.UseCase, builder temsys.SensorBuilder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		var sensorReponse temsys.SensorResponse
		if err := json.NewDecoder(req.Body).Decode(&sensorReponse); err != nil {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		sensorReponse.SensorBuilder = builder
		saveSensorCase.Exec(presenter, sensorReponse)
	}
}

func DeleteSensorByNameHandler(deleteSensorCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		sensorName := chi.URLParam(req, "name")
		if len(sensorName) == 0 {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		deleteSensorCase.Exec(presenter, sensorName)
	}
}

func UpdateSensorHandler(updateSensorCase temsys.UseCase, builder temsys.SensorBuilder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		var viewModel temsys.SensorResponse
		if err := json.NewDecoder(req.Body).Decode(&viewModel); err != nil {
			presenter.PresentError(temsys.MalformedRequestErr)
			return
		}
		viewModel.SensorBuilder = builder
		updateSensorCase.Exec(presenter, viewModel)
	}
}

func GetSensorByNameHandler(getSensorCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		sensorName := chi.URLParam(req, "name")
		getSensorCase.Exec(presenter, sensorName)
	}
}

func SensorNowHandler(sensorNowCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		sensorName := chi.URLParam(req, "name")
		sensorNowCase.Exec(presenter, sensorName)
	}
}

func GetAllSensorsHandler(getAllSensorsCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		wantDeleted := req.URL.Query()["deleted"]
		var reqCase temsys.GetAllSensorsRequest
		reqCase.WantDeleted = !(len(wantDeleted) < 1 || len(wantDeleted[0]) == 0 || wantDeleted[0] == "false")
		getAllSensorsCase.Exec(presenter, reqCase)
	}
}

func AllSensorNowHandler(allSensorNowCase temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		allSensorNowCase.Exec(presenter, temsys.EmptyRequest)
	}
}

func AllSensorsAverageHandler(allSensorsNowAverage temsys.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		presenter := phoenix.NewJSONPresenter(w)
		allSensorsNowAverage.Exec(presenter, temsys.EmptyRequest)
	}
}
