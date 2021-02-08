<template>
  <div ref="container">
    <AnalyticMenu :entryClick="handleEntryClick" />
  </div>
</template>

<script>
import AnalyticMenu from './AnalyticMenu';
import windows from '../windows';
import Vue from 'vue';

export default {
  name: 'analytic-screen',
  components: {
    AnalyticMenu,
    ...windows
  },
  methods: {
    handleEntryClick(handlerName) {
      if(!windows[handlerName]) {
        return;
      }
      const HandlerComponent = Vue.extend(windows[handlerName]);
      const handlerInstance = new HandlerComponent({});
      handlerInstance.$mount();
      this.$refs.container.appendChild(handlerInstance.$el);
    },
  },
};
</script>

<style scoped>
</style>