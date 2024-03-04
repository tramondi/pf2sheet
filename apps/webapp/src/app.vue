<template>
  <v-app>
    <navbar/>
    <v-container>
      <main>
        <router-view></router-view>
      </main>
    </v-container>
  </v-app>
</template>

<script setup lang="ts">
import Navbar from './components/navbar.vue';
</script>
