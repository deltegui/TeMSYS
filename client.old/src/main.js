import Vue from 'vue'
import App from './App.vue'
import * as VueWin from '@hscmap/vue-window';
import * as VueMenu from '@hscmap/vue-menu';

Vue.config.productionTip = false;

if(navigator.serviceWorker) {
  navigator.serviceWorker.register('/service-worker.js');
}

Vue.use(VueWin);
Vue.use(VueMenu);

new Vue({
  render: h => h(App),
}).$mount('#app')
