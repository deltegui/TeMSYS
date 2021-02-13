<template>
  <div class="general-card">
    <div class="average-temp">{{ temperature }}ºC</div>
    <div>Humedad: {{ humidity }}%</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Report } from '@/services/models';
import { reportService } from '@/services';

function getCurrentSeasonName(): string {
  const now = new Date();
  const month = now.getMonth();
  if (month >= 2 && month < 5) {
    return 'spring';
  }
  if (month >= 5 && month < 8) {
    return 'summer';
  }
  if (month >= 8 && month < 11) {
    return 'autumn';
  }
  if (month === 11 || (month >= 0 && month < 2)) {
    return 'winter';
  }
  return 'spring';
}

function setSeasonImage(): void {
  const season = getCurrentSeasonName();
  const element = document.getElementsByClassName('general-card')[0];
  element.setAttribute(
    'style',
    `background-image: linear-gradient(rgba(26, 26, 29, 0.8) 70%, var(--bg-main-color)), url("/img/${season}.jpg")`,
  );
}

export default defineComponent({
  name: 'GeneralComponent',
  data() {
    return {
      temperature: 0,
      humidity: 0,
    };
  },
  mounted() {
    setSeasonImage();
    reportService.getLastReadAverage()
      .then((reports: Report[]) => reports.forEach((report) => {
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
.general-card {
  background: var(--bg-main-color);
  background-image: linear-gradient(
      rgba(65, 179, 163, 0.8) 70%,
      var(--bg-main-color)
    ),
    url("/assets/spring.jpg");
  background-position: right top;
  background-size: cover;
  background-repeat: no-repeat;

  height: 100vh;
  width: 100%;
  color: white;

  text-align: center;
  padding-top: 25vh;
}

.average-temp {
  font-size: 100px;
}

#previous-chart-container {
  margin-top: 50vh;
}

@media only screen and (max-height: 830px) {
  .general-card {
    padding-top: 100px;
  }

  .average-temp {
    font-size: 80px;
  }
}
</style>
