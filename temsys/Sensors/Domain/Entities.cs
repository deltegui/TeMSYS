using System;

namespace temsys.Sensors.Domain {
    public struct Sensor {
        string Name;
        string IP;
        int UpdateInterval;
        string[] Reports;
        bool Deleted;
    }

    public struct Report {
        string Type;
        string Snesor;
        DateTime Date;
        double Value;
    }
}