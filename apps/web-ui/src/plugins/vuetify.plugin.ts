import 'vuetify/styles'
import { createVuetify as _createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi-svg'

export const createVuetify = () => _createVuetify({
  components,
  directives,
  icons: {
    defaultSet: 'mdi',
    aliases: {
      ...aliases,
      // account: mdiAccount,
    },
    sets: {
      mdi,
    },
  },
  theme: {
    defaultTheme: 'light',
  },
})
