import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/views/Home.vue';
import Download from '@/views/Download.vue';

Vue.use(Router);

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/download',
    name: 'download',
    component: Download,
  },
];

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
