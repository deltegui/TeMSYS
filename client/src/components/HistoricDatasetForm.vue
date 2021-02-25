<template>
  <div id="add-form">
    <h3>New dataset</h3>
    <DropDown
      @selection="onSensorSelect"
      title="Sensor"
      type="radio"
      :elements="availableSensors" />
    <DropDown
      @selection="onReportSelect"
      title="Reports"
      :elements="availableReports" />
    <DateTimeInput title="From" v-model="from" />
    <DateTimeInput title="To" v-model="to" />
    <div v-if="error" class="temsys-input error">
    {{error}}
    </div>
    <button @click="addDataset" class="temsys-btn temsys-green"> Add </button>
  </div>
</template>

<script lang="ts">
import { sensorService } from '@/services';
import { Sensor } from '@/services/models';
import { defineComponent } from 'vue';
import DateTimeInput from './DateTimeInput.vue';
import DropDown from './DropDown.vue';

export default defineComponent({
  components: { DropDown, DateTimeInput },
  data(): {
  error: string;
  from: Date;
  to: Date;
  sensor: string;
  reports: string[];
  sensorStore: Sensor[];
  availableSensors: { name: string; checked?: boolean }[];
  availableReports: { name: string; checked?: boolean }[];
  } {
    return {
      error: '',
      from: new Date(),
      to: new Date(),
      sensor: '',
      reports: [],
      sensorStore: [],
      availableSensors: [],
      availableReports: [],
    };
  },
  mounted() {
    sensorService.getAll()
      .then((sensors) => {
        this.sensorStore = sensors;
        this.availableSensors = sensors.map((s) => ({ name: s.name }));
      });
  },
  watch: {
    sensor() {
      this.availableReports = this.sensorStore
        .find((s) => s.name === this.sensor)
        ?.supportedReports
        .map((name) => ({ name })) ?? [];
    },
  },
  methods: {
    addDataset() {
      if (!this.sensor) {
        this.error = 'Empty sensor!';
        return;
      }
      if (this.reports.length === 0) {
        this.error = 'No reports selected!';
        return;
      }
      this.$emit('add', {
        from: this.from,
        to: this.to,
        sensor: this.sensor,
        reports: this.reports,
      });
      this.error = '';
    },

    onReportSelect(selected: string[]) {
      this.reports = selected;
    },

    onSensorSelect(selected: string[]) {
      [this.sensor] = selected;
    },
  },
});
</script>

<style scoped>
#add-form > * {
  margin-top: 10px;
}

.error {
  border-color: var(--fg-weak-color);
  background-color: var(--fg-weak-color);
}
</style>
