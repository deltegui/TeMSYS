package com.deltegui.temsys.sensors.domain;

import com.deltegui.temsys.reports.Report;

import java.util.List;
import java.util.stream.Collectors;

public final class Sensor {
    private String name;
    private String connectionType;
    private String connectionValue;
    private int upateInterval;
    private boolean deleted;

    public Sensor(String name, String connectionType, String connectionValue, int updateInterval, boolean deleted) {
        this.name = name;
        this.connectionValue = connectionValue;
        this.connectionType = connectionType;
        this.upateInterval = updateInterval;
        this.deleted = deleted;
    }

    public List<Report> filterReports(List<Report> reports) {
        return reports.stream()
                .filter(report -> report.comesFromSensor(this))
                .collect(Collectors.toList());
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getConnectionValue() {
        return connectionValue;
    }

    public void setConnectionValue(String connectionValue) {
        this.connectionValue = connectionValue;
    }

    public String getConnectionType() {
        return connectionType;
    }

    public void setConnectionType(String connectionType) {
        this.connectionType = connectionType;
    }

    public int getUpateInterval() {
        return upateInterval;
    }

    public void setUpateInterval(int upateInterval) {
        this.upateInterval = upateInterval;
    }

    public boolean isDeleted() {
        return deleted;
    }

    public void setDeleted(boolean deleted) {
        this.deleted = deleted;
    }
}
