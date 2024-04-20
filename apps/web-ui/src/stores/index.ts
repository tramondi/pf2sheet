import { createPinia as _createPinia } from 'pinia'

export const createPinia = () => {
  return _createPinia()
}

export * from './user.store'
export * from './dashboard.store'
export * from './knowledge.store'
