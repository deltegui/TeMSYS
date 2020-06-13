using System;
using System.Diagnostics;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Linq;
using Microsoft.AspNetCore.Mvc;
using temsys.Models;
using temsys.Services;

namespace temsys.Controllers {
    public class HomeController : Controller {
        private readonly ISensorRepository sensorRepository;

        public HomeController(ISensorRepository sensorRepository) {
            this.sensorRepository = sensorRepository;
        }

        public async Task<IActionResult> Index() {
            IList<Report> reportsList = await this.sensorRepository.GetOneSensorStateByName("salon");
            foreach(var r in reportsList) {
                Console.WriteLine(r.Value);
            }
            return View();
        }

        public IActionResult Privacy() {
            return View();
        }

        [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
        public IActionResult Error() {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}
