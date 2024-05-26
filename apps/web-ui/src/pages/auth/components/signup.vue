<script setup lang="ts">
import { ref } from 'vue'
import { useField, useForm } from 'vee-validate'
import VueRecaptcha from 'vue3-recaptcha2'
//import { useRecaptcha } from 'vue3-recaptcha2'

import { useUserStore } from '../../../stores'
import { signup } from '../../../api'

const props = defineProps<{
  returnUrl: string
}>()

const pattern = /[-0-9a-zA-Z_.]{0,30}/

const { handleSubmit } = useForm({
  validationSchema: {
    login(value) {
      if (value?.length < 2 && value?.length > 30) {
        return 'Длина имени должна быть от 2 до 30 символов'
      }

      if (!pattern.test(value)) {
        return 'Неверный формат логина'
      }

      return true
    },
    password(value) {
      if (value?.length < 6 && value?.length > 30) {
        return 'Длина пароля должна быть от 6 до 30 символов'
      }

      if (!pattern.test(value)) {
        return 'Неверный формат пароля'
      }

      return true
    },
    displayName(value) {
      if (value?.length > 30) {
        return 'Длина имени должна быть до 30 символов'
      }

      if (!pattern.test(value)) {
        return 'Неверный формат имени'
      }

      return true
    },
  },
})

const login = useField('login')
const password = useField('password')
const displayName = useField('displayName')

const recaptchaOK = ref(false)

const submit = async (_) => {
  if (!recaptchaOK.value) {
    alert('Капчу!')
    return
  }

  recaptchaOK.value = false

  await signup(
    login.value.value,
    password.value.value,
    displayName.value.value,
  )
    .then(_ => window.location.href = props.returnUrl)
    .catch(err => console.log(`signup error: ${err.message}`))
}

const recaptchaVerified = (response) => {
  console.log(`recaptcha verified`)
  recaptchaOK.value = true
}

const recaptchaExpired = () => {
  console.log(`recaptcha expired`)
  recaptchaOK.value = false
}

const recaptchaError = (reason) => {
  console.log(`recaptcha error:`, reason)
}

const recaptchaFailed = () => {
  console.log(`recaptcha failed`)
  recaptchaOK.value = false
}
</script>

<template>
  <div>
    <form @submit.prevent="submit" class="px-4 py-4">
      <v-text-field
        v-model="login.value.value"
        :error-messages="login.errorMessage.value"
        label="Логин*"
      ></v-text-field>

      <v-text-field
        v-model="password.value.value"
        :error-messages="password.errorMessage.value"
        label="Пароль*"
      ></v-text-field>

      <v-text-field
        v-model="displayName.value.value"
        :error-messages="displayName.errorMessage.value"
        label="Имя"
      ></v-text-field>

      <vue-recaptcha
        sitekey="6Ld3wugpAAAAACjleYYzRfLFVvh4V40WV2N25Cov"
        size="normal" 
        theme="light"
        hl="tr"
        :loadingTimeout="30000"
        @verify="recaptchaVerified"
        @expire="recaptchaExpired"
        @fail="recaptchaFailed"
        @error="recaptchaError"
        ref="vueRecaptcha"
      ></vue-recaptcha>

      <v-btn class="my-4" type="submit">Зарегистрироваться</v-btn>
    </form>
  </div>
</template>
