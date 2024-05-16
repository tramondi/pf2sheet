import {
  createRouter as _createRouter,
  createWebHashHistory,
} from 'vue-router'

import { useUserStore } from '../stores'

export const createRouter = () => {
  const router = _createRouter({
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

  router.beforeEach(async (to) => {
    console.log('to: ' + to.path)

    const publicPages = ['/auth']
    const authRequired = !publicPages.includes(to.path)

    const userStore = useUserStore()
    await userStore.load()

    if (authRequired && userStore.profile == null) {
      console.log('redierect to auth')
      return '/auth'
    }
  })

  return router
}
