import { defineStore } from 'pinia'
// import axios from 'axios'

import { Ancestry, Class, Sheet } from '../model'

const storeName = 'dashboard'

type DashboardState = {
  sheets: Sheet[]
}

const stubSheets: Sheet[] = [
  {
    charName: 'erza',
    ancestry: {code: 'gnome', title: 'Гном'},
    background: 'preacher',
    class: {code: 'bard', title: 'Бард'},
    level: 2,
    hitPointsMax: 16,
    hitPointsCurrent: 12,
  },
  {
    charName: 'boober',
    ancestry: {code: 'human', title: 'Человек'},
    background: 'miner',
    class: {code: 'fighter', title: 'Воин'},
    level: 1,
    hitPointsMax: 10,
    hitPointsCurrent: 10,
  },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
  // {
  //   charName: 'erza',
  //   ancestry: Ancestry.Gnome,
  //   background: 'student',
  //   class: Class.Bard,
  //   level: 2,
  //   hitPointsMax: 16,
  //   hitPointsCurrent: 12,
  // },
  // {
  //   charName: 'boober',
  //   ancestry: Ancestry.Human,
  //   background: 'miner',
  //   class: Class.Fighter,
  //   level: 1,
  //   hitPointsMax: 10,
  //   hitPointsCurrent: 10,
  // },
]

export const useDashboardStore = defineStore(storeName, {
  state: () => ({
    sheets: stubSheets,
  }),
  // actions: {
  // },
  getters: {
    getSheets: (state) => state.sheets,
    getSheets2: (state) => [state.sheets[0], state.sheets[1]],
  },
})
