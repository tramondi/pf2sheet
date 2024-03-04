import './style.css'

import { createPinia } from 'pinia'
import { createApp } from 'vue'

import { createAppRouter } from './router'
import { createVuetify } from './plugins'

import App from './app.vue'

const router = createAppRouter()
const pinia = createPinia()
const vuetify = createVuetify()

createApp(App)
  .use(router)
  .use(pinia)
  .use(vuetify)
  .mount('#app')
