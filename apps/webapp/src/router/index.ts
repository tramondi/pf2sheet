import {
  createRouter,
  createWebHashHistory,
  RouteRecordRaw,
  Router,
} from 'vue-router'

import Landing from '../views/Landing.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'landing',
    component: Landing,
  },
  {
    path: '/sheet/:id',
    name: 'sheet',
    component: () => import('../views/Sheet.vue'),
    props: true,
  },
]

export const createAppRouter = () => {
  const router: Router = createRouter({
    routes: routes,
    history: createWebHashHistory(),
  })

  // router.beforeEach()
  
  return router
}

export default createAppRouter
