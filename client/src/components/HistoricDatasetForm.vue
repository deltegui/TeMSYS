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
    <label for="from">From:</label>
    <input
      v-model="from"
      name="from"
      class="temsys-input"
      type="datetime-local" />
    <label for="to">To:</label>
    <input
      v-model="to"
      name="to"
      class="temsys-input"
      type="datetime-local" />
    <button @click="addDataset" class="temsys-btn temsys-green"> Add </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import DropDown from './DropDown.vue';

export default defineComponent({
  components: { DropDown },
  data(): {
  from: Date;
  to: Date;
  sensor: string;
  reports: string[];
  availableSensors: { name: string; checked?: boolean }[];
  availableReports: { name: string; checked?: boolean }[];
  } {
    return {
      from: new Date(),
      to: new Date(),
      sensor: '',
      reports: [],
      availableSensors: [
        { name: 'salon' },
        { name: 'habitacion' },
        { name: 'cocina' },
      ],
      availableReports: [
        { name: 'humidity' },
        { name: 'temperature', checked: true },
      ],
    };
  },
  methods: {
    addDataset() {
      this.$emit('add', {
        from: this.from,
        to: this.to,
        sensor: this.sensor,
        reports: this.reports,
      });
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
</style>
