<script setup lang="ts">
import { computed, ref } from 'vue'
import { ListItem } from '../domain/types'

import { Sheet } from '../model'

const props = defineProps<{
  sheet?: Sheet
  minWidth: string
  minHeight: string
}>()

const sheet = ref(props.sheet)

const subtitle = computed(() => {
  const lvl = sheet.value.level
  const ancestry = sheet.value.ancestry.title
  const klass = sheet.value.class.title
  const bg = sheet.value.background

  return `ур. ${lvl} ${ancestry} ${klass} (${bg})`
})

const text = computed(() => {
  const hpCurrent = sheet.value.hitPointsCurrent
  const hpMax = sheet.value.hitPointsMax

  return `${hpCurrent} / ${hpMax} hp`
})

const dialogModel = ref(false)
const tmpSheet = ref<Sheet|undefined>()

const itemsClass: ListItem[] = [
  {value: 'fighter', title: 'Воин'},
  {value: 'bard', title: 'Бард'},
  // {value: 'cleric', title: 'Жрец'},
  // {value: 'druid', title: 'Друид'},
  // {value: 'wizard', title: 'Волшебник'},
]
const selectedClass = ref('')

const openModal = () => {
  tmpSheet.value = {...sheet.value}
  console.log(`tmpsheet: ${JSON.stringify(tmpSheet)}`)

  //selectedClass.value = itemsClass.find(k => k.value === tmpSheet.value.class.value)
  selectedClass.value = tmpSheet.value.class.value
  dialogModel.value = !dialogModel.value
}

const closeModal = (saveChanges: boolean) => {
console.log(`selected: ${JSON.stringify(selectedClass.value)}`)
  if (saveChanges && tmpSheet.value !== undefined) {
    //props.sheet.level = tmpSheet.value.level
    sheet.value = {...tmpSheet.value}
    sheet.value.class = itemsClass.find(k => k.value === selectedClass.value)
    console.log(`sheeted: ${JSON.stringify(sheet.value.class)}`)
  }

  dialogModel.value = !dialogModel.value
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
    <v-btn @click="openModal">Изменить</v-btn>
  </v-card>
  <div class="text-center pa-4">
    <v-dialog
      v-model="dialogModel"
      max-width="400"
      persistent
    >
      <v-card
        prepend-icon="mdi-map-marker"
        title="Редактирование листа персонажа"
      >
        <v-text-field label="Имя" v-model="tmpSheet.charName"></v-text-field>
        <v-text-field label="Уровень" v-model="tmpSheet.level"></v-text-field>
        <v-text-field label="Текущие ОЗ" v-model="tmpSheet.hitPointsCurrent"></v-text-field>
        <v-text-field label="Максимум ОЗ" v-model="tmpSheet.hitPointsMax"></v-text-field>
        <v-text-field label="Предыстория" v-model="tmpSheet.background"></v-text-field>
        <v-text-field label="Происхождение" v-model="tmpSheet.ancestry"></v-text-field>
        <v-select
          label="Класс"
          v-model="selectedClass"
          :items="itemsClass"
        ></v-select>
        <!-- <v-text-field label="Класс" v-model="tmpSheet.class"></v-text-field> -->
        <template v-slot:actions>
          <v-spacer></v-spacer>
          <v-btn @click="closeModal(false)">Отмена</v-btn>
          <v-btn @click="closeModal(true)">Сохранить</v-btn>
        </template>
      </v-card>
    </v-dialog>
  </div>
</template>
