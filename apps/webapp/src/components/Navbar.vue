<template>
  <nav>
    <VToolbar app>
      <VToolbarTitle class="text-uppercase">
        <span>pf2</span>
        <span class="font-weight-light">sheet</span>
      </VToolbarTitle>
      <VBtn color="gray">
        <span>выход</span>
        <VIcon right :icon="mdiLogout"></VIcon>
      </VBtn>
    </VToolbar>
  </nav>
</template>

<script setup lang="ts">
import { mdiLogout } from '@mdi/js'
</script>
