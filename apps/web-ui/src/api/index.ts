import { Sheet, Profile } from '../model'

export const createSheet = async (sheet: Sheet): Promise<Sheet> => {
  return new Promise<Sheet>(async (resolve, reject) => {
    console.log(`creating sheet:`, sheet)

    const params = {
      ancestry_id: sheet.ancestryId,
      class_id: sheet.classId,
      background: sheet.background,
      full_name: sheet.charName,
      level: sheet.level > 0 ? sheet.level : undefined,
      hp_current: sheet.hpCurrent,
      hp_max: sheet.hpMax,
    }

    const body = JSON.stringify(params)
    console.log(body)

    const response = await fetch('//localhost:8081/api/sheets/', {
      method: 'POST',
      headers: {
        "Content-Type": "application/json; charset=utf8",
      },
      credentials: "include",
      body: body,
    })

    if (response.status != 200) {
      console.log('status ' + response.status)
      return reject(new Error(`response error code: ${response.status}`))
    }

    const data = (await response.json()).data
    sheet.id = data.sheet_id

    console.log('sheet_id:', data.sheet_id)

    resolve(sheet)
  })
}

export const updateSheet = (sheet: Sheet) => {
  return new Promise<Sheet>(async (resolve, reject) => {
    if (sheet.id === undefined) {
      return reject(new Error(`sheet.id is undefined`))
    }

    console.log(`updating sheet:`, sheet)

    const params = {
      ancestry_id: sheet.ancestryId,
      class_id: sheet.classId,
      background: sheet.background,
      full_name: sheet.charName,
      level: sheet.level > 0 ? sheet.level : undefined,
      hp_current: sheet.hpCurrent,
      hp_max: sheet.hpMax,
    }

    const body = JSON.stringify(params)
    console.log(body)

    const response = await fetch(`//localhost:8081/api/sheets/${sheet.id}`, {
      method: 'PUT',
      headers: {
        "Content-Type": "application/json; charset=utf8",
      },
      credentials: "include",
      body: body,
    })

    if (response.status != 200) {
      console.log('status ' + response.status)
      return reject(new Error(`response error code: ${response.status}`))
    }

    resolve(sheet)
  })
}

export const deleteSheet = async (id: number) => {
  return new Promise(async (resolve, reject) => {
    const response = await fetch(`//localhost:8081/api/sheets/${id}`, {
      method: 'DELETE',
      headers: {
        "Content-Type": "application/json; charset=utf8",
      },
      credentials: "include",
    })

    if (response.status != 200) {
      return reject(new Error(`response error code: ${response.status}`))
    }

    resolve(true)
  })
}

export const updateProfile = (profile: Profile) => {
  return new Promise<Profile>(async (resolve, reject) => {
    console.log(`updating profile:`, profile)

    const params = {
      display_name: profile.displayName,
    }

    const body = JSON.stringify(params)
    console.log(body)

    const response = await fetch(`//localhost:8081/api/profile/`, {
      method: 'PUT',
      headers: {
        "Content-Type": "application/json; charset=utf8",
      },
      credentials: "include",
      body: body,
    })

    if (response.status != 200) {
      console.log('status ' + response.status)
      return reject(new Error(`response error code: ${response.status}`))
    }

    resolve(profile)
  })
}
