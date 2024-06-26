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
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('../pages/dashboard/dashboard.page.vue'),
      },
      {
        path: '/auth',
        name: 'auth',
        component: () => import('../pages/auth/auth.page.vue'),
      },
      {
        path: '/profile',
        name: 'profile',
        component: () => import('../pages/profile/profile.page.vue')
      },
    ],
  })

  router.beforeEach(async (to) => {
    console.log('to: ' + to.path)

    const publicPages = ['/auth']
    const authRequired = !publicPages.includes(to.path)

    const userStore = useUserStore()
    await userStore.load()

    if (authRequired && userStore.profile === undefined) {
      console.log('redirect to auth')
      return '/auth'
    }
  })

  return router
}
