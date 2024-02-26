import './style.css'
import '@mdi/font/css/materialdesignicons.css'

import { createPinia } from 'pinia'
import { createApp } from 'vue'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as vDirectives from 'vuetify/directives'

import { createAppRouter } from './router'

import App from './App.vue'

const router = createAppRouter()
const pinia = createPinia()

const vuetify = createVuetify({
  directives: vDirectives,
  theme: {
    defaultTheme: 'dark',
  },
})

createApp(App)
  .use(router)
  .use(pinia)
  .use(vuetify)
  .mount('#app')
