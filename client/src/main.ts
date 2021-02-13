import { createApp } from 'vue';
import { stateSymbol, createState } from '@/store';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';

createApp(App)
  .use(router)
  .provide(stateSymbol, createState())
  .mount('#app');
