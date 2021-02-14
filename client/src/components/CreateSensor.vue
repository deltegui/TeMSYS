<template>
  <div id="create_sensor" class="temsys-form">
    <h2>Register new sensor</h2>
    <input
      class="temsys-input"
      placeholder="Sensor name"
      type="text"
      v-model="name"
    />
    <ReportTypeSelection v-on:selection="onReportTypeSelection" />
    <input class="temsys-input" placeholder="IP" type="text" v-model="ip" />
    <input
      class="temsys-input"
      placeholder="Update Interval"
      type="number"
      v-model="updateInterval"
    />
    <p>{{ message }}</p>
    <LoadingRoller v-if="loading" />
    <button v-else class="temsys-btn temsys-green" @click="onCreate">
      Create
    </button>
  </div>
</template>

<script lang="ts">
import { sensorService } from '@/services';
import { defineComponent } from 'vue';
import ReportTypeSelection from './ReportTypeSelection.vue';
import LoadingRoller from './LoadingRoller.vue';

export default defineComponent({
  name: 'CreateSensor',
  components: {
    LoadingRoller,
    ReportTypeSelection,
  },
  data(): {
  name: string;
  ip: string;
  updateInterval: string;
  reportTypes: string[];
  loading: boolean;
  message: string;
  } {
    return {
      name: '',
      ip: '',
      updateInterval: '60',
      reportTypes: [],
      loading: false,
      message: '',
    };
  },
  methods: {
    onCreate() {
      this.loading = true;
      sensorService.create({
        name: this.name,
        connection: {
          type: 'http',
          value: this.ip,
        },
        updateInterval: parseInt(this.updateInterval, 10),
        deleted: false,
        supportedReports: this.reportTypes,
      })
        .then(() => {
          this.message = 'Ok!';
        })
        .catch((err) => {
          this.message = err.reason;
        })
        .then(() => {
          this.loading = false;
          this.name = '';
          this.ip = '';
          this.updateInterval = '60';
        });
    },

    onReportTypeSelection(selection: string[]) {
      this.reportTypes = selection;
    },
  },
});
</script>

<style scoped>
#create_user,
#create_user > button {
  width: 100%;
}

#create_user > button,
#create_user > button > * {
  height: 50px;
}
</style>
