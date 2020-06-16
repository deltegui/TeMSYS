using System.Threading.Tasks;
using System.Collections.Generic;

using temsys.Models;

namespace temsys.Services {
    public interface ISensorRepository {
        Task<IList<Sensor>> GetAllSensors(bool withDeleted = false);
        Task<IList<Report>> GetSensorsState();
        Task<Sensor> GetSensorByName(string name);
        Task<Sensor> SaveSensor(Sensor sensor);
        Task DeleteSensorByName(string name);
        Task<IList<Report>> GetOneSensorStateByName(string name);
    }
}