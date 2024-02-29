<template>
  <VCard flat class="text-xs-center ma-3">
    <VLayout row>
      <VCardText>
        <div class="subheading">({{ props.lvl }}) {{ props.character_name ?? '<no-name>' }}</div>
      </VCardText>
      <VCardActions>
        <VBtn @click="clicked" right small>
          <VIcon small :icon="mdiPencil"></VIcon>
        </VBtn>
        <VSwitch right
          @update:model-value="switched"
          color="success"
          :loading="loading"
          :disabled="disabled"
        ></VSwitch>
        <VSwitch></VSwitch>
        <VBtn @click="clicked" right small>
          <VIcon small :icon="mdiLink"></VIcon>
        </VBtn>
      </VCardActions>
    </VLayout>
  </VCard>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import { mdiLink, mdiPencil } from '@mdi/js'

const props = defineProps<{
  character_name?: String,
  lvl: Number,
  avatar?: String,
}>()

let loading = ref<String|Boolean>(false)
let disabled = ref(false)
const switched = () => {
  loading.value = 'warning'
  disabled.value = true

  setTimeout(() => {
    disabled.value = false
    loading.value = false
  }, 2000)
}

const clicked = () => alert('clicked')
</script>
