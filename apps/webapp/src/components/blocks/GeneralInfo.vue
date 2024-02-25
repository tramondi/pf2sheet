<script setup lang="ts">
  const props = defineProps<{
    charName: string,
    playerName: string,
    ancestry: string,
    background: string,
    klass: string,
  }>()
</script>

<template>
  <div>
    <p>Character name: {{ props.charName }}</p>
    <p>Player name: {{ props.playerName }}</p>
    <p>Ancestry: {{ props.ancestry }}</p>
    <p>Background: {{ props.background }}</p>
    <p>Class: {{ props.klass }}</p>
  </div>
</template>
