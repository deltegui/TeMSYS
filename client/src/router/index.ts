import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Overview from '../views/Overview.vue';
import Login from '../views/Login.vue';
import Panel from '../views/Panel.vue';
import UserAdmin from '../views/UserAdmin.vue';

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
  {
    path: '/panel',
    name: 'Panel',
    component: Panel,
  },
  {
    path: '/useradmin',
    name: 'UserAdmin',
    component: UserAdmin,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;