import { defineStore } from 'pinia'

import { Profile } from '../model'
import { getProfile } from '../api'

type UserState = {
  profile?: Profile
  loaded: boolean
}

export const useUserStore = defineStore<'user', UserState>({
  id: 'user',
  state: () => ({
    profile: undefined,
    loaded: false,
  }),
  actions: {
    async load() {
      if (this.loaded) {
        return
      }

      await getProfile()
        .then(profile => {
          this.profile = profile
          this.loaded = true
        })
        .catch(err => console.log(`unauthorized: ${err.message}`))
    },
  },
})
