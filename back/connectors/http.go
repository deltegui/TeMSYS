package connectors

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"temsys"
	"time"
)

// HTTPConnector is an implementation of a Sensor Connector using HTTP protocol.
type HTTPConnector struct {
	IP string
}

// ReadDataFor sensor. Returns a bunch of reports with current information readed from
// sensor. Returns an error if it cannot read information from sensor
func (connector HTTPConnector) ReadDataFor(sensor temsys.Sensor) ([]temsys.Report, error) {
	body, err := connector.readBodyFromSensor(sensor)
	if err != nil {
		return nil, err
	}
	reports := connector.parseSensorResponse(body, sensor)
	return reports, nil
}

func (connector HTTPConnector) readBodyFromSensor(sensor temsys.Sensor) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s", connector.IP), nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 20*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (connector HTTPConnector) parseSensorResponse(body []byte, sensor temsys.Sensor) []temsys.Report {
	var data map[string]float32
	json.Unmarshal(body, &data)
	var reports []temsys.Report
	for key, value := range data {
		reports = append(reports, temsys.Report{
			ReportType: key,
			SensorName: sensor.Name,
			Date:       time.Now(),
			Value:      value,
		})
	}
	return reports
}
