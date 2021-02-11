import {
  Sensor,
  Report,
  User,
  Token,
} from './models';

export interface SensorRepository {
  getAll(): Promise<Sensor[]>
  getByName(name: string): Promise<Sensor>
  getCurrentStateByName(name: string): Promise<Report[]>
  getCurrentAverageState(): Promise<Report[]>
}

export interface ReportRepository {
  getAll(name: string): Promise<Report[]>
  getByDateRange(range: { name: string, fromDate: Date, toDate: Date, average: boolean }): Promise<Report[]>
  getLatestReports(info: { name: string, trim: number, type: string }): Promise<Report[]>
  getByDate(info : { name: string, date: Date, average: boolean }): Promise<Report[]>
}

export interface UserRepository {
  login(data: { name: string, password: string }): Promise<User>,
}

export interface TokenRepository {
  save(token: Token): void;
  load(): Token;
}
