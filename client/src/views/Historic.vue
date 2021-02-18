<template>
  <div class="container">
    <h1>Humidity</h1>
    <GeneralChart :filters="filters" />
  </div>
</template>

<script lang="ts">
import GeneralChart from '@/components/GeneralChart.vue';
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
  },
  data() {
    return {
      filters,
    };
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
