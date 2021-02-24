<template>
  <div class="container">
    <h1>Sensor Admin</h1>
    <CreateSensor />
    <h2>Sensor list</h2>
    <ModifySensor v-for="s in sensors" :key="s.name" :name="s.name" />
  </div>
</template>

<script lang="ts">
import CreateSensor from '@/components/CreateSensor.vue';
import ModifySensor from '@/components/ModifySensor.vue';
import { sensorService } from '@/services';
import { Sensor } from '@/services/models';
import { defineComponent } from 'vue';

export default defineComponent({
  components: {
    CreateSensor,
    ModifySensor,
  },
  data(): {
  sensors: Sensor[];
  } {
    return {
      sensors: [],
    };
  },
  mounted() {
    sensorService.getAllWithDeleted().then((sensors) => {
      this.sensors = sensors;
    });
  },
});
</script>

<style scoped>
.container {
  bottom: 0px;
  width: 100vw;
  padding: 20px 20vw 20px 20vw;
  margin-top: 40px;

  display: flex;
  flex-direction: column;
}

@media only screen and (max-width: 900px) {
  .container {
    padding: 20px;
  }
}

.container > h1 {
  width: 100%;
}
</style>
