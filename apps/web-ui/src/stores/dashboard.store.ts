import { defineStore } from 'pinia'

import { Sheet } from '../model'
import { getSheets } from '../api'

const storeName = 'dashboard'

type DashboardState = {
  sheets: Sheet[]
}

export const useDashboardStore = defineStore(storeName, {
  state: () => ({
    sheets: [],
  }),
  actions: {
    async load() {
      this.sheets = await getSheets()
    },

    async reloadSheets() {
      this.sheets = []
      await this.load()
    },
  },
  getters: {
    getSheets: (state) => state.sheets,
  },
})
