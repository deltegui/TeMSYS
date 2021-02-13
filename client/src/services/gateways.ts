import {
  Sensor,
  Report,
  Token,
} from '@/services/models';

export interface SensorRepository {
  getAll(): Promise<Sensor[]>;
  getByName(name: string): Promise<Sensor>;
  getCurrentStateByName(name: string, token: string): Promise<Report[]>;
  getCurrentAverageState(token: string): Promise<Report[]>;
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
}

export interface UserRepository {
  login(data: { name: string; password: string }): Promise<Token>;
}

export interface TokenRepository {
  save(token: Token): void;
  load(): Token | undefined;
  clear(): void;
}
