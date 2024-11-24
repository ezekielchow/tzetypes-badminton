<template>
  <div class="login-root">
    <div class="login-container">
      <h2 class="title mb headlines">Login</h2>
      <form @submit.prevent="submitLogin">
        <div class="input-group">
          <label for="email">Email</label>
          <input type="email" id="email" v-model="email" name="email" autocomplete="email" required />
        </div>

        <div class="input-group">
          <label for="password">Password</label>
          <input type="password" id="password" v-model="password" name="password" autocomplete="password" required />
        </div>

        <div class="actions">
          <button class="button button-primary" type="submit">Login</button>
        </div>

        <div class="actions mt">New here? Join us now as a player. <RouterLink to="/signup-player">Signup today!
          </RouterLink>
        </div>

        <div v-if="errorMessage" class="error">
          {{ errorMessage }}
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useSessionStore } from '@/stores/session-store';
import { useUserStore } from '@/stores/user-store';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const email = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()

const sessionStore = useSessionStore()

const submitLogin = async () => {

  const userStore = useUserStore()

  userStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)
  const res = await userStore.login(email.value, password.value)

  if (res instanceof Error) {
    errorMessage.value = res.message
    return
  }

  errorMessage.value = ''
  sessionStorage.setItem('session_token', res.sessionToken);

  if (sessionStore.toRedirectToUrl && sessionStore.toRedirectToUrl !== "") {
    const to = sessionStore.toRedirectToUrl
    sessionStore.toRedirectToUrl = ""
    window.location.href = to
    return
  }

  router.push('/dashboard')
}
</script>

<style scoped>
.login-container {
  min-width: 380px;
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

.mt {
  margin-top: 10px;
}

.mb {
  margin-bottom: 10px;
}

.title {
  text-align: center;
}

.login-root {
  display: flex;
  min-height: 100vh;
  justify-content: center;
  align-items: center;
}
</style>
