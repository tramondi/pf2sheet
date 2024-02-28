import {
  createRouter,
  createWebHashHistory,
  RouteRecordRaw,
  Router,
} from 'vue-router'

import Landing from '../views/Landing.vue'
import MySheets from '../views/MySheets.vue'

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
  {
    path: '/my',
    name: 'mysheets',
    component: MySheets,
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
