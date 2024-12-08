<script setup lang="ts">
import { getAuth, getIdToken, onAuthStateChanged } from "firebase/auth";
import { onBeforeUnmount } from 'vue';
import { RouterView } from 'vue-router';
import { useGameStore } from './stores/game-store';
import { useUserStore } from './stores/user-store';

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)
const userStore = useUserStore()

let gameProgressInterval: number
let isSyncing = false

gameStore.$subscribe((mutation, state) => {
  if (!gameProgressInterval && state.currentGameSettings.isEnded == false) {
    gameProgressInterval = window.setInterval(syncGameProgress, 500)
  }

  if (gameProgressInterval && state.currentGameSettings.isEnded) {
    window.clearInterval(gameProgressInterval)
  }
})

onBeforeUnmount(() => {
  if (gameProgressInterval > 0) {
    window.clearInterval(gameProgressInterval)
  }
})

const syncGameProgress = async () => {
  if (!isSyncing) {
    isSyncing = true
    await syncAddPoints()
    await syncRemovePoints()
    isSyncing = false
  }
}

const syncAddPoints = async () => {
  const toSync = gameStore.currentGameProgress.filter((step) => {
    return !step.isSynced
  })

  if (toSync.length < 1) {
    return
  }

  const res = await gameStore.addGameSteps({
    gameId: gameStore.currentGameSettings.id,
    addGameStepsRequestSchema: {
      steps: toSync
    }
  })

  if (res instanceof Error) {
    return
  }
}

const syncRemovePoints = async () => {
  if (gameStore.stepsToRemove.length < 1) {
    return
  }

  const ids = gameStore.stepsToRemove

  const res = await gameStore.deleteGameSteps({
    gameId: gameStore.currentGameSettings.id,
    requestBody: ids
  })

  if (res instanceof Error) {
    return
  }

  gameStore.stepsToRemove = gameStore.stepsToRemove.filter((step) => {
    return !ids.includes(step)
  })
}

const auth = getAuth();
onAuthStateChanged(auth, async (user) => {
  if (user) {
    const token = await getIdToken(user)
    userStore.firebaseIdToken = token
    userStore.firebaseUser = user
  } else {
    console.log("No user is signed in.");
  }
});

</script>

<template>
  <Suspense>
    <RouterView />
  </Suspense>
</template>

<style scoped></style>
