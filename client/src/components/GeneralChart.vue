<template>
  <div class="general-chart">
    <canvas id="general-chart"></canvas>
  </div>
</template>

<script lang="ts">
import { Report, ReportFilter } from '@/services/models';
import { SensorChart } from '@/impl/chart';
import {
  chartService,
  reportService,
} from '@/services';
import { defineComponent } from 'vue';

const canvasID = 'general-chart';

export default defineComponent({
  name: 'GeneralChart',
  props: {
    filters: {
      type: Array as () => ReportFilter[],
      required: true,
    },
  },
  data(): {
  chart: SensorChart | undefined;
  } {
    return {
      chart: undefined,
    };
  },
  mounted() {
    this.triggerUpload();
  },
  watch: {
    filters() {
      this.triggerUpload();
    },
  },
  methods: {
    triggerUpload() {
      const requests = this.filters?.map(
        (filter) => reportService.readReportsFiltered(filter)
      ) ?? [];
      Promise.all(requests)
        .then(chartService.sortSubets.bind(chartService))
        .then(this.draw.bind(this));
    },

    draw(reports: Report[][]) {
      const data = chartService.generateDataSetsForSensors(reports);
      if (!this.chart) {
        this.chart = new SensorChart({
          mountID: canvasID,
          ...data,
        });
      } else {
        this.chart.update(data);
      }
    },
  },
});
</script>

<style scoped>
.general-chart {
  height: 90%;
  width: 100%;
}
</style>
