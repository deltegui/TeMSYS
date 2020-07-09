const reportTypes = [
  'temperature',
  'humidity',
];

function splitReportsByType(reports) {
  const splitted = [];
  for(let i = 0; i < reportTypes.length; i++) {
    splitted.push(reports.filter(({ type }) => type === reportTypes[i]));
  }
  return splitted;
}

function reduceToAverage(reports) {
  const len = reports.length;
  const { type, sensor } = reports[0];
  const sum = reports.reduce((prev, { value }) => prev + value, 0);
  return {
    type,
    sensor,
    value: sum / len,
    date: new Date(),
  };
}

class ReportService {
  constructor(reportRepo) {
    this.reportRepo = reportRepo;
  }

  getReportsForSensor(name, options = {}) {
    return this.reportRepo.getReportsForSensor(name, options)
      .then(arr => arr.map(({ type, sensor, date, value }) => ({ type, sensor, date, value })))
      .then(reports => {
        if(options.average) return this.calculateReportAverage(reports);
        return reports;
      });
  }

  calculateReportAverage(reports) {
    const splitted = splitReportsByType(reports);
    return splitted.map(reduceToAverage);
  }
}

module.exports = ReportService;
