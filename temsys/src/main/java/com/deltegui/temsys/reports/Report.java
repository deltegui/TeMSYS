package com.deltegui.temsys.reports;

import com.deltegui.temsys.sensors.domain.Sensor;

import java.time.LocalDateTime;

public final class Report {
    private String type;
    private double value;
    private LocalDateTime readDate;
    private String sensor;

    public Report(String type, double value, LocalDateTime readDate, String sensor) {
        this.type = type;
        this.value = value;
        this.readDate = readDate;
        this.sensor = sensor;
    }

    public String getType() {
        return type;
    }

    public double getValue() {
        return value;
    }

    public LocalDateTime getReadDate() {
        return readDate;
    }

    public String getSensor() {
        return sensor;
    }

    public boolean comesFromSensor(Sensor sensor) {
        return sensor.getName().equals(this.sensor);
    }
}
