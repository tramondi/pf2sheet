<script setup lang="ts">
</script>

<template>
  <v-layout>
    <v-app-bar
      color="grey-darken-2"
      density="compact"
    >
      <router-link class="link" to="/profile">Профиль</router-link>
      <router-link class="link" to="/dashboard">Листы</router-link>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>
  </v-layout>
</template>

<style>
.link {
  padding: 16px 32px;
  color: white;
}

.link:hover {
  background: darkgray;
}

.router-link-active {
  background: gray;
}
</style>
