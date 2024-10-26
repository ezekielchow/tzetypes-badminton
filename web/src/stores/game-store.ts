import { type AddGameSteps201Response, type AddGameStepsRequest, type DeleteGameStepsRequest, type Game, type StartGame201Response, type StartGameRequest } from "@/repositories/clients/private";
import { MyApi } from "@/services/requests";
import type { LocalGameStep } from "@/types/game";
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

const initialGameSteps: LocalGameStep[] = []

export const useGameStore = defineStore('game', {
  state: () => ({
    backendUrl: "",
    currentGameSettings: initialGameState,
    currentGameProgress: initialGameSteps,
    isMatchActive: false,
    stepsToRemove: [] as string[]
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

        this.isMatchActive = true
        this.currentGameSettings = res.game
        for (let i = 0; i < res.steps.length; i++) {
          this.currentGameProgress = this.currentGameProgress.concat({
            ...res.steps[i],
            isSynced: true
          })
        }

        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async addGameSteps(params: AddGameStepsRequest
    ): Promise<AddGameSteps201Response | Error> {
      const myApi = new MyApi(this.backendUrl)

      try {
        const res = await myApi.addGameSteps(params)

        for (let i = 0; i < res.gameSteps.length; i++) {
          const step = res.gameSteps[i];

          const matchedIndex = this.currentGameProgress.findIndex((progress) => progress.syncId == step.syncId)
          if (matchedIndex) {
            this.currentGameProgress[matchedIndex] = {
              ...this.currentGameProgress[matchedIndex],
              isSynced: true,
              id: step.id,
              createdAt: step.createdAt,
              updatedAt: step.updatedAt
            }
          }
        }

        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async deleteGameSteps(params: DeleteGameStepsRequest
    ): Promise<void | Error> {
      const myApi = new MyApi(this.backendUrl)

      try {
        const res = await myApi.deleteGameSteps(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
  },
  persist: true
})