
import * as runtime from '@/repositories/clients/private';
import { MyApi } from '@/services/requests';
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
      const myApi = new MyApi(this.backendUrl)

      try {
        const res = await myApi.listPlayers(params)
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async addPlayer(params: runtime.AddPlayerRequest): Promise<void | Error> {
      const myApi = new MyApi(this.backendUrl)

      try {
        const res = await myApi.addPlayer(params)
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
