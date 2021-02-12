/* eslint-disable class-methods-use-this */
/* eslint-disable no-param-reassign */

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

  async getByDateRange(
    {
      name,
      fromDate,
      toDate,
      average,
    }:
    {
      name: string;
      fromDate: Date;
      toDate: Date;
      average: boolean;
    },
  ): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports?from=${fromDate}&to=${toDate}&average=${average}`)
      .then(passReportsToRealDate);
  }

  async getLatestReports(
    {
      name,
      trim,
      type,
    }:
    {
      name: string;
      trim: number;
      type: string;
    },
  ): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports?trim=${trim}&type=${type}`)
      .then(passReportsToRealDate);
  }

  async getByDate(
    {
      name,
      date,
      average,
    }:
    {
      name: string;
      date: Date;
      average: boolean;
    },
  ): Promise<Report[]> {
    return makeRequest(`/sensor/${name}/reports?from=${date}&to=${date}&average=${average}`)
      .then(passReportsToRealDate);
  }
}
