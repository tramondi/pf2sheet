import { defineStore } from 'pinia'
import axios from 'axios'

type AuthDTO = {
  profile: {
    name: string
  }
}

type ResponseError = {
  response: {
    data: {
      detail: string
    }
  }
}

type UserState = {
  profile: {
    privateKey: string
    name: string
  } | null
  returnUrl: string
};

export const useUserStore = defineStore<'moderator', UserState>({
  id: 'moderator',
  state: () => ({
    profile: null,
    returnUrl: '',
    token: '',
  }),
  actions: {
    async auth(privateKey: string) {
      try {
        const result = await axios<AuthDTO>({
          url: '/auth',
          method: 'POST',
          data: new URLSearchParams({
            privateKey,
          }),
        });

        return result.data
      } catch (err) {
        throw (err as ResponseError).response.data.detail
      }
    },
  },
})
