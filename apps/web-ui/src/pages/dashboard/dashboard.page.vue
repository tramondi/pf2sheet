<script setup lang="ts">
import { ref, computed } from 'vue'
import { storeToRefs } from 'pinia'
import { mdiPlus } from '@mdi/js'

import {
  useDashboardStore,
  useKnowledgeStore,
  useUserStore,
} from '../../stores'
import SheetCard from '../../components/sheet-card.vue'
import { createSheet } from '../../api'

const knowledgeStore = useKnowledgeStore()

const dashboardStore = useDashboardStore()

const userStore = useUserStore()
userStore.load()

const { profile } = storeToRefs(userStore)
const playerTitle = computed(() => {
  const playerLoginPart = `Player ${profile.value.login}`
  const playerNamePart = profile.value.displayName == null
    ? '' : `// ${profile.value.displayName}`

  return `${playerLoginPart} ${playerNamePart}`
})

knowledgeStore.load()
  .then((_) => {
    console.log(`store ancestries: ${JSON.stringify(ancestries.value)}`)
    console.log(`store classes: ${JSON.stringify(classes.value)}`)

    dashboardStore.load()
  })

const { ancestries, classes } = storeToRefs(knowledgeStore)
const { sheets } = storeToRefs(dashboardStore)

const dialogModel = ref(false)
const createRequestStatus = ref(true)
const tmpSheet = ref<Sheet>({level: 1})

const newSheet = () => {
  console.log(`tmpsheet: ${JSON.stringify(tmpSheet.value)}`)

  dialogModel.value = true
}

const closeModal = (saveSheet: boolean) => {
  if (saveSheet && tmpSheet.value !== undefined) {
    console.log(`saving sheet: ${JSON.stringify(tmpSheet.value)}`)

    if (tmpSheet.value.hpCurrent !== undefined) {
      tmpSheet.value.hpCurrent = parseInt(tmpSheet.value.hpCurrent!)
    }

    if (tmpSheet.value.hpMax !== undefined) {
      tmpSheet.value.hpMax = parseInt(tmpSheet.value.hpMax!)
    }

    createSheet(tmpSheet.value)
      .then(sheet => {
        dashboardStore.reloadSheets()
        createRequestStatus.value = true
      })
      .catch(err => {
        console.log(`create request failed: ${err.message}`)
        createRequestStatus.value = false
      })
  }

  dialogModel.value = false
}

const w = '240px'
const h = '150px'
</script>

<template>
  <h1 class="my-4">{{ playerTitle }}</h1>
  <div>
    <v-row justify="space-between">
      <v-col cols="4" md="4">
        <v-card :min-width="w" :min-height="h">
          <v-btn :icon="mdiPlus" size="x-large" @click="newSheet"></v-btn>
        </v-card>
      </v-col>
      <v-col cols="4" md="4" v-for="sheet in sheets">
        <SheetCard :sheet="sheet" :min-width="w" :min-height="h"/>
      </v-col>
    </v-row>
  </div>
  <div class="text-center pa-4">
    <v-dialog
      v-model="dialogModel"
      max-width="400"
    >
      <v-card
        prepend-icon="mdi-map-marker"
        title="Создание листа персонажа"
      >
        <v-text-field label="Имя" v-model="tmpSheet.charName"></v-text-field>
        <v-text-field
          label="Уровень"
          type="number"
          v-model="tmpSheet.level"
        ></v-text-field>
        <v-text-field
          label="Текущие ОЗ"
          type="number"
          v-model="tmpSheet.hpCurrent"
        ></v-text-field>
        <v-text-field
          label="Максимум ОЗ"
          type="number"
          v-model="tmpSheet.hpMax"
        ></v-text-field>
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
          v-if="createRequestStatus === false"
          color="error"
          icon="$error"
          title="Ошибка запроса"
          text="Видимо в одном из полей содержится ошибка, либо что-то произошло с сервером"
        ></v-alert>
        <template v-slot:actions>
          <v-spacer></v-spacer>
          <v-btn @click="closeModal(false)">Отмена</v-btn>
          <v-btn @click="closeModal(true)">Сохранить</v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>
