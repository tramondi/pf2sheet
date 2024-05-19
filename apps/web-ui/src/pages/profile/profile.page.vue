<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'

import { Profile } from '../../model'
import { updateProfile } from '../../api'
import { useUserStore } from '../../stores'

const userStore = useUserStore()
userStore.load()

const { profile } = storeToRefs(userStore)

const tmpProfile = ref<Profile>({...profile.value})
const saveRequestStatus = ref(true)

const btnDropChanges = () => {
  tmpProfile.value = {...profile.value}
}

const btnSaveProfile = () => {
  updateProfile(tmpProfile.value)
    .then(_ => {
      profile.value = {...tmpProfile.value}
      saveRequestStatus.value = true
    })
    .catch(err => {
      console.log(`update request failed: ${err.message}`)
      saveRequestStatus.value = false
    })
}
</script>

<template>
  <div class="text-center">
    <v-card>
      <v-card-title>Профиль {{ profile.login }}</v-card-title>
      <v-spacer></v-spacer>
      <v-text-field
        label="Отображаемое имя"
        type="text"
        v-model="tmpProfile.displayName"
      ></v-text-field>
      <v-alert
        v-if="saveRequestStatus === false"
        color="error"
        icon="$error"
        title="Ошибка запроса"
        text="Видимо в одном из полей содержится ошибка, либо что-то произошло с сервером"
      ></v-alert>
      <v-spacer></v-spacer>
      <template v-slot:actions>
        <v-spacer></v-spacer>
        <v-btn @click="btnDropChanges">Сбросить изменения</v-btn>
        <v-btn @click="btnSaveProfile">Сохранить</v-btn>
      </template>
    </v-card>
  </div>
</template>
