using System.Collections.Generic;
using System.Threading.Tasks;

namespace temsys.Services {
    public interface IReportTypeRepository {
        Task<IList<string>> GetAllReportTypes();
        Task SaveReportType(string reportType);
    }
}
