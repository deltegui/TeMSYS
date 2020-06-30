package com.deltegui.temsys.sensors.application;

import com.deltegui.temsys.reports.Report;

import java.util.List;

public class SensorStatusModel {
    private String sensorName;
    private List<Report> reports;

    public SensorStatusModel(String sensorName, List<Report> currentStatus) {
        this.sensorName = sensorName;
        this.reports = currentStatus;
    }

    protected void addReport(Report report) {
        this.reports.add(report);
    }

    public String getSensorName() {
        return sensorName;
    }

    public List<Report> getReports() {
        return reports;
    }

    public boolean haveReports() {
        return !this.reports.isEmpty();
    }
}
