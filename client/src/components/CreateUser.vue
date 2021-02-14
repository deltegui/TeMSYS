<template>
  <div id="create_user" class="temsys-form">
    <h2>Create new user</h2>
    <input class="temsys-input" placeholder="Username" type="text" v-model="username" />
    <input class="temsys-input" placeholder="Password" type="password" v-model="password" />
    <p>{{message}}</p>
    <LoadingRoller v-if="loading" />
    <button v-else class="temsys-btn temsys-green" @click="onCreate">Create</button>
  </div>
</template>

<script lang="ts">
import { userService } from '@/services';
import { defineComponent } from 'vue';
import LoadingRoller from './LoadingRoller.vue';

export default defineComponent({
  components: { LoadingRoller },
  name: 'CreateUser',
  data() {
    return {
      username: '',
      password: '',
      loading: false,
      message: '',
    };
  },
  methods: {
    onCreate() {
      this.loading = true;
      userService.create({ name: this.username, password: this.password })
        .then((user) => {
          this.username = '';
          this.password = '';
          return `Created user with name '${user.name}'`;
        })
        .catch((err) => `Error: [code]: ${err.code}, [reason]: ${err.reason}`)
        .then((msg) => {
          this.loading = false;
          this.message = msg;
        });
    },
  },
});
</script>

<style scoped>
#create_user, #create_user > button {
  width: 100%;
}

#create_user > button, #create_user > button > * {
  height: 50px;
}
</style>
