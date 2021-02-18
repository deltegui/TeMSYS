/* eslint-disable no-new */
import Chart from 'chart.js';

export type ChartOptions = {
  mountID: string;
  datasets: any;
  labels: string[];
  tooltips?: string[][];
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
              return `[${items.yLabel}]: ${tooltips[items.datasetIndex ?? 0][items.index ?? 0]}`;
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
