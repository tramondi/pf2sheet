<template>
  <nav>
    <v-toolbar app>
      <v-toolbar-title class="text-uppercase">
        <span>pf2</span>
        <span class="font-weight-light">sheet</span>
      </v-toolbar-title>
      <v-btn color="gray">
        <span>выход</span>
        <v-icon right :icon="mdiLogout"></v-icon>
      </v-btn>
    </v-toolbar>
  </nav>
</template>

<script setup lang="ts">
import { mdiLogout } from '@mdi/js'
</script>
