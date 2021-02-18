<template>
  <div class="general-chart">
    <canvas id="general-chart"></canvas>
  </div>
</template>

<script lang="ts">
import { Report, ReportFilter } from '@/services/models';
import { drawChart } from '@/impl/chart';
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
  mounted() {
    this.triggerUpload();
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
      drawChart({
        mountID: canvasID,
        ...data,
      });
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
