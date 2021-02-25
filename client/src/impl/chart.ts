/* eslint-disable no-new */
import Chart from 'chart.js';

function createYAxesFrom(datasets: any): any {
  const yAxes = [
    {
      id: datasets[0].label,
      position: 'left',
      gridLines: {
        display: false,
      },
    },
  ];
  if (datasets.length === 2) {
    yAxes.push({
      id: datasets[1].label,
      position: 'right',
      gridLines: {
        display: false,
      },
    });
  }
  return yAxes;
}

function label(tooltips: any): any {
  return (items: any) => {
    if (tooltips) {
      return `[${items.yLabel}]: ${tooltips[items.datasetIndex ?? 0][items.index ?? 0]}`;
    }
    return `${items.yLabel}: ${items.xLabel}`;
  };
}

export function createChart({
  mountID,
  datasets,
  labels,
  tooltips = undefined,
  showLegend = false,
  showTitle = false,
}: ChartOptions): Chart {
  const yAxes = (datasets.length === 0) ? [] : createYAxesFrom(datasets);
  return new Chart(mountID, {
    type: 'line',
    data: {
      labels,
      datasets,
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      legend: {
        display: showLegend,
      },
      title: {
        display: showTitle,
      },
      tooltips: {
        enabled: true,
        mode: 'single',
        callbacks: {
          label: label(tooltips),
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
        yAxes,
      },
    },
  });
}

export type ChartOptions = {
  mountID: string;
  datasets: any;
  labels: string[];
  tooltips?: string[][];
  showLegend?: boolean;
  showTitle?: boolean;
}

export class SensorChart {
  private chart: Chart;

  constructor(private options: ChartOptions) {
    this.chart = createChart(this.options);
  }

  update(data: {
    datasets: any;
    labels: string[];
    tooltips?: any;
  }) {
    this.chart.data = {
      labels: data.labels,
      datasets: data.datasets,
    };
    this.chart.options = {
      tooltips: {
        callbacks: {
          label: label(data.tooltips),
        },
      },
      scales: {
        yAxes: (data.datasets.length === 0) ? [] : createYAxesFrom(data.datasets),
      },
      legend: {
        display: this.options.showLegend,
      },
      title: {
        display: this.options.showTitle,
      },
    };
    this.chart.update();
  }
}

export function drawChart(options: ChartOptions) {
  new SensorChart(options);
}
