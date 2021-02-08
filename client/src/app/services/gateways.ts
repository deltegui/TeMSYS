import {
  Sensor,
  Report,
} from './models';

export interface SensorRepository {
  getAll(): Promise<Sensor[]>
  getByName(name: string): Promise<Sensor>
  getCurrentStateByName(name: string): Promise<Report[]>
  getCurrentAverageState(): Promise<Report[]>
}

export interface ReportRepository {
  getAll(name: string): Promise<Report[]>
  getByDateRange(
    {
      name,
      fromDate,
      toDate,
      average
    }:
    {
      name: string,
      fromDate: Date,
      toDate: Date,
      average: boolean,
    },
  ): Promise<Report[]>
  getLatestReports(
    {
      name,
      trim,
      type
    }:
    {
      name: string,
      trim: number,
      type: string
    },
  ): Promise<Report[]>
  getByDate(
    {
      name,
      date,
      average
    }:
    {
      name: string,
      date: Date,
      average: boolean,
    },
  ): Promise<Report[]>
}