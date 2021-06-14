<template>
  <DropDown title="ReportTypes" :elements="elements" />
</template>

<script lang="ts">
import { reportService } from '@/services';
import { defineComponent, PropType } from 'vue';
import DropDown, { DropDownElement } from './DropDown.vue';

type Data = {
  elements: DropDownElement[];
};

export default defineComponent({
  components: { DropDown },
  props: {
    loadInitialSelection: {
      type: Function as PropType<() => Promise<string[]>>,
      default: () => Promise.resolve([]),
    },
  },
  data(): Data {
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
        this.elements = this.elements
          .map(({ name }) => {
            if (initial.includes(name)) {
              return {
                name,
                checked: true,
              };
            }
            return {
              name,
              checked: false,
            };
          });
      });
  },
});
</script>
