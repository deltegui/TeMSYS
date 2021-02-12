<template>
  <GeneralOverview></GeneralOverview>
  <div class="cards-grid">
    <SensorCard name="salon"></SensorCard>
    <SensorCard name="habitacion"></SensorCard>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import GeneralOverview from '@/components/GeneralOverview.vue';
import SensorCard from '@/components/SensorCard.vue';

export default defineComponent({
  name: 'Overview',
  components: {
    GeneralOverview,
    SensorCard,
  },
});
</script>

<style scoped>
.cards-grid {
  position: absolute;
  bottom: 0px;

  width: 100vw;

  display: grid;
  grid-template-columns: auto auto;
  gap: 20px;

  padding: 20px;
}

@media only screen and (max-width: 1100px) {
  .cards-grid {
    grid-template-columns: auto;
  }
}
</style>
