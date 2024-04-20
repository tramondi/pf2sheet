import {
  createRouter as _createRouter,
  createWebHashHistory,
} from 'vue-router'

export const createRouter = () => {
  return _createRouter({
    // eslint-disable-next-line
    // @ts-ignore
    history: createWebHashHistory(),
    routes: [
      {
        path: '/',
        name: 'home',
        component: () => import('../pages/home/home.page.vue'),
      },
      {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('../pages/dashboard/dashboard.page.vue'),
      },
      {
        path: '/auth',
        name: 'auth',
        component: () => import('../pages/auth/auth.page.vue'),
      },
    ],
  })
}
