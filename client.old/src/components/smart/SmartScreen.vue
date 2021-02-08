<template>
  <div id="smart-screen-container">
    <GeneralCard/>
    <SensorCard sensor-name="salon"/>
    <SensorCard sensor-name="habitacion"/>
    <WaterCard sensor-name="BombonaSalon"/>
    <AnalyticButton :click="changeMode"/>
  </div>
</template>

<script>
import GeneralCard from './GeneralCard';
import SensorCard from './SensorCard';
import AnalyticButton from './AccessAnalyticButton';
import WaterCard from './WaterCard';

export default {
  name: 'SmartScreen',
  components: {
    GeneralCard,
    SensorCard,
    AnalyticButton,
    WaterCard,
  },
  props: {
    changeMode: Function,
  },
};
</script>

<style scoped>
#smart-screen-container {
  display: flex;
  flex-direction: column;

  width: 100%;
}
</style>