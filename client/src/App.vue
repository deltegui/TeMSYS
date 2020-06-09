<template>
  <div id="app">
    <SmartScreen :changeMode="onChangeMode" v-if="smartmode"/>
    <AnalyticScreen v-else/>
  </div>
</template>

<script>
import SmartScreen from './components/smart/SmartScreen';
import AnalyticScreen from './components/analytic/AnalyticScreen';

export default {
  name: 'app',
  components: {
    SmartScreen,
    AnalyticScreen,
  },
  data: function() {
    return {
      smartmode: true,
    };
  },
  methods: {
    onChangeMode() {
      this.smartmode = false;
    },
  },
};
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Roboto&display=swap');

:root {
  --main_color: #41B3A3;
  --inverted_color: #b34151;
}

* {
  box-sizing: border-box;
  font-family: 'Roboto', sans-serif;
}

body, html {
  padding: 0px;
  margin: 0px;
}

#app {
  height: 100vh;
  width: 100vw;
}

.card {
  padding: 20px;
  display: flex;
  flex-direction: column;
  background-color: var(--main_color);
}
</style>
