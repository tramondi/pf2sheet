<script setup lang="ts">
import { ref } from 'vue'

import GeneralInfo from './blocks/GeneralInfo.vue'

const props = defineProps<{ hpPerLvl: number }>()

const lvl = ref(1)
</script>

<template>
  <div>
    <h1>Pathfinder 2e</h1>
  </div>
  <GeneralInfo
    charName="abobus"
    playerName="bibunus"
    ancestry="gnome"
    background="acrobat"
    klass="ranger"
  />
  <div>
    <div>
      <h3>Total HP: {{ lvl * props.hpPerLvl }}</h3>
    </div>
    <div>
      <button type="button" class="btn btn-blue" @click="lvl--">Lvl -</button>
      <button type="button" class="btn btn-blue" @click="lvl++">Lvl +</button>
    </div>
  </div>
</template>

<style scoped>
  .btn {
    @apply font-bold py-2 px-4 rounded;
  }
  .btn-blue {
    @apply bg-blue-500 text-white;
  }
  .btn-blue:hover {
    @apply bg-blue-700;
  }
</style>
