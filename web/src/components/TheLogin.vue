<template>
  <div class="login-container">
    <form @submit.prevent="submitLogin">
      <div class="input-group">
        <label for="email">Email</label>
        <input type="email" id="email" v-model="email" required />
      </div>

      <div class="input-group">
        <label for="password">Password</label>
        <input type="password" id="password" v-model="password" required />
      </div>

      <div class="actions">
        <button type="submit">Login</button>
      </div>

      <div v-if="errorMessage" class="error">
        {{ errorMessage }}
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { MyApi } from '@/services/requests';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const email = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()

const submitLogin = async () => {

  const myApi = new MyApi(import.meta.env.VITE_BACKEND_URL)

  try {
    const res = await myApi.login({
      loginRequestSchema: {
        email: email.value,
        password: password.value
      }
    })

    errorMessage.value = ''

    sessionStorage.setItem('session_token', res.sessionToken);
    router.push('/dashboard')

  } catch (error: any) {

    if (error.response) {
      const errorBody = await error.response.json() // Parse the error response body as JSON
      errorMessage.value = `Error: ${errorBody.message || 'Something went wrong'}`
      return
    }

    errorMessage.value = "Network error or unexpected error occurred"
    return
  }

}
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: auto;
  padding: 2rem;
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

button {
  padding: 0.5rem 1rem;
  background-color: #27ae60;
  border: none;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #2ecc71;
}

.error {
  color: red;
  margin-top: 1rem;
  text-align: center;
}
</style>
