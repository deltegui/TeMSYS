export type Sensor = {
  name: string,
  connection: {
    connType: string,
    connValue: string,
  },
  updateInterval: number,
  deleted: boolean,
  supportedReports: string[],
}

export type Report = {
  type: string,
  sensor: string,
  date: Date,
  value: number,
}

export type Token = {
  value: string,
  expires: Date,
  owner: string,
  role: string,
}

export type User = {
  name: string,
  role: string,
  token: Token,
}