<script lang="ts" setup>
import { ref } from 'vue'

import Signin from './components/signin.vue'
import Signup from './components/signup.vue'

const tab = ref<string>()

const returnUrl = '/#/dashboard'
</script>

<template>
  <v-card :width="400" :height="400">
    <v-tabs v-model="tab" align-tabs="center">
      <v-tab value="signin">Вход</v-tab>
      <v-tab value="signup">Регистрация</v-tab>
    </v-tabs>

    <v-card-text>
      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="signin">
          <Signin :return-url="returnUrl"/>
        </v-tabs-window-item>
        <v-tabs-window-item value="signup">
          <Signup :return-url="returnUrl"/>
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card-text>
  </v-card>
</template>
