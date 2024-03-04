<template>
  <v-card
    flat
    class="text-xs-center ma-3"
    width="270px"
    height="350px"
  >
    <v-card-text>
      <div class="subheading">{{ props.name ?? '<no-name>' }}</div>
    </v-card-text>
    <div class="d-flex justify-center align-center overflow-hidden">
      <v-avatar
        size="220px"
        rounded="0"
        :image="props.avatar_url ?? 'https://http.cat/404'"
      >
      </v-avatar>
    </div>
    <v-card-actions>
      <v-layout class="px-4 py-2 justify-end">
        <v-card-text>
          <p>item link</p>
        </v-card-text>
        <v-switch
          color="warning"
          hide-details="true"
          @update:model-value="switched"
          :model-value="active"
          :loading="loading"
          :disabled="disabled"
        ></v-switch>
        <div class="d-flex align-center pl-2">
          <v-btn @click="clicked">
            <v-icon small :icon="mdiLink"></v-icon>
          </v-btn>
        </div>
      </v-layout>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import { mdiLink } from '@mdi/js'

const props = defineProps<{
  name?: string,
  avatar_url?: string,
}>()

let active = ref(false)
let disabled = ref(false)
let loading = ref<String|Boolean>(false)

const switched = () => {
  loading.value = 'warning'
  disabled.value = true

  setTimeout(() => {
    disabled.value = false
    loading.value = false

    if (true) {
      active.value = !active.value
    }
  }, 1000)
}

const clicked = () => alert('link copied')
</script>
