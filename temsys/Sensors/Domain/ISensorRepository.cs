using System.Threading.Tasks;

namespace temsys.Sensors.Domain {
    public interface ISensorRepository {
        Task<Sensor[]> GetAllSensors(bool withDeleted = false);
        Report[] GetSensorsState();
        Sensor GetSensorByName(string name);
        void SaveSensor(Sensor sensor);
        void DeleteSensorByName(string name);
        Report[] GetOneSensorStateByName(string name);
    }

    public interface IReportTypeRepository {
        string[] GetAllReportTypes();
        void SaveReportType(string reportType);
    }
}