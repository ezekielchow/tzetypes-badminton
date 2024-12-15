<template>
  <div class="wrapper">
    <div class="login-root">
      <div class="login-container">
        <h2 class="title mb headlines">Login</h2>
        <div>
          <div class="input-group">
            <label for="email">Email</label>
            <input type="email" id="email" v-model="email" name="email" autocomplete="email" required />
          </div>

          <div class="input-group">
            <label for="password">Password</label>
            <input type="password" id="password" v-model="password" name="password" autocomplete="password" required />
          </div>

          <div class="actions">
            <ButtonComponent type="primary" :isLoading="isLoading" @click="submitLogin">
              Login
            </ButtonComponent>
          </div>

          <div class="actions mt">New here? Join us now as a player. <RouterLink to="/signup-player">Signup today!
            </RouterLink>
          </div>

          <div v-if="errorMessage" class="error">
            {{ errorMessage }}
          </div>
        </div>
      </div>
    </div>
    <div class="seperator"></div>
    <div class="ig-wrapper">
      <h2 class="title mb headlines mt">Latest Posts</h2>
      <swiper-container :loop="true" :slides-per-view="1" :space-between="10" :centered-slides="true" :pagination="true"
        :navigation="true">
        <swiper-slide v-for="media in medias" :key="media.id">
          <div class="slider-wrapper">
            <a :href="media.permalink">
              <img :src="media.mediaUrl" class="slider-img" alt="">
            </a>
          </div>
        </swiper-slide>
      </swiper-container>
    </div>
  </div>
</template>

<script setup lang="ts">
import { instanceOfGetInstagramFeed200Response, type InstagramMedia } from '@/repositories/clients/public';
import { useGameStore } from '@/stores/game-store';
import { useUserStore } from '@/stores/user-store';
import { signInWithEmailAndPassword, type Auth } from 'firebase/auth';
import { register } from 'swiper/element/bundle';
import { inject, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import ButtonComponent from './ButtonComponent.vue';

register();

const email = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()
const isLoading = ref(false);
const medias = ref([] as InstagramMedia[])
const auth = inject<Auth>("auth");

const userStore = useUserStore()
userStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

onMounted(async () => {
  await getInstagramFeed()
})

const getInstagramFeed = async () => {
  try {
    const data = await gameStore.getInstagramFeed()
    if (instanceOfGetInstagramFeed200Response(data)) {
      medias.value = data.feed
      console.log(medias.value);

    }
  } catch (error: any) {
    errorMessage.value = "unable to get insta feed"
  }
}

const submitLogin = async () => {

  isLoading.value = true

  try {
    if (auth) {
      await signInWithEmailAndPassword(auth, email.value, password.value);
      if (auth) {
        userStore.firebaseUser = auth.currentUser
      }

      router.push('/dashboard')
      isLoading.value = false
    } else {
      errorMessage.value = "no auth to sign in"
      isLoading.value = false
    }
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
  min-height: 40vh;
  justify-content: center;
  align-items: center;
}

.wrapper {
  display: flex;
  flex-direction: column;
}

.ig-wrapper {
  min-height: 55vh;
}

.slider-img {
  max-height: 50vh;
}

.slider-wrapper {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.seperator {
  width: 100%;
  height: 0.2rem;
  background-color: silver;
}

@media (min-width: 768px) {
  .wrapper {
    display: flex;
    flex-direction: row;
  }

  .ig-wrapper {
    min-width: 58vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .login-root {
    min-width: 40vw;
    height: 100vh;
  }

  .seperator {
    height: 100vh;
    width: 0.5rem;
    background-color: silver;
  }
}
</style>
