<template>
  <div class="container">
    <SensorCard v-for="sensor in sensors" :key="sensor.name" :name="sensor.name"></SensorCard>
  </div>
</template>

<script lang="ts">
import { State, useState } from '@/store';
import SensorCard from '@/components/SensorCard.vue';
import { defineComponent } from 'vue';
import { sensorService } from '@/services';
import { Sensor } from '@/services/models';

export default defineComponent({
  name: 'Panel',
  components: {
    SensorCard,
  },
  data(): {
  store?: Readonly<State>;
  sensors: Sensor[];
  } {
    return {
      store: useState(),
      sensors: [],
    };
  },
  mounted() {
    if (!this.store?.token) {
      this.$router.push('/login');
    }
    sensorService.getAll()
      .then((sensors) => {
        this.sensors = sensors;
      });
  },
});
</script>

<style scoped>
.container {
  bottom: 0px;
  width: 100vw;
  padding: 20px;
  margin-top: 40px;

  display: grid;
  grid-template-columns: 49% 49%;
  gap: 20px;
}

@media only screen and (max-width: 900px) {
  .container {
    grid-template-columns: 100%;
  }
}

@media only screen and (min-width: 1417px) {
  .container {
    grid-template-columns: 32% 33% 32%;
    width: 100%;
    max-width: 1600px;
    margin-left: auto;
    margin-right: auto;
  }
}
</style>
