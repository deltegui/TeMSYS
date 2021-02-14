<template>
  <div id="create_user" class="temsys-form">
    <h4>{{name}}</h4>
    <button class="temsys-btn" @click="onDelete">Delete</button>
    <p>{{message}}</p>
  </div>
</template>

<script lang="ts">
import { userService } from '@/services';
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    name: String,
  },
  data() {
    return {
      message: '',
    };
  },
  methods: {
    onDelete() {
      userService.delete(this.name ?? '')
        .then(() => {
          this.message = 'ok';
        })
        .catch(({ reason }) => {
          this.message = reason;
        });
    },
  },
});
</script>

<style scoped>
#create_user {
  width: 100%;
  display: grid;
  grid-template-columns: 140px auto;
  gap: 10px;
}
</style>
