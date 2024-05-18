export type Class = {
  id: number
  title: string
}

export type Ancestry = {
  id: number
  title: string
}

export type Sheet = {
  id?: number
  charName?: string
  level: number
  hpCurrent?: number
  hpMax?: number
  background?: string
  ancestryId?: number
  classId?: number
}
