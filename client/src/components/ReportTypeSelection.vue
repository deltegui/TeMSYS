<template>
  <DropDown title="ReportTypes" :elements="elements" />
</template>

<script lang="ts">
import { reportService } from '@/services';
import { defineComponent, PropType } from 'vue';
import DropDown from './DropDown.vue';

export default defineComponent({
  components: { DropDown },
  props: {
    loadInitialSelection: {
      type: Function as PropType<() => Promise<string[]>>,
      default: () => Promise.resolve([]),
    },
  },
  data(): {
  elements: { checked?: boolean; name: string }[];
  } {
    return {
      elements: [],
    };
  },
  mounted() {
    reportService.getAllReportTypes()
      .then((types) => {
        this.elements = types.map((name) => ({ name }));
      })
      .then(this.loadInitialSelection)
      .then((initial) => {
        this.elements
          .filter((e) => initial.includes(e.name))
          .forEach((e) => {
            e.checked = true;
          });
      });
  },
});
</script>
