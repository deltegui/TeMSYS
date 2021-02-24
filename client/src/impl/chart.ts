/* eslint-disable no-new */
import Chart from 'chart.js';

type LabeledObject = {
  label: string;
}

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
        yAxes,
      },
    },
  });
}
