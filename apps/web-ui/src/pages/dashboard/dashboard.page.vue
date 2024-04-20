<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { mdiPlus } from '@mdi/js'

import { useDashboardStore, useKnowledgeStore } from '../../stores'
import SheetCard from '../../components/sheet-card.vue'

const dashboardStore = useDashboardStore()
const knowledgeStore = useKnowledgeStore()

knowledgeStore.load()

const { ancestries, classes } = storeToRefs(knowledgeStore)
console.log(`store ancestries: ${JSON.stringify(ancestries.value)}`)
console.log(`store classes: ${JSON.stringify(classes.value)}`)

const selectedClass = ref('')
const selectedAncestry = ref('')

const dialogModel = ref(false)
const tmpSheet = ref<Sheet>({level: 1})

const newSheet = () => {
  console.log(`tmpsheet: ${JSON.stringify(tmpSheet.value)}`)

  dialogModel.value = true
}

const closeModal = (saveSheet: boolean) => {
  if (saveSheet) {
    console.log(`saving sheet: ${JSON.stringify(tmpSheet.value)}`)
  }

  dialogModel.value = false
}

const w = '240px'
const h = '150px'
</script>

<template>
  <div>
    <v-row justify="space-between">
      <v-col cols="4" md="4">
        <v-card :min-width="w" :min-height="h">
          <v-btn :icon="mdiPlus" size="x-large" @click="newSheet"></v-btn>
        </v-card>
      </v-col>
      <v-col cols="4" md="4" v-for="sheet in dashboardStore.getSheets2">
        <SheetCard :sheet="sheet" :min-width="w" :min-height="h"/>
      </v-col>
    </v-row>
  </div>
  <div class="text-center pa-4">
    <v-dialog
      v-model="dialogModel"
      max-width="400"
      persistent
    >
      <v-card
        prepend-icon="mdi-map-marker"
        title="Создание листа персонажа"
      >
        <v-text-field label="Имя" v-model="tmpSheet.charName"></v-text-field>
        <v-text-field label="Уровень" v-model="tmpSheet.level"></v-text-field>
        <v-text-field label="Текущие ОЗ" v-model="tmpSheet.hitPointsCurrent"></v-text-field>
        <v-text-field label="Максимум ОЗ" v-model="tmpSheet.hitPointsMax"></v-text-field>
        <v-text-field label="Предыстория" v-model="tmpSheet.background"></v-text-field>
        <v-select
          label="Происхождение"
          v-model="selectedAncestry"
          item-value="code"
          :items="ancestries"
        ></v-select>
        <v-select
          label="Класс"
          v-model="selectedClass"
          item-value="code"
          :items="classes"
        ></v-select>
        <template v-slot:actions>
          <v-spacer></v-spacer>
          <v-btn @click="closeModal(false)">Отмена</v-btn>
          <v-btn @click="closeModal(true)">Сохранить</v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>
