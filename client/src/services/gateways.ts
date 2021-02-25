import {
  Sensor,
  Report,
  Token,
  UserResponse,
} from '@/services/models';

export interface SensorRepository {
  getAll(deleted: boolean): Promise<Sensor[]>;
  getByName(name: string): Promise<Sensor>;
  getCurrentStateByName(name: string, token: string): Promise<Report[]>;
  getCurrentAverageState(token: string): Promise<Report[]>;
  updateSensor(sensor: Sensor, token: string): Promise<Sensor>;
  delete(name: string, token: string): Promise<void>;
  create(sensor: Sensor, token: string): Promise<Sensor>;
}

export interface ReportRepository {
  getAll(name: string): Promise<Report[]>;
  getFiltered(options: {
    name: string;
    type?: string;
    trim?: number;
    fromDate?: Date;
    toDate?: Date;
    average?: boolean;
  }): Promise<Report[]>;
  getAllReportsAverage(options: {
    from: Date;
    to: Date;
  }): Promise<Report[]>;
  getAllReportTypes(): Promise<string[]>;
  saveReportType(name: string, token: string): Promise<boolean>;
}

export interface UserRepository {
  login(data: { name: string; password: string }): Promise<Token>;
  createUser(user: { name: string; password: string }, token: string): Promise<UserResponse>;
  deleteUser(user: string, token: string): Promise<string>;
  getAll(token: string): Promise<UserResponse[]>;
}

export interface TokenRepository {
  save(token: Token): void;
  load(): Token | undefined;
  clear(): void;
}
