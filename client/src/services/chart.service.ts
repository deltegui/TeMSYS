/* eslint-disable no-param-reassign */

import { Report } from '@/services/models';

type DataEntry = {
  label: string;
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

function compareReportsByDate(a: Report, b: Report): number {
  return a.date.getTime() - b.date.getTime();
}

function* generateColor() {
  const chartColors = [
    '#b35a41',
    '#0269A4',
    '#87E698',
    '#E566E5',
  ];
  let currentIndex = 0;
  for (;;) {
    yield chartColors[currentIndex];
    currentIndex = (currentIndex + 1) % chartColors.length;
  }
}

function generateDataSets(separatedReports: any): DataEntry[] {
  const colorGenerator = generateColor();
  return Object
    .keys(separatedReports)
    .map((key) => separatedReports[key].map(({ value }: { value: string }) => value))
    .map((data) => {
      const color = colorGenerator.next().value;
      return {
        label: 'pussy',
        data,
        backgroundColor: color,
        borderColor: color,
        fill: false,
      };
    });
}

function generateLabels(separatedReports: any): string[] {
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

  calculateElementsByChart(spacingBetweenChartElements = 20) {
    return parseInt(String(window.innerWidth / spacingBetweenChartElements), 10);
  },

  genearateDataSetsForOneSensor(sensorReports: Report[]): ChartData {
    const a = toChartData(separateReportsByType(sensorReports));
    console.log(a);
    return a;
  },

  generateDataSetsForSensors(reports: Report[][]): ChartData {
    const b = toChartData(separateReportsBySensorAndType(reports));
    console.log(b);
    return b;
  },
};
