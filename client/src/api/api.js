export class SensorApi {

  constructor(api) {
    this.api = api;
  }

  getAll() {
    return this.api.makeRequest('/sensors');
  }

  getByName(name) {
    return this.api.makeRequest(`/sensor/${name}`);
  }

  getCurrentStateByName(name) {
    return this.api.makeRequest(`/sensor/${name}/now`);
  }

  getCurrentAverageState() {
    return this.api.makeRequest('/sensors/now/average');
  }
}

export class ReportApi {

  constructor(api) {
    this.api = api;
  }

  getAll(name) {
    return this.api.makeRequest(`/sensor/${name}/reports`)
      .then(passReportsToRealDate);
  }

  getByDateRange({name, fromDate, toDate, average}) {
    return this.api.makeRequest(`/sensor/${name}/reports`, {
      date: {
        from: fromDate,
        to: toDate,
      },
      average,
    }).then(passReportsToRealDate);
  }

  getLatestReports({name, trim, type}) {
    return this.api.makeRequest(`/sensor/${name}/reports`, {
      latestReports: trim,
      type,
    }).then(passReportsToRealDate);
  }

  getByDate({name, date, average}) {
    return this.api.makeRequest(`/sensor/${name}/reports`, {
      date: {
        from: date,
        to: date,
      },
      average,
    }).then(passReportsToRealDate);
  }

}

function passReportsToRealDate(reports) {
  return reports.map(r => {
    r.date = new Date(r.date)
    return r;
  });
}