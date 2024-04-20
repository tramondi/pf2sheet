import { Ancestry, Class, ListItem } from '../types'

export type Sheet = {
  charName?: string
  level: number
  hitPointsCurrent?: number
  hitPointsMax?: number
  background?: string
  ancestry?: Ancestry
  class?: ListItem
}
