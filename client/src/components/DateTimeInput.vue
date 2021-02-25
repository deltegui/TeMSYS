<template>
  <div class="datetime-input">
    <label for="from">{{title}}:</label>
    <span>
      <input
        @change="onChanged"
        v-model="date"
        class="temsys-input"
        type="date" />
      <input
        @change="onChanged"
        v-model="hours"
        class="temsys-input"
        type="number"
        max="23"
        min="0" />
      <div>:</div>
      <input
        @change="onChanged"
        v-model="minutes"
        class="temsys-input"
        type="number"
        max="59"
        min="0" />
    </span>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

function createDate(formattedDate: string, hours: number, minutes: number): Date {
  const date = new Date(formattedDate);
  date.setHours(hours);
  date.setMinutes(minutes);
  return date;
}

function formatDate(date: Date): string {
  return `${date.getFullYear()}-${date.getMonth()}-${date.getDay()}`;
}

export default defineComponent({
  props: {
    title: {
      required: true,
      type: String,
    },
    modelValue: {
      required: true,
      type: Object as () => Date,
    },
  },
  data(): {
  date: string;
  hours: number;
  minutes: number;
  } {
    return {
      date: formatDate(this.modelValue ?? new Date()),
      hours: this.modelValue?.getHours() ?? 21,
      minutes: this.modelValue?.getMinutes() ?? 5,
    };
  },
  methods: {
    onChanged() {
      const date = createDate(this.date, this.hours, this.minutes);
      this.$emit('update:modelValue', date);
    },
  },
});
</script>

<style scoped>
.datetime-input > span {
  display: grid;
  grid-template-columns: auto 100px 2px 100px;
  gap: 10px;
}
</style>
