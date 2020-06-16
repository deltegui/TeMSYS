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
}