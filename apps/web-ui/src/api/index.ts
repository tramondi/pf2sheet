import { Sheet, Profile, Ancestry, Class } from '../model'

const apiHost = '//localhost:8081'

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

    const response = await fetch(`${apiHost}/api/sheets/`, {
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

    const response = await fetch(`${apiHost}/api/sheets/${sheet.id}`, {
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
    const response = await fetch(`${apiHost}/api/sheets/${id}`, {
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

    const response = await fetch(`${apiHost}/api/profile/`, {
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

export const getProfile = () => {
  return new Promise<Profile>(async (resolve, reject) => {
    const response = await fetch(`${apiHost}/api/profile/`, {
      method: 'GET',
      credentials: "include",
    })

    if (response.status != 200) {
      return reject(new Error(`response error code: ${response.status}`))
    }

    const body = await response.json()
    console.log(JSON.stringify(body))

    const profile = {
      displayName: body.data.display_name,
      login: body.data.login,
    } as Profile

    resolve(profile)
  })
}

export const getSheets = () => {
  return new Promise<Sheet[]>(async (resolve, reject) => {
    const response = await fetch(`${apiHost}/api/sheets/`, {
      method: 'GET',
      credentials: 'include',
    })

    if (response.status != 200) {
      return reject(new Error(`response error code: ${response.status}`))
    }

    const body = await response.json()
    console.log(JSON.stringify(body))

    const sheets: Sheet[] = body.data.sheets.map(item => ({
      id: item.id,
      charName: item.full_name,
      level: item.level,
      hpCurrent: item.hp_current,
      hpMax: item.hp_max,
      background: item.background,
      ancestryId: item.ancestry_id,
      classId: item.class_id,
    } as Sheet))

    return resolve(sheets)
  })
}

export const getAncestries = () => {
  return new Promise<Ancestry[]>(async (resolve, reject) => { 
    const response = await fetch(`${apiHost}/api/knowledge/ancestries`, {
      method: 'GET',
      credentials: 'include',
    })

    if (response.status != 200) {
      return reject(new Error(`response error code: ${response.status}`))
    }

    const body = await response.json()
    console.log(JSON.stringify(body))

    const ancestries: Ancestry[] = body.data.ancestries.map(item => ({
      id: item.id,
      title: item.title,
    } as Ancestry))

    resolve(ancestries)
  })
}

export const getClasses = () => {
  return new Promise<Class[]>(async (resolve, reject) => {
    const response = await fetch(`${apiHost}/api/knowledge/classes`, {
      method: 'GET',
      credentials: 'include',
    })

    if (response.status != 200) {
      return reject(new Error(`response error code: ${response.status}`))
    }

    const body = await response.json()
    console.log(JSON.stringify(body))

    const classes = body.data.classes.map(item => ({
      id: item.id,
      title: item.title,
    } as Class))

    resolve(classes)
  })
}

export const signout = () => {
  return new Promise<void>(async (resolve, reject) => {
    const response = await fetch(`${apiHost}/api/profile/signout`, {
      method: 'POST',
      credentials: 'include',
    })

    if (response.status != 200) {
      return reject(new Error(`response error code: ${response.status}`))
    }

    resolve()
  })
}

export const signin = (login: string, password: string) => {
  return new Promise<void>(async (resolve, reject) => {
    const params = {
      login: login,
      password: password,
    }

    const body = JSON.stringify(params)
    console.log(body)

    await fetch(`${apiHost}/auth/signin`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json; charset=utf8',
      },
      credentials: 'include',
      body: body,
    })
      .then(response => resolve())
      .catch(err => {
        const msg = `fetch error: ${err.message}`

        console.log(msg)
        reject(new Error(msg))
      })
  })
}

export const signup = (login: string, password: string, displayName: string) => {
  return new Promise<void>(async (resolve, reject) => {
    const params = {
      login: login,
      password: password,
      display_name: displayName,
    }

    const body = JSON.stringify(params)
    console.log(body)

    const response = await fetch(`${apiHost}/auth/signup`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json; charset=utf8',
      },
      credentials: 'include',
      body: body,
    })

    if (response.status != 200) {
      console.log('status ' + response.status)
      return reject(new Error(`response error code: ${response.status}`))
    }

    resolve()
  })
}
