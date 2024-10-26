<script setup lang="ts">
import { onBeforeUnmount } from 'vue';
import { RouterView } from 'vue-router';
import { useGameStore } from './stores/game-store';

const gameStore = useGameStore()
gameStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

let gameProgressInterval: number

gameStore.$subscribe((mutation, state) => {
  if (!gameProgressInterval && state.isMatchActive) {
    gameProgressInterval = window.setInterval(syncGameProgress, 3000)
  }
})

onBeforeUnmount(() => {
  if (gameProgressInterval > 0) {
    window.clearInterval(gameProgressInterval)
  }
})

const syncGameProgress = async () => {
  await syncAddPoints()
  await syncRemovePoints()
}

const syncAddPoints = async () => {
  const toSync = gameStore.currentGameProgress.filter((step) => {
    return !step.isSynced
  })

  if (toSync.length < 1) {
    return
  }

  const res = await gameStore.addGameSteps({
    id: gameStore.currentGameSettings.id,
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

</script>

<template>
  <Suspense>
    <RouterView />
  </Suspense>
</template>

<style scoped></style>
