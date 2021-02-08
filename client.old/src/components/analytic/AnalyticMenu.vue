<template>
  <hsc-menu-style-metal>
    <hsc-menu-bar>
      <hsc-menu-bar-item label="Show data">
        <hsc-menu-item label="Show variable data" @click="() => this.entryClick('ShowData')"/>
      </hsc-menu-bar-item>
    </hsc-menu-bar>
  </hsc-menu-style-metal>
</template>

<script>
export default {
  name: 'analytic-menu',
  props: {
    entryClick: Function,
  },
};
</script>

<style scoped>
hsc-menu-style-metal {
  position: fixed;
  z-index: 2;
}

hsc-menu-bar {
  border-radius: 0 0 4pt 0;
}
</style>