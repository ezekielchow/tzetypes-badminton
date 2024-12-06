import './assets/styles/main.css';

import firebase from "@/services/firebase";
import { createPinia } from 'pinia';
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import { createApp } from 'vue';

import App from './App.vue';
import router from './router';

import 'sweetalert2/dist/sweetalert2.min.css';


const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(firebase)
app.use(pinia)
app.use(router)

app.mount('#app')
