<template>
  <GeneralOverview></GeneralOverview>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import GeneralOverview from '@/components/GeneralOverview.vue';

export default defineComponent({
  name: 'Overview',
  components: {
    GeneralOverview,
  },
});
</script>
