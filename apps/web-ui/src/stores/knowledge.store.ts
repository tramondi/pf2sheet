import { defineStore } from 'pinia'

import { Class, Ancestry } from '../model'
import { getAncestries, getClasses } from '../api'

const storeName = 'knowledge'

type KnowledgeState = {
  ancestries: Ancestry[]
  classes: Class[]
}

export const useKnowledgeStore = defineStore(storeName, {
  state: () => ({
    ancestries: [],
    classes: [],
  }),
  actions: {
    async load() {
      this.ancestries = await getAncestries()
      this.classes = await getClasses()
    },

    getAncestryById(id: number): Ancestry | undefined {
      return this.ancestries.find(item => item.id == id)
    },

    getClassById(id: number): Class | undefined {
      return this.classes.find(item => item.id == id)
    },
  },
})
