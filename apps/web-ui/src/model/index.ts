export type Class = {
  code: string
  title: string
}

export type Ancestry = {
  code: string
  title: string
}

export type Sheet = {
  charName?: string
  level: number
  hpCurrent?: number
  hpMax?: number
  background?: string
  ancestry?: Ancestry
  class?: Class
}
