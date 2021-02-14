<template>
  <GeneralOverview></GeneralOverview>
  <LatestTempChart :average="true" />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import GeneralOverview from '@/components/GeneralOverview.vue';
import LatestTempChart from '@/components/LatestTempChart.vue';

export default defineComponent({
  name: 'Overview',
  components: {
    GeneralOverview,
    LatestTempChart,
  },
});
</script>

<style scoped>
.latest-temp-chart {
  width: 100vw;
  height: 260px;

  position: absolute;
  bottom: 0px;
  left: 0px;
}
</style>
