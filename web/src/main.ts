import './assets/main.css';

import { createPinia } from 'pinia';
import { createApp } from 'vue';

import { VueGoodTablePlugin } from 'vue-good-table';
import 'vue-good-table/dist/vue-good-table.css';

import App from './App.vue';
import router from './router';

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(VueGoodTablePlugin)

app.mount('#app')
