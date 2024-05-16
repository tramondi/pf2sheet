import { defineStore } from 'pinia'

import { Class, Ancestry } from '../model'

const storeName = 'knowledge'

type KnowledgeState = {
  ancestries: Ancestry[]
  classes: Class[]
}

const fetchAncestries = async () => {
  const response = await fetch('//localhost:8081/api/knowledge/ancestries', {
    method: 'GET',
    credentials: 'include',
  })

  const body = await response.json()
  console.log(JSON.stringify(body))

  return body.data.ancestries.map(item => ({
    id: item.id,
    title: item.title,
  }))
}

const fetchClasses = async () => {
  const response = await fetch('//localhost:8081/api/knowledge/classes', {
    method: 'GET',
    credentials: 'include',
  })

  const body = await response.json()
  console.log(JSON.stringify(body))

  return body.data.classes.map(item => ({
    id: item.id,
    title: item.title,
  }))
}

export const useKnowledgeStore = defineStore(storeName, {
  state: () => ({
    ancestries: [],
    classes: [],
  }),
  actions: {
    async load() {
      this.ancestries = await fetchAncestries()
      this.classes = await fetchClasses()
    }
  },
})
