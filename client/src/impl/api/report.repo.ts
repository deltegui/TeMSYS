/* eslint-disable class-methods-use-this */
/* eslint-disable no-param-reassign */
/* eslint-disable @typescript-eslint/no-explicit-any */

import {
  Report,
} from '@/services/models';

import {
  ReportRepository,
} from '@/services/gateways';

import makeRequest from './core';

function passReportsToRealDate(reports: any[]): Report[] {
  return reports.map((r) => {
    r.date = new Date(r.date);
    return r;
  });
}

export default class ApiReportRepository implements ReportRepository {
  async getAll(name: string): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports`)
      .then(passReportsToRealDate);
  }

  async getFiltered(
    {
      name,
      type,
      trim,
      fromDate,
      toDate,
      average,
    }:
    {
      name: string;
      type?: string;
      trim?: number;
      fromDate?: Date;
      toDate?: Date;
      average?: boolean;
    },
  ): Promise<Report[]> {
    let baseQuery = `/sensor/${name}/reports?`;
    if (fromDate && toDate) {
      baseQuery += `from=${fromDate.toJSON()}&to=${toDate.toJSON()}&`;
    }
    if (trim) {
      baseQuery += `trim=${trim}&`;
    }
    if (average) {
      baseQuery += `average=${average}&`;
    }
    if (type) {
      baseQuery += `type=${type}&`;
    }
    return makeRequest(baseQuery)
      .then(passReportsToRealDate);
  }

  async getAllReportsAverage(
    {
      from,
      to,
    }:
    {
      from: Date;
      to: Date;
    },
  ): Promise<Report[]> {
    let baseQuery = '/reports/average';
    if (from && to) {
      baseQuery += `?from=${from.toJSON()}&to=${to.toJSON()}`;
    }
    return makeRequest(baseQuery)
      .then(passReportsToRealDate);
  }

  async getAllReportTypes(): Promise<string[]> {
    return makeRequest('/reports/types');
  }
}
