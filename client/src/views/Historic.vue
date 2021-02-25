<template>
  <div class="container">
    <HistoricSettings
      :show="showSettings"
      @close="close"
      @dataset="onFiltersChange"/>
    <h1>Historic</h1>
    <GeneralChart :filters="filters" />
    <button class="temsys-btn temsys-green" @click="show">
      Settings
    </button>
  </div>
</template>

<script lang="ts">
import GeneralChart from '@/components/GeneralChart.vue';
import HistoricSettings from '@/components/HistoricSettings.vue';
import { ReportFilter } from '@/services/models';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Historic',
  components: {
    GeneralChart,
    HistoricSettings,
  },
  data(): {
  filters: ReportFilter[];
  showSettings: boolean;
  } {
    return {
      filters: [],
      showSettings: false,
    };
  },
  methods: {
    show() {
      this.showSettings = true;
    },

    close() {
      this.showSettings = false;
    },

    onFiltersChange(datasets: {
    id: number;
    from: Date;
    to: Date;
    sensor: string;
    reports: string[];
    }[]) {
      this.filters = datasets
        .map((dataset) => dataset.reports.map((report) => ({
          fromDate: dataset.from,
          toDate: dataset.to,
          name: dataset.sensor,
          type: report,
        })))
        .flat();
    },
  },
});
</script>

<style scoped>
.container {
  padding: 30px 20px 20px 20px;
  width: 100vw;
  height: 100vh;
}

.container > button {
  z-index: 0;
  position: absolute;
  top: 10px;
  right: 10px;
}
</style>
