<script setup lang="ts">
  import { VApp, VContainer } from 'vuetify/components'
</script>

<template>
  <VApp>
    <VContainer>
      <main>
        <router-view></router-view>
      </main>
    </VContainer>
  </VApp>
</template>
