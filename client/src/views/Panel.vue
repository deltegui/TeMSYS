<template>
  <div class="container">
    <SensorCard name="salon"></SensorCard>
    <SensorCard name="habitacion"></SensorCard>
  </div>
</template>

<script lang="ts">
import { useState } from '@/store';
import SensorCard from '@/components/SensorCard.vue';
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Panel',
  components: {
    SensorCard,
  },
  data() {
    return {
      store: useState(),
    };
  },
  /*
  mounted() {
    if (this.store?.state.token) {
      this.$router.push('/login');
    }
  },
  */
});
</script>

<style scoped>
.container {
  bottom: 0px;
  width: 100vw;
  display: grid;
  grid-template-columns: auto auto;
  gap: 20px;
  padding: 20px;
}

@media only screen and (max-width: 1100px) {
  .container {
    grid-template-columns: auto;
  }
}
</style>
