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

enum Month {
  JANUARY,
  FEBRUARY,
  MARCH,
  APIRL,
  MAY,
  JUNE,
  JULY,
  AUGUST,
  SEPTEMBER,
  OCTOBER,
  NOBEMBER,
  DICEMBER,
}

function isWinter(month: number) {
  return month === Month.DICEMBER || (month >= Month.JANUARY && month < Month.MARCH);
}

const isSummer = (month: number) => month >= Month.JUNE && month < Month.SEPTEMBER;
const isSpring = (month: number) => month >= Month.MARCH && month < Month.JUNE;
const isAutumn = (month: number) => month >= Month.SEPTEMBER && month < Month.DICEMBER;

function getCurrentSeasonName(): string {
  const now = new Date();
  const month = now.getMonth();
  if (isSpring(month)) {
    return 'spring';
  }
  if (isSummer(month)) {
    return 'summer';
  }
  if (isAutumn(month)) {
    return 'autumn';
  }
  if (isWinter(month)) {
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
