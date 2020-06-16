using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace temsys.Models {
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

        public string Type { get; }
        public string Sensor { get; }
        public DateTime Date { get; }
        public double Value { get; }
    }
}
