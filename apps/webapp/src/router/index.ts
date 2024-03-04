import {
  createRouter,
  createWebHashHistory,
  RouteRecordRaw,
  Router,
} from 'vue-router'

import MySheets from '../views/my-sheets.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/sheet/:id',
    name: 'sheet',
    component: () => import('../views/sheet.vue'),
    props: true,
  },
  {
    path: '/my',
    name: 'my-sheets',
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
