import Dashboard from '@/components/TheDashboard.vue';
import { MyApi } from '@/services/requests';
import GameSetupView from '@/views/GameSetupView.vue';
import LoginView from '@/views/LoginView.vue';
import OngoingGameView from '@/views/OngoingGameView.vue';
import PlayersAddView from '@/views/PlayersAddView.vue';
import PlayersDashboardView from '@/views/PlayersDashboardView.vue';
import PlayersEditView from '@/views/PlayersEditView.vue';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/login',
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: { onlyPublic: true }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: Dashboard,
      meta: { requiresAuth: true }
    },
    {
      path: '/players',
      name: 'players',
      component: PlayersDashboardView,
      meta: { requiresAuth: true }
    },
    {
      path: '/players/add',
      name: 'players/add',
      component: PlayersAddView,
      meta: { requiresAuth: true }
    },
    {
      path: '/players/:id',
      name: 'players/edit',
      component: PlayersEditView,
      meta: { requiresAuth: true }
    },
    {
      path: '/game/setup',
      name: 'game/setup',
      component: GameSetupView,
      meta: { requiresAuth: true }
    },
    {
      path: '/game/:id',
      name: 'game/playing',
      component: OngoingGameView,
      meta: { requiresAuth: true }
    },
  ]
})

router.beforeEach(async (to, from, next) => {

  let isAuthenticated = !!sessionStorage.getItem('session_token');

  // first time browsing
  if (!isAuthenticated) {
    const api = new MyApi(import.meta.env.VITE_BACKEND_URL)

    try {
      await api.refreshToken()
      isAuthenticated = !!sessionStorage.getItem('session_token');
    } catch (error) {
      console.log(error);
    }
  }

  const metas = Object.keys(to.meta)
  if (metas.includes("onlyPublic")) {
    if (!isAuthenticated) {
      next();
      return
    } else {
      next({ name: 'dashboard' });
      return
    }
  } else if (metas.includes("requiresAuth")) {
    if (isAuthenticated) {
      next();
      return
    } else {
      next({ name: 'login' });
      return
    }
  }

  next()
});

export default router
