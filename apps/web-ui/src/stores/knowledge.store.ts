import { defineStore } from 'pinia'

import { Class, Ancestry } from '../model'

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
      const response = await fetch('//localhost:8081/api/knowledge/ancestries', {
        method: 'GET',
      })

      const body = await response.json()

      console.log(JSON.stringify(body))

      this.ancestries = body.data.ancestries.map(item => ({
        code: item.code,
        title: item.title,
      }))
    }
  },
})
