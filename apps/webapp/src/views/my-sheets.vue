<template>
  <v-container>
    <v-layout row wrap>
      <v-flex xs12 sm6 md4 lg3 v-for="sheet in data" :key="sheet.name">
        <SheetCard :character_name="sheet.name" :lvl="sheet.avatar_url"/>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script setup lang="ts">
import SheetCard from '../components/sheet-card.vue'

const data = [
  { name: null, avatar_url: '' },
  { name: 'biba', avatar_url: '' },
  { name: 'boba', avatar_url: '' },
]
</script>
