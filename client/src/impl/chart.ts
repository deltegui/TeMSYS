/* eslint-disable no-new */
/* eslint-disable no-param-reassign */
/* eslint-disable no-unused-expressions */
import { Report } from '@/services/models';
import Chart from 'chart.js';

function separateReportsByType(reports: Report[]): any {
  return reports.reduce((prev: any, report: Report) => {
    if (!prev[report.type]) {
      prev[report.type] = [];
    }
    prev[report.type].push(report);
    return prev;
  }, {});
}

function* generateColor() {
  const chartColors = [
    '#b35a41',
    '#0269A4',
  ];
  let currentIndex = 0;
  for (;;) {
    yield chartColors[currentIndex];
    currentIndex = (currentIndex + 1) % chartColors.length;
  }
}

type DataEntry = {
  label: string;
  data: number;
  backgroundColor: string | void;
  borderColor: string | void;
  fill: boolean;
}

function generateDataSets(separatedReports: any): DataEntry[] {
  const colorGenerator = generateColor();
  return Object
    .keys(separatedReports)
    .map((key) => separatedReports[key].map(({ value }: { value: string }) => value))
    .map((data) => {
      const color = colorGenerator.next().value;
      return {
        label: '',
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

function generateTooltips(separatedReports: any): string[] {
  const firstGroup = separatedReports[Object.keys(separatedReports)[0]];
  return firstGroup.map((e: Report) => `${e.date.toDateString()} ${e.date.getHours()}:${e.date.getMinutes()}:${e.date.getSeconds()}`);
}

export type ChartOptions = {
  mountID: string;
  datasets: any;
  labels: string[];
  tooltips?: string[];
  showLegend?: boolean;
  showTitle?: boolean;
}

export function drawChart({
  mountID,
  datasets,
  labels,
  tooltips = undefined,
  showLegend = false,
  showTitle = false,
}: ChartOptions) {
  new Chart(mountID, {
    type: 'line',
    data: {
      labels,
      datasets,
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      legend: {
        display: !!showLegend,
      },
      title: {
        display: !!showTitle,
      },
      tooltips: {
        enabled: true,
        mode: 'single',
        callbacks: {
          label(items) {
            if (tooltips) {
              return `[${items.yLabel} ºC]: ${tooltips[items.index ?? 0]}`;
            }
            return `${items.yLabel}: ${items.xLabel}`;
          },
        },
      },
      scales: {
        xAxes: [
          {
            gridLines: {
              display: false,
            },
          },
        ],
        yAxes: [
          {
            gridLines: {
              display: false,
            },
          },
        ],
      },
    },
  });
}

export class ReportsChart {
  private datasets: any;

  private labels: string[];

  private mountID: string;

  private tooltips: string[];

  constructor(mountID: string) {
    this.mountID = mountID;
    this.labels = [];
    this.datasets = {};
    this.tooltips = [];
  }

  set data(reports: Report[]) {
    const separatedReports = separateReportsByType(reports);
    this.datasets = generateDataSets(separatedReports);
    this.labels = generateLabels(separatedReports);
    this.tooltips = generateTooltips(separatedReports);
  }

  draw() {
    drawChart({
      mountID: this.mountID,
      datasets: this.datasets,
      labels: this.labels,
      tooltips: this.tooltips,
    });
  }
}
