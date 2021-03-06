/* eslint-disable class-methods-use-this */
/* eslint-disable no-useless-constructor */
import { store } from '@/store';
import { ReportRepository, SensorRepository } from './gateways';
import { Report, ReportFilter, Sensor } from './models';

function hoursFromNow(hours: number) {
  const date = new Date();
  date.setHours(date.getHours() - hours);
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
    const from = hoursFromNow(2);
    const to = new Date();
    return this.reportRepo.getAllReportsAverage({ from, to });
  }

  async getAllReportTypes(): Promise<string[]> {
    return this.reportRepo.getAllReportTypes();
  }

  async saveReportType(name: string): Promise<boolean> {
    return this.reportRepo.saveReportType(name, store.token?.value ?? '');
  }

  async getLastReadForSensor(name: string): Promise<Report[]> {
    const from = hoursFromNow(2);
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

  async getLatestReports({
    name,
    trim = 48,
    type = undefined,
  }: {
    name: string;
    trim: number;
    type?: string;
  }): Promise<Report[]> {
    const from = hoursFromNow(trim);
    const to = new Date();
    return this.reportRepo.getFiltered({
      name,
      fromDate: from,
      toDate: to,
      trim,
      type,
    });
  }

  roundReportValue(value: number): number {
    return Math.round(value * 100) / 100;
  }

  roundAllReports(report: Report[]): Report[] {
    return report.map((r) => ({
      type: r.type,
      sensor: r.sensor,
      date: r.date,
      value: this.roundReportValue(r.value),
    }));
  }

  readReportsFiltered(filters: ReportFilter): Promise<Report[]> {
    return this.reportRepo.getFiltered(filters);
  }
}
