<template>
  <div class="card sensor-card">
    <h2 class="sensor-name">{{sensorName}}</h2>
    <div class="bottle">
      <div v-bind:style="dynBottle" class="bottle-fill"></div>
      <div class="bottle-percent">{{currentLevel}}%</div>
    </div>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'Card',
  props: {
    sensorName: String,
  },
  components: {},
  data: function() {
    return {
      currentLevel: 90,
      dynBottle: {
        height: '0%',
        top: '0%',
      }
    };
  },
  methods: {
    setDynBottle() {
      this.dynBottle = {
        height: `${this.currentLevel}%`,
        top: `${this.inverseLevel}%`,
      }
      if(this.currentLevel >= 90) {
        Object.assign(this.dynBottle, {
          'animation-play-state': 'running',
        });
      }
    }
  },
  computed: {
    inverseLevel: function() {
      return 100 - this.currentLevel;
    },
  },
  mounted() {
    api.sensor.getCurrentStateByName(this.sensorName)
      .then(reports => {
        this.currentLevel = reports[0].value;
      })
      .then(this.setDynBottle.bind(this));
  }
};
</script>

<style scoped>
.sensor-name {
  padding-top: 50px;
  margin-right: 150px;
}

.sensor-card {
  display: flex;
  flex-direction: row;
  color: white;
}

.bottle {
  position: relative;
  display: block;
  background-color: rgba(0,0,0,0);
  width: 100px;
  height: 200px;

  border-bottom-style: solid;
  border-bottom-width: 5px;
  border-bottom-color: black;
  border-left-style: solid;
  border-left-width: 5px;
  border-left-color: black;
  border-right-style: solid;
  border-right-width: 5px;
  border-right-color: black;

  padding-top: 5px;
}

.bottle-fill {
  background-color: blue;
  width: 100%;
  height: 0%;
  position: absolute;
  top: 0%;
  animation-name: bottle-danger;
  animation-duration: 3s;
  animation-iteration-count: infinite;
  animation-timing-function: cubic-bezier(1,.05,.46,.97);;
  animation-play-state: paused;
}

.bottle-percent {
  display: inline;
  text-align: center;
  position: absolute;
  top: 40%;
  left: 10px;
  font-size: 30px;
}

@keyframes bottle-danger {
  from {
    background-color: blue;
  }
  50% {
    background-color: red;
  }
  to {
    background-color: blue;
  }
}

@media only screen and (max-width: 470px) {
  .sensor-card {
    flex-direction: column;
  }
}
</style>