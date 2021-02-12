import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Overview from '../views/Overview.vue';
import Login from '../views/Login.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Overview',
    component: Overview,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
