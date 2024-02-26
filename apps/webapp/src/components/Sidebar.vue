<template>
    <div class="text-subtitle-2 mb-2">Default</div>
    <v-expansion-panels>
      <v-expansion-panel
        v-for="i in 3"
        :key="i"
        title="Item"
        text="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
      ></v-expansion-panel>
    </v-expansion-panels>
</template>

<script setup lang="ts">
  import { VExpansionPanels, VExpansionPanel } from 'vuetify/components';
</script>
