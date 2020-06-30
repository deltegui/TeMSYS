package com.deltegui.temsys.sensors.application;

import com.deltegui.temsys.reports.Report;
import com.deltegui.temsys.sensors.domain.Sensor;

import java.util.List;

public class SensorStatusModel {
    private Sensor sensor;
    private List<Report> reports;

    public SensorStatusModel(Sensor sensor, List<Report> currentStatus) {
        this.sensor = sensor;
        this.reports = currentStatus;
    }

    public String getSensorName() {
        return this.sensor.getName();
    }

    public String getSensorIp() {
        return this.sensor.getConnectionValue();
    }

    public List<Report> getReports() {
        return reports;
    }

    public boolean haveReports() {
        return !this.reports.isEmpty();
    }
}
