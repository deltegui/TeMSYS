<template>
  <input type="checkbox" id="menu" />
  <label id="labelMenu" for="menu" onclick></label>
  <div id="container">
    <nav id="menubar">
      <img src="@/assets/logo.png" />
      <div v-if="!!store.token">
        <h2>{{store.token.owner}}</h2>
        <p>{{store.token.role}}</p>
      </div>
      <router-link to="/">Overview</router-link>
      <router-link to="/api/doc">Api Documentation</router-link>
      <span v-if="!!store.token">
        <div v-if="store.token.role === 'admin'">
          <router-link to="/useradmin">User Admin</router-link>
          <router-link to="/sensoradmin">Sensor Admin</router-link>
          <router-link to="/typesadmin">ReportTypes Admin</router-link>
        </div>
      </span>
      <router-link v-if="!!store.token" to="/historic">Historical</router-link>
      <router-link v-if="!!store.token" to="/panel">Panel</router-link>
      <a v-if="!!store.token" v-on:click="onLogout">Logout</a>
      <router-link v-else to="/login">Login</router-link>
    </nav>
    <main>
      <router-view />
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useState } from '@/store';
import { userService } from './services';

export default defineComponent({
  data() {
    return {
      store: useState(),
    };
  },

  mounted() {
    window.addEventListener('click', this.onMenuClick.bind(this));
  },

  unmounted() {
    window.removeEventListener('click', this.onMenuClick.bind(this));
  },

  methods: {
    onMenuClick(evt: MouseEvent) {
      const nav = document.getElementById('menubar');
      const menuButton = document.getElementById('menu') as HTMLInputElement;
      const label = document.getElementById('labelMenu');
      if (evt.target === label || evt.target === menuButton || evt.target === nav) {
        return;
      }
      menuButton.checked = false;
    },

    onLogout() {
      userService.logout();
      this.$router.push('/');
    },

    role() {
      return this.store?.token?.role;
    },
  },
});
</script>

<style scoped>
#container {
  overflow: hidden;
  display: grid;
  grid-template-columns: 100vw;
  grid-template-rows: auto;
}

input[type="checkbox"] {
  position: absolute;
  top: -9999px;
  left: -9999px;
}

input:checked ~ #container nav {
  left: 0px;
  box-shadow: 10px 0px 20px #0a0a0d;
}

label {
  z-index: 2;
  position: fixed;
  left: 10px;
  display: block;
  width: 3em;
  height: 3.25em;
  padding: 0.35em;
  font-size: 1.1em;
  color: white;
  transition: color 0.3s ease-in-out;
  cursor: pointer;
  user-select: none;
  margin: 0;
}

label:after {
  position: absolute;
  right: 0.25em;
  top: 0;
  content: "\2261";
  font-size: 2.8em;
}

label:hover,
input:checked ~ label {
  color: var(--fg-semi-color);
}

nav {
  z-index: 1;
  display: block;
  text-align: center;
  position: fixed;
  left: -250px;
  width: 250px;
  height: 100vh;
  background-color: var(--bg-menu-color);
  box-shadow: 0px 0px 0px #0a0a0d;
  transition: left 0.3s, box-shadow 0.3s;
}

nav > img {
  width: 100px;
  padding: 20px;
}

nav > a, div > a {
  margin-bottom: 10px;
  width: 100%;
  display: block;
  --main-color: var(--fg-semi-color);
  --text-color: white;

  border-bottom-style: solid;
  border-bottom-width: 1px;
  border-bottom-color: var(--bg-alternative-color);
  border-radius: 3px;

  padding: 10px 20px 10px 20px;

  color: var(--text-color);
  text-decoration: none;
  background-color: rgba(0, 0, 0, 0);

  transition: background-color 0.2s;
}

nav > a:hover, div > a:hover {
  background-color: var(--main-color);
}
</style>

<style>
:root {
  --bg-menu-color: #18181c;
  /*--bg-main-color: #1a1a1d;*/ /* gray color. It can cause contrast problems */
  --bg-main-color: #10101d; /* kind of blue color. Better contrast */
  --bg-alternative-color: #202020;
  --fg-strong-color: #c3073f;
  --fg-semi-color: #a31c52;
  --fg-weak-color: #8f1a49;
  --fg-green-color: #14a690;
}

@font-face {
  font-family: assistant;
  src: url("assets/Assistant-Medium.ttf");
}

@font-face {
  font-family: opensans;
  src: url("assets/OpenSans-Regular.ttf");
}

* {
  box-sizing: border-box;
  color: white;
  font-family: "opensans";
}

body {
  padding: 0px;
  margin: 0px;
  background-color: var(--bg-main-color);
}

h1 {
  font-family: "assistant";
  color: var(--fg-semi-color);
}

.temsys-input {
  --main-color: var(--fg-weak-color);
  --text-color: white;

  width: 100%;
  display: block;

  border-style: solid;
  border-width: 2px;
  border-color: #29292e;
  border-radius: 3px;

  padding: 10px;

  background-color: #29292e;
  color: white;

  transition: border-color 0.2s;
}

.temsys-input:hover,
.temsys-input:focus-within {
  border-color: var(--main-color);
}

.temsys-btn {
  --main-color: var(--fg-semi-color);
  --text-color: white;

  padding: 10px 20px 10px 20px;

  border-style: solid;
  border-width: 2px;
  border-color: var(--main-color);
  border-radius: 3px;

  color: var(--text-color);
  background-color: rgba(0, 0, 0, 0);

  box-shadow: 5px 5px black;

  transition: background-color 0.2s;
}

.temsys-btn:hover {
  background-color: var(--main-color);
}

.temsys-btn:active {
  box-shadow: 0px 0px black;
}

.temsys-form * {
  margin-top: 10px;
}

.temsys-green {
  --main-color: var(--fg-green-color);
}

.temsys-gray {
  --main-color: var(--bg-alternative-color);
}
</style>
