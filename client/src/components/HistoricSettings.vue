<template>
  <div id="historic-settings" v-bind:class="classes">
    <button @click="close" class="temsys-btn temsys-red close-btn">
      <img src="@/assets/close-icon.png"/>
    </button>
    <h1>Settings</h1>
    <HistoricDatasetForm @add="addDataset" />
    <div>
      <h3>Selected datasets</h3>
      <div class="dataset" v-for="(set) in datasets" v-bind:key="set.id">
        <h4>{{set.id}} {{set.sensor}}</h4>
        {{set.reports.join(", ")}}
        {{set.to}}
        {{set.from}}
        <button class="temsys-btn temsys-red" @click="() => onDatasetClose(set.id)">
          <img src="@/assets/close-icon.png"/>
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import HistoricDatasetForm from './HistoricDatasetForm.vue';

class ID {
  private ids: Generator<number>;

  constructor() {
    this.ids = ID.generator();
  }

  private static* generator(): Generator<number> {
    let i = 0;
    for (;;) {
      yield i;
      i += 1;
    }
  }

  public next(): number {
    return this.ids.next().value;
  }
}

const id = new ID();

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
    id: number;
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
      this.datasets.push({
        id: id.next(),
        ...dataset,
      });
      this.onChange();
    },

    onDatasetClose(selectedId: number) {
      this.datasets = this.datasets
        .filter((data) => data.id !== selectedId);
      this.onChange();
    },

    onChange() {
      this.$emit('dataset', this.datasets);
    },
  },
});
</script>

<style scoped>
#historic-settings {
  z-index: 1;
  position: absolute;

  width: 80vw;
  height: 80vh;
  padding: 20px;

  border-radius: 5px;

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
  position: relative;

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

.dataset > button {
  position: absolute;
  top: 5px;
  right: 5px;
  width: 40px;
  height: 40px;
}

.dataset > button > img {
  position: absolute;
  top: 8px;
  left: 10px;
  width: 20px;
  height: 20px;
}
</style>
