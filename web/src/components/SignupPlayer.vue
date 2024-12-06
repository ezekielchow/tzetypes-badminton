<template>
  <div class="signup-root">
    <div class="signup-container">
      <h2 class="title mb headlines">Player Signup</h2>
      <div>
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
          <ButtonComponent type="primary" :isLoading="isLoading" @click="submitSignup">
            Start playing!
          </ButtonComponent>
        </div>

        <div class="actions mt">Already have an account? <RouterLink to="/login">Login here</RouterLink>
        </div>

        <div v-if="errorMessage" class="error">
          {{ errorMessage }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { createUserWithEmailAndPassword, type Auth } from 'firebase/auth';
import { inject, ref } from 'vue';
import { useRouter } from 'vue-router';
import ButtonComponent from './ButtonComponent.vue';

const email = ref('')
const password = ref('')
const passwordRepeat = ref("")
const errorMessage = ref('')
const isLoading = ref(false)

const router = useRouter()
const auth = inject<Auth>("auth");

const submitSignup = async () => {
  isLoading.value = true

  if (email.value.length < 1) {
    errorMessage.value = "email is not valid"
    isLoading.value = false
    return
  }

  if (password.value.length < 1 || passwordRepeat.value.length < 1) {
    errorMessage.value = "password are invalid"
    isLoading.value = false
    return
  }

  if (password.value !== passwordRepeat.value) {
    errorMessage.value = "passwords must match"
    isLoading.value = false
    return
  }

  try {
    if (auth) {
      await createUserWithEmailAndPassword(auth, email.value, password.value);
      errorMessage.value = ''
      router.push('/login')
      isLoading.value = false
    } else {
      errorMessage.value = 'no login detected'
      isLoading.value = false
    }
  } catch (error: any) {
    errorMessage.value = error.message
    isLoading.value = false
  }
}
</script>

<style scoped>
.signup-container {
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

.signup-root {
  display: flex;
  min-height: 100vh;
  justify-content: center;
  align-items: center;
}
</style>
