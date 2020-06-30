package com.deltegui.temsys.sensors.application;

import com.deltegui.temsys.reports.Report;
import com.deltegui.temsys.sensors.domain.Sensor;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class SensorService {
    private final SensorRepository sensorRepository;

    public SensorService(SensorRepository sensorRepository) {
        this.sensorRepository = sensorRepository;
    }

    public List<SensorStatusModel> getAllSensorsWithStatus() {
        List<Sensor> sensors = this.sensorRepository.getAll();
        List<Report> reports = this.sensorRepository.getStatusForAll();
        return this.mergeSensorsWithReports(sensors, reports);
    }

    private List<SensorStatusModel> mergeSensorsWithReports(List<Sensor> sensors, List<Report> reports) {
        var models = new ArrayList<SensorStatusModel>();
        for(Sensor currentSensor : sensors) {
            List<Report> filteredReports = currentSensor.filterReports(reports);
            models.add(new SensorStatusModel(currentSensor, filteredReports));
        }
        return models;
    }

}
