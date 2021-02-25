<template>
  <div id="historic-settings" v-bind:class="classes">
    <button @click="close" class="temsys-btn temsys-red close-btn">
      <img src="@/assets/close-icon.png"/>
    </button>
    <h1>Settings</h1>
    <HistoricDatasetForm @add="addDataset" />
    <div>
      <h3>Selected datasets</h3>
      <div class="dataset" v-for="(set, index) in datasets" v-bind:key="index">
        <h4>{{set.sensor}}</h4>
        {{set.reports.join(", ")}}
        {{set.to}}
        {{set.from}}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import HistoricDatasetForm from './HistoricDatasetForm.vue';

export default defineComponent({
  components: { HistoricDatasetForm },
  props: {
    show: {
      type: Boolean,
      required: true,
    },
  },
  data(): {
  datasets: {
    from: Date;
    to: Date;
    sensor: string;
    reports: string[];
  }[];
  } {
    return {
      datasets: [],
    };
  },
  computed: {
    classes(): string {
      return (this.show) ? 'displayed' : 'hide';
    },
  },
  methods: {
    close() {
      this.$emit('close');
    },

    addDataset(dataset: {
    from: Date;
    to: Date;
    sensor: string;
    reports: string[];
    }) {
      this.datasets.push(dataset);
    },
  },
});
</script>

<style scoped>
#historic-settings {
  position: absolute;

  width: 80vw;
  height: 80vh;
  padding: 20px;

  top: 0px;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;

  opacity: 1;

  overflow: scroll;

  background-color: var(--bg-menu-color);
  box-shadow: 0px 10px 30px #0a0a0d;

  animation: appear 1s ease-in-out;
}

.displayed {
  display: block;
}

.hide {
  display: none;
}

@keyframes appear {
  from {
    top: 30px;
    opacity: 0;
  }

  to {
    top: 0px;
    opacity: 1;
  }
}

.close-btn {
  position: absolute;
  right: 10px;
  top: 10px;

  width: 50px;
  height: 50px;

  background-color: var(--bg-alternative-color);
}

.close-btn:hover {
  background-color: var(--fg-weak-color);
}

.close-btn > img {
  position: absolute;
  top: 6px;
  left: 6px;
  width: 35px;
  height: 35px;
}

@media only screen and (max-width: 900px) {
  #historic-settings {
    width: 100vw;
    height: 100vh;
  }
}

#add-form > * {
  margin-top: 10px;
}

.dataset {
  border-style: solid;
  border-width: 2px;
  border-radius: 5px;
  border-color: var(--fg-weak-color);

  padding: 10px;
  margin-top: 10px;
}

.dataset > h4 {
  padding: 0px 0px 10px 0px;
  margin: 0px;
}
</style>
