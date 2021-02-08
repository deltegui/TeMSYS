<template>
  <DataSelector v-if="!show" :close="close" :show="onShow"/>
  <DataViewer v-else :showdata="data" :close="close"/>
</template>

<script>
import DataSelector from './DataSelector';
import DataViewer from './DataViewer';

export default {
  components: {
    DataSelector,
    DataViewer,
  },
  name: 'showdata-handler',
  data: function() {
    return {
      data: {},
      show: false,
    };
  },
  methods: {
    close () {
      this.$destroy();
      this.$el.parentNode.removeChild(this.$el);
    },
    onShow(data) {
      this.data = data;
      this.show = true;
    },
  },
};
</script>