<template>
  <button id="analytic-btn" class="float-button" v-if="desktop" v-on:click="click">AN</button>  
</template>

<script>
function runningOnDesktop() {
  const isMobile = /iPhone|iPad|iPod|Android/i.test(navigator.userAgent);
	return !isMobile;
}

export default {
  name: 'AccessAnalyticButton',
  data: function() {
    return {
      desktop: runningOnDesktop(),
    };
  },
  props: {
    click: Function,
  },
};
</script>

<style scoped>
.float-button {
    position: fixed;
    padding: 10px;
    border-radius: 50%;
    border-style: solid;
    border-width: 2px;
    color: black;
    background-color: var(--main_color);
    border-color: var(--main-color);
    transition: box-shadow 2s, color 2s background-color 2s;
}

.float-button:hover {
    color: lightgrey;
    background-color: #2C3037;
    border-color: white;
    box-shadow: 2px 2px 10px grey;
}

#analytic-btn {
  top: 30px;
  right: 30px;
}

</style>