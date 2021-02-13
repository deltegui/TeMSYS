<template>
  <div id="container">
    <div id="item">
      <img src="@/assets/logo.png" />
      <h1>Temsys</h1>
      <div class="temsys-form">
        <input
          class="temsys-input"
          type="text"
          placeholder="Name"
          name="name"
          v-model="username"
        />
        <input
          class="temsys-input"
          type="password"
          placeholder="Password"
          name="password"
          v-model="password"
        />
        <button class="temsys-input temsys-btn" @click="onLogin">
          Send!
        </button>
        <div v-if="loginErr" class="temsys-input">
          {{ loginErr }}
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { userService } from '@/services';
import { useState } from '@/store';

export default defineComponent({
  name: 'Login',
  data() {
    return {
      username: '',
      password: '',
      loginErr: '',
      store: useState(),
    };
  },
  /*
  mounted() {
    if (this.store?.state.token) {
      this.$router.push('/panel');
    }
  },
  */
  methods: {
    onLogin() {
      this.loginErr = '';
      userService.login({
        name: this.username,
        password: this.password,
      })
        .then((token) => {
          this.$router.push('/panel');
          if (!this.store) return;
          this.store.setToken(token);
        })
        .catch((err) => {
          this.loginErr = err.reason;
        });
    },
  },
});
</script>

<style scoped>
#container {
  display: grid;
  grid-template-columns: auto 500px auto;
  grid-template-rows: 100px 550px auto;
}

#item {
  align-items: center;
  text-align: center;
  grid-column-start: 2;
  grid-row-start: 2;
}

@media only screen and (max-width: 520px) {
  #container {
    grid-template-columns: auto 80% auto;
  }
}
</style>
