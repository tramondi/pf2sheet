import './style.css'
// import '@mdi/font/css/materialdesignicons.css'
// import 'material-design-icons-iconfont/dist/material-design-icons.css' // Ensure your project is capable of handling css files

import { createPinia } from 'pinia'
import { createApp } from 'vue'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi-svg'
import * as vComponents from 'vuetify/components'
import * as vDirectives from 'vuetify/directives'

import { createAppRouter } from './router'

import App from './App.vue'

const router = createAppRouter()
const pinia = createPinia()

const vuetify = createVuetify({
  components: vComponents,
  directives: vDirectives,
  theme: {
    defaultTheme: 'dark',
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: { mdi },
  },
})

createApp(App)
  .use(router)
  .use(pinia)
  .use(vuetify)
  .mount('#app')
