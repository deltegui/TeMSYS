import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import Overview from '../views/Overview.vue';
import Login from '../views/Login.vue';
import Panel from '../views/Panel.vue';
import UserAdmin from '../views/UserAdmin.vue';
import SensorAdmin from '../views/SensorAdmin.vue';
import Historic from '../views/Historic.vue';

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
  {
    path: '/sensoradmin',
    name: 'SensorAdmin',
    component: SensorAdmin,
  },
  {
    path: '/historic',
    name: 'Historic',
    component: Historic,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
