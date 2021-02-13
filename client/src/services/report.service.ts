/* eslint-disable no-useless-constructor */
import { ReportRepository, SensorRepository } from './gateways';
import { Report, Sensor } from './models';

function twoHoursFromNow() {
  const date = new Date();
  date.setHours(date.getHours() - 2);
  return date;
}

function getNumberSupportedReportTypes(sensor: Sensor) {
  return sensor.supportedReports.length;
}

export default class ReportService {
  constructor(
    private reportRepo: ReportRepository,
    private sensorRepo: SensorRepository,
  ) {}

  async getLastReadAverage(): Promise<Report[]> {
    const from = twoHoursFromNow();
    const to = new Date();
    return this.reportRepo.getAllReportsAverage({ from, to });
  }

  async getLastReadForSensor(name: string): Promise<Report[]> {
    const from = twoHoursFromNow();
    const to = new Date();
    return this.sensorRepo.getByName(name)
      .then(getNumberSupportedReportTypes)
      .then((trim) => this.reportRepo.getFiltered({
        name,
        fromDate: from,
        toDate: to,
        trim,
      }));
  }
}
