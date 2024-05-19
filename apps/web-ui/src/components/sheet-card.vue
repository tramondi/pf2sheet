<script setup lang="ts">
import { computed, ref } from 'vue'
import { storeToRefs } from 'pinia'

import { ListItem } from '../domain/types'
import { Sheet } from '../model'
import { updateSheet, deleteSheet } from '../api'

import {
  useKnowledgeStore,
  useDashboardStore,
} from '../stores'

const props = defineProps<{
  sheet?: Sheet
  minWidth: string
  minHeight: string
}>()

const knowledgeStore = useKnowledgeStore()
const { ancestries, classes } = storeToRefs(knowledgeStore)

const dashboardStore = useDashboardStore()

const sheet = ref(props.sheet)

knowledgeStore.load()
  .then((_) => {
    console.log(`store ancestries: ${JSON.stringify(ancestries.value)}`)
    console.log(`store classes: ${JSON.stringify(classes.value)}`)
  })

const subtitle = computed(() => {
  const lvl = sheet.value.level
  const bg = sheet.value.background
  const ancestryId = sheet.value.ancestryId
  const classId = sheet.value.classId

  let parts = []

  if (lvl !== undefined) {
    parts.push(`ур. ${lvl}`)
  }

  if (ancestryId !== undefined) {
    const ancestry = knowledgeStore.getAncestryById(ancestryId)
    if (ancestry !== undefined) {
      parts.push(ancestry.title)
    }
  }

  if (classId !== undefined) {
    const klass = knowledgeStore.getClassById(classId)
    if (klass !== undefined) {
      parts.push(klass.title)
    }
  }

  if (bg !== undefined) {
    parts.push(`(${bg})`)
  }

  return parts.join(' ')
})

const text = computed(() => {
  if (sheet.value.hpCurrent === undefined || sheet.value.hpMax === undefined) {
    return ''
  }

  const hpCurrent = sheet.value.hpCurrent
  const hpMax = sheet.value.hpMax

  return `${hpCurrent} / ${hpMax} ОЗ`
})

const dialogModel = ref(false)
const updateRequestStatus = ref(true)
const tmpSheet = ref<Sheet|undefined>()

const openModal = () => {
  tmpSheet.value = {...sheet.value}
  console.log(`tmpsheet: ${JSON.stringify(tmpSheet)}`)

  dialogModel.value = !dialogModel.value
}

const closeModal = (saveChanges: boolean) => {
  if (saveChanges && tmpSheet.value !== undefined) {
    tmpSheet.value.level = parseInt(tmpSheet.value.level)

    if (tmpSheet.value.hpCurrent !== undefined) {
      tmpSheet.value.hpCurrent = parseInt(tmpSheet.value.hpCurrent!)
    }

    if (tmpSheet.value.hpMax !== undefined) {
      tmpSheet.value.hpMax = parseInt(tmpSheet.value.hpMax!)
    }

    updateSheet(tmpSheet.value)
      .then(_ => {
        updateRequestStatus.value = true
        dialogModel.value = false

        sheet.value = {...tmpSheet.value}
        dashboardStore.reloadSheets()
      })
      .catch(err => {
        console.log(`failed to update sheet: ${err}`)
        updateRequestStatus.value = false
      })
  } else {
    dialogModel.value = false
  }
}

const btnDelete = (id: number) => {
  deleteSheet(id)
    .then(_ => {
      dashboardStore.reloadSheets()
    })
    .catch(err => {
      console.log(`create request failed: ${err.message}`)
    })
}

const btnExport = () => {
  const json = JSON.stringify(sheet.value, null, '\t')
  console.log(`sheet to export: ${json}`)

  const blob = new Blob([json], {
    type: 'application/json',
  })
  const blobUrl = URL.createObjectURL(blob)

  let href = document.createElement('a')
  href.href = blobUrl
  href.download = `sheet_${sheet.value.id}`

  href.click()
}

const clearSheet = () => {
  sheet.value = {} as Sheet
}
</script>

<template>
  <v-card
    :title="props.sheet.charName"
    :subtitle="subtitle"
    :text="text"
    :min-width="props.minWidth"
    :min-height="props.minHeight"
  >
    <v-btn @click="btnExport">Экспорт</v-btn>
    <v-btn @click="openModal">Изменить</v-btn>
    <v-btn @click="btnDelete(sheet.id)">Удалить</v-btn>
  </v-card>
  <div class="text-center pa-4">
    <v-dialog
      v-model="dialogModel"
      max-width="500"
    >
      <v-card
        prepend-icon="mdi-map-marker"
        title="Редактирование листа персонажа"
      >
        <v-text-field label="Имя" v-model="tmpSheet.charName"></v-text-field>
        <v-text-field label="Уровень" v-model="tmpSheet.level"></v-text-field>
        <v-text-field label="Текущие ОЗ" v-model="tmpSheet.hpCurrent"></v-text-field>
        <v-text-field label="Максимум ОЗ" v-model="tmpSheet.hpMax"></v-text-field>
        <v-text-field label="Предыстория" v-model="tmpSheet.background"></v-text-field>
        <v-select
          label="Происхождение"
          v-model="tmpSheet.ancestryId"
          item-value="id"
          :items="ancestries"
        ></v-select>
        <v-select
          label="Класс"
          v-model="tmpSheet.classId"
          item-value="id"
          :items="classes"
        ></v-select>
        <v-alert
          v-if="updateRequestStatus === false"
          color="error"
          icon="$error"
          title="Ошибка запроса"
          text="Видимо в одном из полей содержится ошибка, либо что-то произошло с сервером"
        ></v-alert>
        <template v-slot:actions>
          <v-spacer></v-spacer>
          <v-btn @click="clearSheet">Очистить</v-btn>
          <v-btn @click="closeModal(false)">Отмена</v-btn>
          <v-btn @click="closeModal(true)">Сохранить</v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>
