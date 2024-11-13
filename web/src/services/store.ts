import { useGameStore } from "@/stores/game-store"
import { usePlayerStore } from "@/stores/player-store"
import { useUserStore } from "@/stores/user-store"

export const resetStores = () => {
  const gameStore = useGameStore()
  const playerStore = usePlayerStore()
  const userStore = useUserStore()

  gameStore.$reset()
  playerStore.$reset()
  userStore.$reset()
}