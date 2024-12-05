import { type AddGameSteps201Response, type AddGameStepsRequest, type CreateOrUpdateGameHistoryRequest, type DeleteGameStepsRequest, type EndGameOperationRequest, type Game, type GetGameHistory200Response, type GetGameHistoryRequest, type GetRecentStatistics200Response, type StartGameRequest } from "@/repositories/clients/private";
import type { StartGame201Response } from "@/repositories/clients/private/models/StartGame201Response";
import type { GetGame200Response, GetGameRequest } from "@/repositories/clients/public";
import { MyPrivateApi } from "@/services/requests-private";
import { MyPublicApi } from "@/services/requests-public";
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
  isEnded: false,
  createdAt: "",
  updatedAt: ""
}

const initialGameSteps: LocalGameStep[] = []

export const useGameStore = defineStore('game', {
  state: () => ({
    backendUrl: "",
    currentGameSettings: initialGameState,
    currentGameProgress: initialGameSteps,
    stepsToRemove: [] as string[]
  }),
  actions: {
    setBackendUrl(backendUrl: string) {
      this.backendUrl = backendUrl
    },
    async startGame(params: StartGameRequest
    ): Promise<StartGame201Response | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.startGame(params)
        if (res) {
          this.currentGameSettings = res.game
          this.currentGameProgress = []
          for (let i = 0; i < res.steps.length; i++) {
            this.currentGameProgress = this.currentGameProgress.concat({
              ...res.steps[i],
              isSynced: true
            })
          }

          return res
        }
        throw new Error("request failed");


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
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.addGameSteps(params)
        if (res) {
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
        }
        throw new Error("request failed")

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
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.deleteGameSteps(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async endGame(params: EndGameOperationRequest
    ): Promise<void | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.endGame(params)
        this.currentGameSettings.isEnded = true
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async getGameStatistics(params: GetGameRequest
    ): Promise<GetGame200Response | Error> {
      const publicApi = new MyPublicApi(this.backendUrl)

      try {
        const res = await publicApi.getGame(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async createOrUpdateGameHistory(params: CreateOrUpdateGameHistoryRequest
    ): Promise<GetGameHistory200Response | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.createOrUpdateGameHistory(params)
        if (res) {
          return res
        }
        throw new Error("request failed")

      } catch (error: any) {
        if (error.response) {
          if (error.response.statusText) {
            return new Error(`Error: ${error.response.statusText || 'Something went wrong'}`)
          }

          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async getGameHistory(params: GetGameHistoryRequest
    ): Promise<GetGameHistory200Response | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.getGameHistory(params)
        if (res) {
          return res
        }
        throw new Error("request failed")

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async getRecentStatistics(): Promise<GetRecentStatistics200Response | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.getRecentStatistics()
        if (res) {
          return res
        }
        throw new Error("request failed")

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    loadGame(oldData: GetGame200Response) {
      this.stepsToRemove = []
      this.currentGameSettings = oldData.game

      const steps: LocalGameStep[] = []

      for (let i = 0; i < oldData.steps.length; i++) {
        const step = oldData.steps[i];

        steps.push({
          ...step,
          isSynced: true
        } as LocalGameStep)
      }

      this.currentGameProgress = steps
    }
  },
  persist: true
})