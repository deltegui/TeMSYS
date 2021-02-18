<template>
  <div class="latest-temp-chart">
    <canvas class="previous-chart" :id="`previous-chart-${name}`"></canvas>
  </div>
</template>

<script lang="ts">
/* eslint-disable no-plusplus */

import { Report } from '@/services/models';
import { drawChart } from '@/impl/chart';
import {
  reportService,
  sensorService,
  chartService,
} from '@/services';
import { defineComponent } from 'vue';

function generateCanvasID(name: string) {
  return `previous-chart-${name}`;
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

export default defineComponent({
  name: 'chart',
  props: {
    sensor: String,
    average: Boolean,
  },
  data() {
    return {
      name: '',
    };
  },
  methods: {
    loadAverageChart() {
      sensorService.getAll()
        .then((sensors) => Promise.all(
          sensors.map(({ name }) => reportService.getTemperatureLatestReports(
            name,
            chartService.calculateElementsByChart(),
          )),
        ))
        .then(chartService.flatAndSort.bind(chartService))
        .then(groupReportsByDate)
        .then(calculateAverageByGroup)
        .then(reportService.roundAllReports.bind(reportService))
        .then((reports) => this.draw(reports));
    },

    loadChartFor(sensorName: string) {
      reportService.getTemperatureLatestReports(sensorName, 40)
        .then(chartService.sortReports)
        .then((reports) => this.draw(reports));
    },

    draw(reports: Report[]) {
      const data = chartService.genearateDataSetsForOneSensor(reports);
      drawChart({
        mountID: generateCanvasID(this.name),
        ...data,
      });
    },
  },
  mounted() {
    if (this.average) {
      this.name = 'average';
      this.loadAverageChart();
    } else {
      this.name = this.sensor ?? '';
      this.loadChartFor(this.sensor ?? '');
    }
  },
});
</script>

<style scoped>
.previous-chart {
  height: 150px;
}
</style>
