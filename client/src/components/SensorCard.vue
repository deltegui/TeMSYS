<template>
  <div class="card">
    <aside>
      <span v-bind:class="enabled ? 'sensor-enabled' : 'sensor-disabled'"></span>
      <h1>{{ name }}</h1>
    </aside>
    <main>
      <div v-if="enabled">
        <h1>{{ temperature }}ºC</h1>
        <h2 v-if="!!humidity">Humedad: {{ humidity }}%</h2>
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
import ApiSensorRepository from '@/impl/api/sensor.repo';

const sensorRepo = new ApiSensorRepository();

export default defineComponent({
  name: 'SensorCard',
  props: {
    name: {
      type: String,
      required: true,
    },
  },
  data(): {
    temperature: number | null;
    humidity: number | null;
    enabled: boolean;
    } {
    return {
      temperature: null,
      humidity: null,
      enabled: false,
    };
  },
  mounted() {
    sensorRepo.getCurrentStateByName(this.name)
      .then((reports: Report[]) => reports.forEach((report) => {
        this.enabled = true;
        if (report.type === 'temperature') {
          this.temperature = report.value;
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
  border-style: none;
  padding: 10px;
  height: 150px;
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
</style>
