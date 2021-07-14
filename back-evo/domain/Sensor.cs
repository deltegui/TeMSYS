using System.Collections.Generic;
using Domain.Core;

namespace Domain {
    public interface ISensorConnector {
        public Maybe<List<Report>> ReadDataFrom(string connectionValue);
    }

    public class Sensor {
        public string Name { get; }
        public int UpdateInterval { get; set; }
        public bool Deleted { get; set; }
        public List<string> SupportedReports { get; set; }
        public string ConnectionValue { get; set; }

        public Sensor (
            string name,
            int updateInterval,
            bool deleted,
            List<string> supportedReports,
            string connValue
        ) {
            this.Name = name;
            this.UpdateInterval = updateInterval;
            this.Deleted = deleted;
            this.SupportedReports = supportedReports;
            this.ConnectionValue = connValue;
        }

        public Maybe<List<Report>> GetCurrentState(ISensorConnector connector) {
            for (int i = 0; i < 2; i++) {
                var reports = connector.ReadDataFrom(this.ConnectionValue);
                if (reports.IsFilled()) {
                    return reports;
                }
            }
            return Maybe<List<Report>>.Empty();
        }
    }
}
