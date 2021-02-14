<template>
  <div id="create_sensor" class="temsys-form">
    <h4>{{name}}</h4>
    <ReportTypeSelection
      :loadInitialSelection="loadReportTypes"
      v-on:selection="onReportTypeSelection"
    />
    <input class="temsys-input" placeholder="IP" type="text" v-model="ip" />
    <input
      class="temsys-input"
      placeholder="Update Interval"
      type="number"
      v-model="updateInterval"
    />
    <input type="checkbox" :checked="deleted ? true : false" /> Deleted?
    <p>{{ message }}</p>
    <LoadingRoller v-if="loading" />
    <span class="buttons" v-else>
      <button class="temsys-btn temsys-green" @click="onUpdate">
        Update
      </button>
      <button class="temsys-btn" @click="onDeleted">
        Delete
      </button>
    </span>
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
  props: {
    name: {
      type: String,
      required: true,
    },
  },
  data(): {
  ip: string;
  updateInterval: string;
  reportTypes: string[];
  deleted: boolean;
  loading: boolean;
  message: string;
  } {
    return {
      ip: '',
      updateInterval: '60',
      reportTypes: [],
      deleted: false,
      loading: false,
      message: '',
    };
  },
  methods: {
    onUpdate() {
      this.loading = true;
      sensorService.update({
        name: this.name,
        connection: {
          type: 'http',
          value: this.ip,
        },
        updateInterval: parseInt(this.updateInterval, 10),
        deleted: this.deleted,
        supportedReports: this.reportTypes,
      })
        .then(() => {
          this.message = 'Done updating!';
        })
        .catch((err) => {
          this.message = err.reason;
        })
        .then(() => {
          this.loading = false;
        });
    },

    onReportTypeSelection(selection: string[]) {
      this.reportTypes = selection;
    },

    onDeleted() {
      this.deleted = true;
      this.onUpdate();
    },

    async loadReportTypes(): Promise<string[]> {
      return sensorService.getOne(this.name)
        .then((sensor) => {
          this.ip = sensor.connection.value;
          this.updateInterval = String(sensor.updateInterval);
          this.reportTypes = sensor.supportedReports;
          this.deleted = sensor.deleted;
          return this.reportTypes;
        });
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

.buttons {
  display: grid;
  grid-template-columns: auto auto;
  gap: 10px;
}
</style>
