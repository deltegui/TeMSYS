<template>
  <div class="card sensor-card">
    <h2>{{sensorName}}</h2>
    <div class="average-temp">
      {{currentTemp}}ºC
    </div>

    <div v-if="!isNaN(currentHumidity)">
      Humedad: {{currentHumidity}}%
    </div>

    <div>
      <h4>De las últimas dos semanas:</h4>
      <p>Temperatura máxima: {{maxTemp}}ºC</p>
      <p>Temperatura mínima: {{minTemp}}ºC</p>
      <p>Media de temperaturas: {{average}}ºC</p>
      <p>Desviación éstandar de temperatura: {{variation}}ºC</p>
    </div>

    <div id="sensor-chart-container">
      <h4>Temperatura de las últimas horas</h4>
      <LatestTempChart :sensor="sensorName"/>
    </div>
  </div>
</template>

<script>
import api from '../../api';
import LatestTempChart from './LatestTempChart';
import {DataAsset} from '../../analytics';
import {formatSensorData} from './format';

export default {
  name: 'Card',
  props: {
    sensorName: String,
  },
  components: {
    LatestTempChart,
  },
  data: function() {
    return {
      currentTemp: 0,
      currentHumidity: 0,
      maxTemp: 0,
      minTemp: 0,
      average: 0,
      variation: 0,
    };
  },
  methods: {
    loadData() {
      api.sensor.getCurrentStateByName(this.sensorName)
        .then(reports => reports
          .map(({type, value}) => ({[type]: value}))
          .reduce((prev, current) => Object.assign(prev, current), {})
        )
        .then(reports => {
          this.currentTemp = formatSensorData(reports.temperature);
          this.currentHumidity = formatSensorData(reports.humidity);
        });
    },

    calculateAnalytics() {
      const latestDays = 15 * 24;
      api.report.getLatestReports({
        name: this.sensorName,
        trim: latestDays,
        type: 'temperature',
      })
        .then(reports => new DataAsset(reports.map(e => e.value)))
        .then(asset => {
          this.maxTemp = asset.max();
          this.minTemp = asset.min();
          this.average = Math.round(asset.average());
          this.variation = Math.round(asset.standardDeviation() * 100) / 100;
        });
    },
  },
  mounted() {
    this.loadData();
    this.calculateAnalytics();
  }
};
</script>

<style scoped>
.sensor-card {
  height: 630px;
  color: white;
}

.average-temp {
  font-size: 80px;
}

#sensor-chart-container{
  margin-top: 5px;
}
</style>