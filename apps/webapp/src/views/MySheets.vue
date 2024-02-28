<template>
  <VContainer>
    <VLayout row wrap>
      <VFlex xs12 sm6 md4 lg3 v-for="sheet in data" :key="sheet.name">
        <SheetCard :character_name="sheet.name" :lvl="sheet.lvl"/>
      </VFlex>
    </VLayout>
  </VContainer>
</template>

<script setup lang="ts">
import SheetCard from '../components/SheetCard.vue'

const data = [
  { name: null, lvl: 1 },
  { name: 'biba', lvl: 4 },
  { name: 'boba', lvl: 11 },
]
</script>
