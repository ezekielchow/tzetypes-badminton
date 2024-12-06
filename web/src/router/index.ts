import { useUserStore } from '@/stores/user-store';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/login',
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: { onlyPublic: true }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/components/TheDashboard.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/players',
      name: 'players',
      component: () => import('@/views/PlayersDashboardView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/players/add',
      name: 'players/add',
      component: () => import('@/views/PlayersAddView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/players/:id',
      name: 'players/edit',
      component: () => import('@/views/PlayersEditView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/game/setup',
      name: 'game/setup',
      component: () => import('@/views/GameSetupView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/game/:id',
      name: 'game/playing',
      component: () => import('@/views/OngoingGameView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/game/:id/statistics',
      name: 'game/statistics',
      component: () => import('@/views/GameStatisticsView.vue'),
    },
    {
      path: '/signup-player',
      name: 'signup-player',
      component: () => import('@/components/SignupPlayer.vue'),
      meta: { onlyPublic: true }
    },
  ]
})


router.beforeEach(async (to, from, next) => {

  const userStore = useUserStore()

  let isAuthenticated = null
  if (userStore.firebaseUser) {
    isAuthenticated = userStore.firebaseUser
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
