/* eslint-disable no-param-reassign */
/* eslint-disable no-plusplus */

import { Report } from '@/services/models';

type DataEntry = {
  label: string;
  yAxisID?: string;
  data: number;
  backgroundColor: string | void;
  borderColor: string | void;
  fill: boolean;
}

export type ChartData = {
  datasets: DataEntry[] | DataEntry[][];
  labels: string[];
  tooltips?: string[][];
}

function dateIsEqual(first: Date, second: Date): boolean {
  return first.getFullYear() === second.getFullYear()
    && first.getMonth() === second.getMonth()
    && first.getDay() === second.getDay()
    && first.getHours() === second.getHours();
}

function groupReportsByDate(reports: Report[]): Report[][] {
  const groups = [];
  for (let i = 0; i < reports.length; i++) {
    const report = reports[i];
    let found = false;
    for (let j = 0; j < groups.length; j++) {
      const grp = groups[j];
      if (dateIsEqual(grp[0].date, report.date)) {
        grp.push(report);
        found = true;
      }
    }
    if (!found) {
      groups.push([report]);
    }
  }
  return groups;
}

function calculateAverageByGroup(groups: Report[][]) {
  const averageReports = [];
  for (let j = 0; j < groups.length; j++) {
    const grp = groups[j];
    const sum = grp.reduce((prev, current) => prev + current.value, 0);
    const value = sum / grp.length;
    const { type, date } = grp[0];
    averageReports.push({
      type,
      date,
      sensor: 'average',
      value,
    });
  }
  return averageReports;
}

function compareReportsByDate(a: Report, b: Report): number {
  return a.date.getTime() - b.date.getTime();
}

class ChartColors {
  private colors = [
    '#b35a41',
    '#0269A4',
    '#87E698',
    '#E566E5',
  ];

  private currentGenerator: Generator<string, void, unknown>;

  constructor() {
    this.currentGenerator = this.generator();
  }

  * generator() {
    let currentIndex = 0;
    for (;;) {
      yield this.colors[currentIndex];
      currentIndex = (currentIndex + 1) % this.colors.length;
    }
  }

  next(): IteratorResult<string> {
    return this.currentGenerator.next();
  }

  getColor(reportType = ''): string {
    switch (reportType) {
      case 'temperature': return this.colors[0];
      case 'humidity': return this.colors[1];
      default: return this.next().value;
    }
  }
}

function generateDataSets(separatedReports: any): DataEntry[] {
  const colors = new ChartColors();
  return Object
    .keys(separatedReports)
    .map((key) => ({
      key,
      label: separatedReports[key][0].type,
      data: separatedReports[key].map(({ value }: { value: string }) => value),
    }))
    .map(({ key, label, data }) => {
      const color = colors.getColor(key);
      return {
        label,
        yAxisID: label,
        data,
        backgroundColor: color,
        borderColor: color,
        fill: false,
      };
    });
}

function generateLabels(separatedReports: any): string[] {
  if (Object.keys(separatedReports).length === 0) return [];
  const firstGroup = separatedReports[Object.keys(separatedReports)[0]];
  return firstGroup.map((e: Report) => e.date.getHours());
}

function generateTooltips(separatedReports: any): string[][] {
  return Object.keys(separatedReports)
    .map((key) => {
      const firstGroup = separatedReports[key];
      return firstGroup.map((e: Report) => `(${key}) ${e.date.toDateString()} ${e.date.getHours()}:${e.date.getMinutes()}:${e.date.getSeconds()}`);
    });
}

function toChartData(separatedReports: any): ChartData {
  return {
    datasets: generateDataSets(separatedReports),
    labels: generateLabels(separatedReports),
    tooltips: generateTooltips(separatedReports),
  };
}

function separateReportsByType(reports: Report[]): any {
  return reports.reduce((prev: any, report: Report) => {
    if (!prev[report.type]) {
      prev[report.type] = [];
    }
    prev[report.type].push(report);
    return prev;
  }, {});
}

function separateReportsBySensorAndType(reportPerSensor: Report[][]): any {
  return reportPerSensor.reduce((prev: any, r: Report[]) => {
    if (r.length === 0) {
      return prev;
    }
    const first = r[0];
    const key = `${first.sensor}, ${first.type}`;
    if (!prev[key]) {
      prev[key] = [];
    }
    prev[key] = r;
    return prev;
  }, {});
}

export default {
  sortReports(reports: Report[]): Report[] {
    return reports.sort(compareReportsByDate);
  },

  flatAndSort(reports: Report[][]): Report[] {
    return this.sortReports(reports.flat());
  },

  sortSubets(reports: Report[][]): Report[][] {
    return reports.map(this.sortReports.bind(this));
  },

  calculateElementsByChart(spacingBetweenChartElements = 40) {
    return parseInt(String(window.innerWidth / spacingBetweenChartElements), 10);
  },

  genearateDataSetsForOneSensor(sensorReports: Report[]): ChartData {
    return toChartData(separateReportsByType(sensorReports));
  },

  generateDataSetsForSensors(reports: Report[][]): ChartData {
    return toChartData(separateReportsBySensorAndType(reports));
  },

  getAverageAllReports(reportsPerSensor: Report[][]) {
    const flattenAndSorted = this.flatAndSort(reportsPerSensor);
    const separatedByType = separateReportsByType(flattenAndSorted);
    const typesAndGrouped = Object.keys(separatedByType)
      .map((reportType) => groupReportsByDate(separatedByType[reportType]));
    return typesAndGrouped.map(calculateAverageByGroup);
  },
};
