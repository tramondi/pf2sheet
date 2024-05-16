<script setup lang="ts">
import { useField, useForm } from 'vee-validate'

import { useUserStore } from '../../../stores'

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

const submit = async (_) => {
  const params = {
    login: login.value.value,
    password: password.value.value,
    display_name: displayName.value.value,
  }

  const body = JSON.stringify(params)
  console.log(body)

  const response = await fetch('//localhost:8081/auth/signup', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json; charset=utf8",
    },
    credentials: "include",
    body: body,
  })

  if (response.status != 200) {
    console.log('status ' + response.statue)
    return
  }

  window.location.replace('/#/dashboard')
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

      <v-btn class="me-4" type="submit">Зарегистрироваться</v-btn>
    </form>
  </div>
</template>
