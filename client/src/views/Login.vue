<template>
  <div id="container">
    <div id="item">
      <img src="@/assets/logo.png" />
      <h1>Temsys</h1>
      <div class="temsys-form">
        <input
          id="inputName"
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
/* eslint-disable no-unused-expressions */
import { defineComponent } from 'vue';
import { useState } from '@/store';
import { userService } from '@/services';

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
  mounted() {
    document.getElementById('inputName')?.focus();
    if (this.store?.token) {
      this.$router.push('/panel');
    }
  },
  methods: {
    onLogin() {
      this.loginErr = '';
      userService.login({
        name: this.username,
        password: this.password,
      })
        .then(() => this.$router.push('/panel'))
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
