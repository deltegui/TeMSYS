import { environment } from '../../environments/environment';

import {
  Sensor,
  Report,
  User,
} from './models';
import {
  SensorRepository,
  ReportRepository,
  UserRepository,
} from './gateways';

function isApiError(err: any): boolean {
  return (!!err.Code || err.Code === 0) && !!err.Reason;
}

async function makeRequest(endpoint: string, body: any = undefined, method: string = 'GET'): Promise<any> {
  let reqConfig = undefined;
  if (body) {
    reqConfig = {
      method,
      body: JSON.stringify(body),
      headers: {
        'Content-Type': 'application/json',
      },
    };
  }
  return fetch(`${environment.apiURL}${endpoint}`, reqConfig)
    .then(async (res) => {
      if (res.ok) {
        return res.json();
      }
      throw await res.json()
    })
    .catch((err) => {
      if (isApiError(err)) {
        throw {
          code: err.Code,
          reason: err.Reason,
        };
      }
      throw {
        code: -1,
        reason: err.toString(),
      };
    });
}

export class ApiSensorRepository implements SensorRepository {
  async getAll(): Promise<Sensor[]> {
    return makeRequest('/sensors');
  }

  async getByName(name: string): Promise<Sensor> {
    return makeRequest(`/sensor/${name}`);
  }

  async getCurrentStateByName(name: string): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/now`);
  }

  async getCurrentAverageState(): Promise<Report[]> {
    return makeRequest('/sensors/now/average');
  }
}

export class ApiReportRepository implements ReportRepository {
  async getAll(name: string): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports`)
      .then(passReportsToRealDate);
  }

  async getByDateRange(
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
  ): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports?from=${fromDate}&to=${toDate}&average=${average}`)
      .then(passReportsToRealDate);
  }

  async getLatestReports(
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
  ): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports?trim=${trim}&type=${type}`)
      .then(passReportsToRealDate);
  }

  async getByDate(
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
  ): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports?from=${date}&to=${date}&average=${average}`)
      .then(passReportsToRealDate);
  }
}

function passReportsToRealDate(reports: any[]): Report[] {
  return reports.map(r => {
    r.date = new Date(r.date)
    return r;
  });
}

export class ApiUserRepository implements UserRepository {
  async login(body: { name: string, password: string }): Promise<User> {
    return makeRequest('/user/login', body, 'POST')
      .then(({ name, role, token }) => {
        return {
          name,
          role,
          token: {
            value: token.value,
            expires: token.expires,
            owner: token.owner,
            role: token.role,
          },
        };
      })
  }
}