<template>
  <div>
  <canvas class="previous-chart" :id="`previous-chart-${name}`"></canvas>
  </div>
</template>

<script>
import api from '../../api';
import {ReportsChart} from '../../charts';

const spacingBetweenChartElements = 50;
const elementsByChart = parseInt(window.innerWidth / spacingBetweenChartElements);

export default {
  name: 'chart',
  props: {
    'sensor': String,
    'average': Boolean,
  },
  data: function() {
    return {
      name: '',
    };
  },
  methods: {
    loadAverageChart() {
      api.sensor.getAll()
        .then(sensors => Promise.all(
          sensors.map(({name}) => api.report.getLatestReports({
            name,
            trim: elementsByChart,
            type: 'temperature',
          }))
        ))
        .then(res => res.flat())
        .then(groupReportsByDate)
        .then(calculateAverageByGroup)
        .then(reports => reports.sort(compareReportsByDate))
        .then(reports => this.draw(reports));
    },

    loadChartFor(sensorName) {
      api.report.getLatestReports({
        name: sensorName,
        trim: elementsByChart,
        type: 'temperature',
      })
        .then(reports => reports.sort(compareReportsByDate))
        .then(reports => this.draw(reports));
    },

    draw(reports) {
      const chart = new ReportsChart(generateCanvasID(this.name));
      chart.data = reports;
      chart.draw();
    }
  },
  mounted() {
    if(this.average) {
      this.name = 'average';
      this.loadAverageChart();
    } else {
      this.name = this.sensor;
      this.loadChartFor(this.sensor);
    }
  }
}

function generateCanvasID(name) {
  return `previous-chart-${name}`;
}

function groupReportsByDate(reports) {
  const groups = [];
  for(let r of reports) {
    let found = false;
    for(let grp of groups) {
      if(grp[0].date.getHours() === r.date.getHours()) {
        grp.push(r);
        found = true;
      }
    }
    if(!found) {
      groups.push([r]);
    }
  }
  return groups;
}

function calculateAverageByGroup(groups) {
  const averageReports = [];
  for(let grp of groups) {
    const sum = grp.reduce((prev, current) => prev + current.value, 0);
    const value = sum / grp.length;
    const {type, date} = grp[0];
    averageReports.push({
      type,
      date,
      sensor: 'average',
      value,
    });
  }
  return averageReports;
}

function compareReportsByDate(a, b) {
  return a.date - b.date;
}
</script>

<style scoped>
.previous-chart {
  height: 150px;
}
</style>