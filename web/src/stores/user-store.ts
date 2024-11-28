
import type { User } from '@/repositories/clients/private'
import type { LoginResponseSchema } from '@/repositories/clients/public'
import { MyPrivateApi } from '@/services/requests-private'
import { MyPublicApi } from '@/services/requests-public'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    backendUrl: "",
    currentUser: null as User | null
  }),
  actions: {
    setBackendUrl(backendUrl: string) {
      this.backendUrl = backendUrl
    },
    async login(email: string, password: string): Promise<LoginResponseSchema | Error> {
      const publicApi = new MyPublicApi(this.backendUrl)

      try {
        const res = await publicApi.login({
          loginRequestSchema: {
            email: email,
            password: password
          }
        })
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async signupPlayer(email: string, password: string, passwordRepeat: string): Promise<void | Error> {
      const publicApi = new MyPublicApi(this.backendUrl)

      try {
        const res = await publicApi.signupPlayer({
          signupRequestSchema: {
            email: email,
            password: password,
            passwordRepeat: passwordRepeat
          }
        })
        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async logout(): Promise<void | Error> {
      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.logoutRequest()
        privateApi.deleteSession()

        return res

      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    },
    async getCurrentUser(): Promise<User | Error> {
      if (this.currentUser != null) {
        return this.currentUser
      }

      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.currentUser()
        this.currentUser = res.user

        return this.currentUser

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
