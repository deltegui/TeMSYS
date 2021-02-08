import Chart from "chart.js";

export class ReportsChart {
  constructor(mountID) {
    this.mountID = mountID;
  }

  set data(reports) {
    const separatedReports = separateReportsByType(reports);
    this.datasets = generateDataSets(separatedReports);
    this.labels = generateLabels(separatedReports);
  }

  draw() {
    drawChart({
      mountID: this.mountID,
      datasets: this.datasets,
      labels: this.labels,
    });
  }
}

function* generateColor() {
  const chartColors = [
    '#b35a41',
    '#0269A4',
  ];
  let currentIndex = 0;
  for(;;) {
    yield chartColors[currentIndex];
    currentIndex = (currentIndex + 1) % chartColors.length;
  }
}

function separateReportsByType(reports) {
  const separated = {};
  for(let report of reports) {
    if(!separated[report.type]) {
      separated[report.type] = [];
    }
    separated[report.type].push(report);
  }
  return separated;
}

function generateDataSets(separatedReports) {
  const colorGenerator = generateColor();
  return Object
    .keys(separatedReports)
    .map(key => {
      return separatedReports[key].map(({value}) => value);
    })
    .map(data => {
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

function generateLabels(separatedReports) {
  const firstGroup = separatedReports[Object.keys(separatedReports)[0]];
  return firstGroup.map(e => e.date.getHours());
}

export function drawChart({
 mountID,
 datasets,
 labels,
 showLegend,
 showTitle,
}) {
  const ctx = document.getElementById(mountID);
  new Chart(ctx, {
    type: "line",
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
    },
  });
}