<template>
  <div class="signup-container">
    <form @submit.prevent="submitSignup">
      <div class="input-group">
        <label for="email">Email</label>
        <input type="email" id="email" v-model="email" name="email" autocomplete="email" required />
      </div>

      <div class="input-group">
        <label for="password">Password</label>
        <input type="password" id="password" v-model="password" name="password" required />
      </div>

      <div class="input-group">
        <label for="password-repeat">Password Repeat</label>
        <input type="password" id="password-repeat" v-model="passwordRepeat" name="password-repeat" required />
      </div>

      <div class="actions">
        <button class="primary-button" type="submit">Signup</button>
      </div>

      <div v-if="errorMessage" class="error">
        {{ errorMessage }}
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user-store';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const email = ref('')
const password = ref('')
const passwordRepeat = ref("")
const errorMessage = ref('')
const router = useRouter()

const submitSignup = async () => {

  const userStore = useUserStore()

  userStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)
  const res = await userStore.signupPlayer(email.value, password.value, passwordRepeat.value)

  if (res instanceof Error) {
    errorMessage.value = res.message
    return
  }

  errorMessage.value = ''
  router.push('/login')
}
</script>

<style scoped>
.signup-container {
  max-width: 400px;
  margin: auto;
  padding: 1rem;
  border: 1px solid darkgray;
  border-radius: 8px;
}

.input-group {
  margin-bottom: 1rem;
}

.input-group label {
  display: block;
  margin-bottom: 0.5rem;
}

.input-group input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.actions {
  text-align: center;
}

.error {
  color: red;
  margin-top: 1rem;
  text-align: center;
}
</style>
