<template>
  <div class="card">
    <aside>
      <span v-bind:class="enabled ? 'sensor-enabled' : 'sensor-disabled'"></span>
      <h1>{{ name }}</h1>
    </aside>
    <main v-if="loading">
      <LoadingRoller />
    </main>
    <main v-else>
      <div v-if="enabled">
        <h1 v-if="!!temperature">{{ temperature }}ºC</h1>
        <h1 v-if="!!watts">{{ watts }} W</h1>
        <h2 v-if="!!humidity">Humedad: {{ humidity }}%</h2>
        <LatestTempChart :sensor="name" />
      </div>
      <div v-else>
        <h2>Sensor is not responding</h2>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Report } from '@/services/models';
import { sensorService } from '@/services';
import { State, useState } from '@/store';
import LoadingRoller from './LoadingRoller.vue';
import LatestTempChart from './LatestTempChart.vue';

export default defineComponent({
  name: 'SensorCard',
  components: {
    LoadingRoller,
    LatestTempChart,
  },
  props: {
    name: {
      type: String,
      required: true,
    },
  },
  data(): {
    temperature: number | null;
    humidity: number | null;
    watts: number | null;
    enabled: boolean;
    loading: boolean;
    store?: State;
    } {
    return {
      temperature: null,
      humidity: null,
      watts: null,
      enabled: false,
      loading: true,
      store: useState(),
    };
  },
  mounted() {
    const timeout = setTimeout(() => {
      this.loading = false;
    }, 4000);
    sensorService.readState(this.name)
      .then((reports: Report[]) => reports.forEach((report) => {
        clearTimeout(timeout);
        this.loading = false;
        this.enabled = true;
        if (report.type === 'temperature') {
          this.temperature = report.value;
        }
        if (report.type === 'watts') {
          this.watts = report.value;
        }
        if (report.type === 'humidity') {
          this.humidity = report.value;
        }
      }));
  },
});
</script>

<style scoped>
.card {
  border-style: solid;
  border-width: 2px;
  border-color: var(--fg-weak-color);
  padding: 10px;
  height: 300px;
}

.card > aside {
  width: 100%;
}

.card > aside span {
  display: inline-block;
  height: 10px;
  width: 10px;
  border-radius: 50%;
}

.card > aside h1 {
  display: inline-block;
  width: 100px;
  margin: 0px 0px 0px 5px;
  font-size: 24px;
}

.sensor-enabled {
  background-color: var(--fg-green-color);
}

.sensor-disabled {
  background-color: red;
}

main {
  text-align: center;
}

main h1 {
  margin: 0px 0px 5px 0px;
  font-size: 40px;
}

main h2 {
  margin: 0px;
  font-size: 15px;
}

.latest-temp-chart, .latest-temp-chart > * {
  width: 100%;
  height: 150px;
}
</style>
