import { type Game, type GameStep, type StartGame201Response, type StartGameRequest } from "@/repositories/clients/private";
import { MyApi } from "@/services/requests";
import { defineStore } from "pinia";

const initialGameState: Game = {
  id: "",
  clubId: "",
  leftEvenPlayerName: "",
  leftOddPlayerName: "",
  rightEvenPlayerName: "",
  rightOddPlayerName: "",
  gameType: "",
  servingSide: "",
  createdAt: "",
  updatedAt: ""
}

const initialGameSteps: GameStep[] = []

export const useGameStore = defineStore('game', {
  state: () => ({
    backendUrl: "",
    currentGameSettings: initialGameState,
    currentGameProgress: initialGameSteps
  }),
  actions: {
    setBackendUrl(backendUrl: string) {
      this.backendUrl = backendUrl
    },
    async startGame(params: StartGameRequest
    ): Promise<StartGame201Response | Error> {
      const myApi = new MyApi(this.backendUrl)

      try {
        const res = await myApi.startGame(params)

        this.currentGameSettings = res.game
        this.currentGameProgress = res.steps

        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
  }
})