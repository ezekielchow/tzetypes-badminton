import './assets/styles/main.css';

import { auth } from '@/services/firebase';
import { browserLocalPersistence, setPersistence } from "firebase/auth";
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import { createApp } from 'vue';

import App from './App.vue';
import router from './router';

import 'sweetalert2/dist/sweetalert2.min.css';

await setPersistence(auth, browserLocalPersistence)

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(router)

app.mount('#app')
