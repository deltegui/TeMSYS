<template>
  <div class="dropdown_selection">
    <span>
      <button class="temsys-btn temsys-gray" @click="onOpen">{{title}}</button>
      <span>Selected: {{selected.join(', ')}}</span>
    </span>
    <div v-if="dropdownOpen" class="temsys-btn dropdown">
      <div v-for="t in elements" :key="t">
        <input
          :name="title"
          :checked="selected.includes(t.name) ? true : false || t.checked"
          :type="type"
          @change="() => onCheck(t.name)"
        />
        {{t.name}}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    title: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      default: 'checkbox',
    },
    elements: {
      type: Array as () => { checked: boolean; name: string }[],
      required: true,
    },
  },
  data(): {
  selected: string[];
  dropdownOpen: boolean;
  } {
    return {
      selected: [],
      dropdownOpen: false,
    };
  },
  mounted() {
    this.load();
  },
  watch: {
    elements() {
      this.selected = [];
      this.load();
    },
  },
  methods: {
    load() {
      this.elements
        .filter((e) => e.checked)
        .forEach((e) => this.selected.push(e.name));
    },

    onOpen() {
      this.dropdownOpen = !this.dropdownOpen;
      this.$emit('selection', this.selected);
    },

    onCheck(selection: string) {
      if (this.type === 'checkbox') {
        if (this.selected.includes(selection)) {
          this.selected = this.selected.filter((element) => element !== selection);
          return;
        }
      }
      if (this.type === 'radio') {
        this.selected = [];
      }
      this.selected.push(selection);
    },
  },
});
</script>

<style scoped>
.dropdown_selection {
  position: relative;
  width: 100%;
}

.dropdown_selection > span {
  display: grid;
  gap: 20px;
  grid-template-columns: 150px auto;
}

.dropdown_selection > .dropdown {
  z-index: 2;
  background-color: var(--bg-alternative-color);

  position: absolute;
  width: 100%;
  height: 130px;

  padding: 5px;

  overflow: scroll;
}
</style>