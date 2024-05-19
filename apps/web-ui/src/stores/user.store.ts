import { defineStore } from 'pinia'

import { Profile } from '../model'

type UserState = {
  profile?: Profile
  loaded: boolean
};

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

      const response = await fetch('//localhost:8081/api/profile/', {
        method: 'GET',
        credentials: "include",
      })

      if (response.status != 200) {
        return
      }

      const body = await response.json()
      console.log(JSON.stringify(body))

      this.profile = {
        displayName: body.data.display_name,
        login: body.data.login,
      }

      this.loaded = true
    },
  },
})
