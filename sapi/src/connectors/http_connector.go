package connectors

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sensorapi/src/domain"
	"time"
)

type HTTPConnector struct {
	IP string
}

func (connector HTTPConnector) ReadDataFor(sensor domain.Sensor) ([]domain.Report, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s", connector.IP), nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return []domain.Report{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return []domain.Report{}, err
	}
	var data map[string]float32
	json.Unmarshal(body, &data)
	var reports []domain.Report
	for key, value := range data {
		reports = append(reports, domain.Report{
			ReportType: key,
			SensorName: sensor.Name,
			Date:       time.Now(),
			Value:      value,
		})
	}
	return reports, nil
}
