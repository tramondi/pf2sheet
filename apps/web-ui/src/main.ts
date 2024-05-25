import { createApp } from 'vue'

import CKEditor from '@ckeditor/ckeditor5-vue'

import './style.css'
import App from './App.vue'
import { createRouter } from './router'
import { createPinia } from './stores'
import { createVuetify } from './plugins'

const router = createRouter()
const store = createPinia()
const vuetify = createVuetify()

const app = createApp(App)
  .use(router)
  .use(store)
  .use(vuetify)
  .use(CKEditor)

router.isReady().then(() => app.mount('#app'))
