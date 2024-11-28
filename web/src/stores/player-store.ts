
import * as runtime from '@/repositories/clients/private';
import { MyPrivateApi } from '@/services/requests-private';
import { defineStore } from 'pinia';

export const usePlayerStore = defineStore('player', {
  state: () => ({
    backendUrl: "",
  }),
  actions: {
    setBackendUrl(backendUrl: string) {
      this.backendUrl = backendUrl
    },
    async listPlayers(params: runtime.ListPlayersRequest
    ): Promise<runtime.ListPlayers200Response | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.listPlayers(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async addPlayer(params: runtime.AddPlayerRequest): Promise<runtime.Player | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.addPlayer(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async getPlayer(params: runtime.GetPlayerWithIdRequest): Promise<runtime.Player | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.getPlayer(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async updatePlayer(params: runtime.UpdatePlayerWithIdOperationRequest): Promise<runtime.Player | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.updatePlayer(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    }
  }
})
