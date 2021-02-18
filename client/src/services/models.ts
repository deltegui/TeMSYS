export type Sensor = {
  name: string;
  connection: {
    type: string;
    value: string;
  };
  updateInterval: number;
  deleted: boolean;
  supportedReports: string[];
}

export type Report = {
  type: string;
  sensor: string;
  date: Date;
  value: number;
}

export type Token = {
  value: string;
  expires: Date;
  owner: string;
  role: string;
}

export type UserResponse = {
  name: string;
  role: string;
}

export type ReportFilter = {
  name: string;
  type?: string;
  trim?: number;
  fromDate?: Date;
  toDate?: Date;
  average?: boolean;
}
