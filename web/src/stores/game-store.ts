import { CurrentServer } from "@/enums/game";
import type { GameState } from "@/types/game";
import { defineStore } from "pinia";

const initialGameState: GameState = {
  leftEvenPlayer: {
    id: "",
    name: ""
  },
  rightEvenPlayer: {
    id: "",
    name: ""
  },
  firstServer: CurrentServer.SERVER_LEFT_EVEN,
  progress: []
}

export const useGameStore = defineStore('game', {
  state: () => ({
    backendUrl: "",
    currentState: initialGameState
  }),
})