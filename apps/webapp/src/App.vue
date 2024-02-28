<template>
  <VApp>
    <Navbar/>
    <VContainer>
      <main>
        <router-view></router-view>
      </main>
    </VContainer>
  </VApp>
</template>

<script setup lang="ts">
import Navbar from './components/Navbar.vue';
</script>
