<template>
  <div class="container">
    <HistoricSettings :show="showSettings" @close="close"/>
    <h1 v-on:click="show">Historic</h1>
    <GeneralChart :filters="filters" />
  </div>
</template>

<script lang="ts">
import GeneralChart from '@/components/GeneralChart.vue';
import HistoricSettings from '@/components/HistoricSettings.vue';
import { ReportFilter } from '@/services/models';
import { defineComponent } from 'vue';

const filters: ReportFilter[] = [
  {
    name: 'habitacion',
    type: 'humidity',
    trim: 30,
  },
  {
    name: 'cocina',
    type: 'humidity',
    trim: 30,
  },
  {
    name: 'salon',
    type: 'humidity',
    trim: 30,
  },
];

export default defineComponent({
  name: 'Historic',
  components: {
    GeneralChart,
    HistoricSettings,
  },
  data() {
    return {
      filters,
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
  },
});
</script>

<style scoped>
.container {
  padding: 30px 20px 20px 20px;
  width: 100vw;
  height: 100vh;
}
</style>
