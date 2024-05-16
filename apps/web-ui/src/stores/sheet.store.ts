import { defineStore } from 'pinia'

import { Ancestry, Class, Sheet } from '../model'

const storeName = 'sheet'

type SheetDto = {
  fullName?: string
  level?: number
  hpCurrent?: number
  hpMax?: number
  background?: string
  ancestryId?: number
  classId?: number
}

type SheetState = {
  sheets: SheetDto[]
}

// const stubSheets: Sheet[] = [
//   {
//     charName: 'erza',
//     ancestry: {code: 'gnome', title: 'Гном'},
//     background: 'preacher',
//     class: {code: 'bard', title: 'Бард'},
//     level: 2,
//     hitPointsMax: 16,
//     hitPointsCurrent: 12,
//   },
//   {
//     charName: 'boober',
//     ancestry: {code: 'human', title: 'Человек'},
//     background: 'miner',
//     class: {code: 'fighter', title: 'Воин'},
//     level: 1,
//     hitPointsMax: 10,
//     hitPointsCurrent: 10,
//   },
// ]

export const useSheetStore = defineStore(storeName, {
  state: () => ({
    sheets: [],
  }),
  actions: {
    async load() {
      const response = await fetch('//localhost:8081/api/sheets/', {
        method: 'GET',
      })

      const body = await response.json()
      console.log(JSON.stringify(body))

      this.sheets = body.data.sheets.map(item => {
        charName: item?.fullName,
        level: item?.level,
        ancestryId: item?.ancestryId,
        classId: item?.classId,
        background: item?.background,
        hpCurrent: item?.hpCurrent,
        hpMax: item?.hpMax,
      })
    },
  },
  getters: {
    getSheets: (state) => state.sheets,
    getSheets2: (state) => [state.sheets[0], state.sheets[1]],
  },
})
