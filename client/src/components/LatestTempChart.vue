<template>
  <div class="latest-temp-chart">
    <canvas class="previous-chart" :id="`previous-chart-${name}`"></canvas>
  </div>
</template>

<script lang="ts">

import { drawChart } from '@/impl/chart';
import {
  reportService,
  sensorService,
  chartService,
} from '@/services';
import { defineComponent } from 'vue';
import { ChartData } from '@/services/chart.service';

function generateCanvasID(name: string) {
  return `previous-chart-${name}`;
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
          sensors.map(({ name }) => reportService.getLatestReports({
            name,
            trim: chartService.calculateElementsByChart(),
          })),
        ))
        .then(chartService.getAverageAllReports.bind(chartService))
        .then((groups) => groups.map(reportService.roundAllReports.bind(reportService)))
        .then((reports) => chartService.generateDataSetsForSensors(reports))
        .then((data) => this.draw(data));
    },

    loadChartFor(sensorName: string) {
      reportService.getLatestReports({ name: sensorName, trim: 48 })
        .then(chartService.sortReports)
        .then((reports) => chartService.genearateDataSetsForOneSensor(reports))
        .then((data) => this.draw(data));
    },

    draw(data: ChartData) {
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
