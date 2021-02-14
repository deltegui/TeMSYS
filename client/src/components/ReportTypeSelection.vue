<template>
  <div class="report_type_selection">
    <span>
      <button class="temsys-btn temsys-gray" @click="onOpen">Report types</button>
      <span>Selected: {{selected.join(', ')}}</span>
    </span>
    <div v-if="dropdownOpen" class="temsys-btn dropdown">
      <div v-for="t in types" :key="t">
        <input
          :checked="selected.includes(t) ? true : false"
          type="checkbox"
          @change="() => onCheck(t)"
        />
        {{t}}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { reportService } from '@/services';
import { defineComponent, PropType } from 'vue';

export default defineComponent({
  props: {
    loadInitialSelection: {
      type: Function as PropType<() => Promise<string[]>>,
      default: () => Promise.resolve([]),
    },
  },
  data(): {
  types: string[];
  selected: string[];
  dropdownOpen: boolean;
  } {
    return {
      types: [],
      selected: [],
      dropdownOpen: false,
    };
  },
  mounted() {
    reportService.getAllReportTypes()
      .then((types) => {
        this.types = types;
      })
      .then(this.loadInitialSelection)
      .then((initial) => {
        this.selected = initial;
      });
  },
  methods: {
    onOpen() {
      this.dropdownOpen = !this.dropdownOpen;
      this.$emit('selection', this.selected);
    },

    onCheck(selection: string) {
      if (this.selected.includes(selection)) {
        this.selected = this.selected.filter((element) => element !== selection);
        return;
      }
      this.selected.push(selection);
    },
  },
});
</script>

<style scoped>
.report_type_selection {
  position: relative;
  width: 100%;
}

.report_type_selection > span {
  display: grid;
  gap: 20px;
  grid-template-columns: 150px auto;
}

.report_type_selection > .dropdown {
  background-color: var(--bg-alternative-color);

  position: absolute;
  width: 100%;
  height: 130px;

  padding: 5px;

  overflow: scroll;
}
</style>
