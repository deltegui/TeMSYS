<template>
  <div class="card general-card">
    <div class="average-temp">{{currentAverageTemp}}ºC</div>

    <div>Humedad media: {{currentAverageHumidity}}%;</div>

    <div id="previous-chart-container">
      <LatestTempChart :average="true" />
    </div>
  </div>
</template>

<script>
import {formatSensorData} from './format';
import api from "../../api";
import LatestTempChart from "./LatestTempChart";

function getCurrentSeasonName() {
  const now = new Date();
  const month = now.getMonth();
  if (month >= 2 && month < 5) {
    return "spring";
  }
  if (month >= 5 && month < 8) {
    return "summer";
  }
  if (month >= 8 && month < 11) {
    return "autumn";
  }
  if (month === 11 || (month >= 0 && month < 2)) {
    return "winter";
  }
  return "spring";
}

export default {
  name: "Card",
  data: function() {
    return {
      currentAverageTemp: 0,
      currentAverageHumidity: 0
    };
  },
  components: {
    LatestTempChart
  },
  methods: {
    loadData() {
      api.sensor
        .getCurrentAverageState()
        .then(reports =>
          reports
            .map(({ type, value }) => ({ [type]: value }))
            .reduce((prev, current) => Object.assign(prev, current), {})
        )
        .then(reports => {
          this.currentAverageTemp = formatSensorData(reports.temperature);
          this.currentAverageHumidity = formatSensorData(reports.humidity);
        });
    },

    setSeasonImage() {
      const season = getCurrentSeasonName();
      const element = document.getElementsByClassName("general-card")[0];
      element.style.backgroundImage = `linear-gradient(rgba(65, 179, 163, 0.8) 70%, var(--main_color)), url("/${season}.jpg")`;
    }
  },
  mounted() {
    this.loadData();
    this.setSeasonImage();
  }
};

</script>

<style scoped>
.general-card {
  background: #000;
  background-image: linear-gradient(
      rgba(65, 179, 163, 0.8) 70%,
      var(--main_color)
    ),
    url("/spring.jpg");
  background-position: right top;
  background-size: cover;
  background-repeat: no-repeat;
  height: 100vh;
  width: 100%;
  color: white;
}

.average-temp {
  font-size: 80px;
}

#previous-chart-container {
  margin-top: 50vh;
}
</style>