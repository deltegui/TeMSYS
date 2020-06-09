<template>
<hsc-window-style-metal style="position: fixed; z-index: 1">
  <hsc-window
    key="data-selector"
    title="Select Data"
    :closeButton="true"
    :isOpen.sync="isOpen"
    :width="350"
    :height="400">
    <h3>Select Data you would like to see</h3>
    <fieldset>
      Select sensor: 
      <select name="cars" @change="onSelectChange">
        <option disabled selected value> -- select an option -- </option>
        <option v-for="s in sensors" :key="s.name" :value="s.name">{{s.name}}</option>
      </select>
    </fieldset>
    <fieldset>
      <p>Select reports to show</p>
      <div v-for="r in reports" :key="r">
      <input type="checkbox" :value="r" v-model="checkedReports" checked>{{r}}
      </div>
    </fieldset>
    <fieldset>
      <p>Select reports to show</p>
      From <input type="date" v-model="fromDate"/>
      <br/>
      To <input type="date" v-model="toDate"/>
      <br/>
    </fieldset>
    <button @click="onShow">Show</button>
  </hsc-window>
</hsc-window-style-metal>
</template>

<script>
import api from '../../../api';

export default {
  name: 'data-selector',
  props: {
    close: Function,
    show: Function,
  },
  data() {
    return {
      isOpen: true,
      sensors: [],
      selectedSensor: undefined,
      reports: [],
      checkedReports: [],
      fromDate: undefined,
      toDate: undefined,
    }
  },
  watch: {
    isOpen: function() {
      this.close();
    },
  },
  mounted() {
    api.sensor.getAll()
      .then(sensors => {
        this.sensors = sensors;
      });
  },
  methods: {
    onSelectChange(evt) {
      this.selectedSensor = evt.target.value;
      const sensor = this.sensors.filter(e => e.name === this.selectedSensor)[0];
      if(!this.selectedSensor || !sensor) {
        return;
      }
      this.reports = sensor.supportedReports;
    },
    onShow() {
      const from = new Date(this.fromDate);
      const to = new Date(this.toDate);
      if(!this.selectedSensor) {
        return;
      }
      if(this.reports.length === 0) {
        return;
      }
      if(!this.fromDate) {
        return;
      }
      if(!this.toDate) {
        return;
      }
      this.show({
        sensor: this.selectedSensor,
        reports: this.checkedReports,
        date: {
          from,
          to,
        },
      });
    }
  },
};
</script>

<style scoped>
hsc-window {
  width: 1200px;
}
</style>