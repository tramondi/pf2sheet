<script setup lang="ts">
import { useField, useForm } from 'vee-validate'

import { useUserStore } from '../../../stores'
import { signin } from '../../../api'

const props = defineProps<{
  returnUrl: string
}>()

const pattern = /[-0-9a-zA-Z_.]/

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
  },
})

const login = useField('login')
const password = useField('password')

const submit = async (_) => {
  await signin(
    login.value.value,
    password.value.value,
  )
    .then(_ => window.location.href = props.returnUrl)
    .catch(err => console.log(`signin error: ${err.message}`))
}
</script>

<template>
  <div>
    <form @submit.prevent="submit" class="px-4 py-4">
      <v-text-field
        v-model="login.value.value"
        :error-messages="login.errorMessage.value"
        label="Логин"
      ></v-text-field>

      <v-text-field
        v-model="password.value.value"
        :error-messages="password.errorMessage.value"
        label="Пароль"
      ></v-text-field>

      <v-btn class="me-4" type="submit">Войти</v-btn>
    </form>
  </div>
</template>
