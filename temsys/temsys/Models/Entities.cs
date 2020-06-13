using System;
using System.Collections.Generic;

namespace temsys.Models {
    public class Sensor {
        public Sensor(string name, string ip, int updateInterval, string[] reports, bool deleted) {
            this.Name = name;
            this.IP = ip;
            this.UpdateInterval = updateInterval;
            this.Reports = reports;
            this.Deleted = deleted;
        }

        public static Sensor CreateWithDefaults(string name, string ip) {
            return new Sensor(name, ip, 1, new string[0], false);
        }

        public override bool Equals(object obj) {
            return obj is Sensor sensor &&
                   Name == sensor.Name &&
                   IP == sensor.IP;
        }

        public override int GetHashCode() {
            return HashCode.Combine(Name, IP, UpdateInterval, Reports, Deleted);
        }

        public string Name {get;}

        public string IP {get;}

        public int UpdateInterval {get; set;}

        public string[] Reports {get;}

        public bool Deleted {get;}
    }

    public class Report {

        public Report(string type, string sensor, DateTime date, double value) {
            this.Type = type;
            this.Sensor = sensor;
            this.Date = date;
            this.Value = value;
        }

        public static Report CreateUsingFixedDate(string type, string sensor, double value) {
            return new Report(type, sensor, DateTime.MinValue, value);
        }

        public override bool Equals(object obj) {
            return obj is Report report &&
                   Type == report.Type &&
                   Sensor == report.Sensor &&
                   Date == report.Date &&
                   Value == report.Value;
        }

        public override int GetHashCode() {
            return HashCode.Combine(Type, Sensor, Date, Value);
        }

        public string Type {get;}
        public string Sensor {get;}
        public DateTime Date {get;}
        public double Value {get;}
    }
}