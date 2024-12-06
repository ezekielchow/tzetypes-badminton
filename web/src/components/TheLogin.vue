<template>
  <div class="login-root">
    <div class="login-container">
      <h2 class="title mb headlines">Login</h2>
      <form>
        <div class="input-group">
          <label for="email">Email</label>
          <input type="email" id="email" v-model="email" name="email" autocomplete="email" required />
        </div>

        <div class="input-group">
          <label for="password">Password</label>
          <input type="password" id="password" v-model="password" name="password" autocomplete="password" required />
        </div>

        <div class="actions">
          <ButtonComponent type="primary" :isLoading="isLoading" @click.prevent="submitLogin">
            Login
          </ButtonComponent>
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
import { auth } from '@/services/firebase';
import { signInWithEmailAndPassword } from 'firebase/auth';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import ButtonComponent from './ButtonComponent.vue';

const email = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()
const isLoading = ref(false);

const submitLogin = async () => {

  isLoading.value = true

  try {
    await signInWithEmailAndPassword(auth, email.value, password.value);
    router.push('/dashboard')
    isLoading.value = false
  } catch (err: any) {
    errorMessage.value = err.message
    isLoading.value = false
  }
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
