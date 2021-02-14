import { createApp } from 'vue';
import { stateSymbol, store } from '@/store';
import App from './App.vue';
import './registerServiceWorker';
import './debug';
import router from './router';

createApp(App)
  .use(router)
  .provide(stateSymbol, store)
  .mount('#app');
