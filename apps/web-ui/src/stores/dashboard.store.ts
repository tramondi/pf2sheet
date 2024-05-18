import { defineStore } from 'pinia'

import { Sheet } from '../model'

const storeName = 'dashboard'

type DashboardState = {
  sheets: Sheet[]
}

const fetchSheets = async () => {
  const response = await fetch('//localhost:8081/api/sheets/', {
    method: 'GET',
    credentials: 'include',
  })

  const body = await response.json()
  console.log(JSON.stringify(body))

  return body.data.sheets.map(item => ({
    id: item.id,
    charName: item.full_name,
    level: item.level,
    hpCurrent: item.hp_current,
    hpMax: item.hp_max,
    background: item.background,
    ancestryId: item.ancestry_id,
    classId: item.class_id,
  }))
}

export const useDashboardStore = defineStore(storeName, {
  state: () => ({
    sheets: [],
  }),
  actions: {
    async load() {
      this.sheets = await fetchSheets()
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
