
import type { User } from '@/repositories/clients/private'
import { MyPrivateApi } from '@/services/requests-private'
import type { User as AuthUser } from 'firebase/auth'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    backendUrl: "",
    currentUser: null as User | null,
    firebaseUser: null as AuthUser | null,
    firebaseIdToken: ""
  }),
  actions: {
    setBackendUrl(backendUrl: string) {
      this.backendUrl = backendUrl
    },
    async getCurrentUser(): Promise<User | Error> {
      if (this.currentUser != null) {
        return this.currentUser
      }

      const privateApi = new MyPrivateApi(this.backendUrl)

      try {
        const res = await privateApi.currentUser()
        if (res?.user) {
          this.currentUser = res.user
          return this.currentUser
        }
        throw new Error("failed request")
      } catch (error: any) {
        if (error.response) {
          const errorBody = await error.response.json() // Parse the error response body as JSON
          return new Error(`Error: ${errorBody.message || 'Something went wrong'}`)
        }
        return new Error("Network error or unexpected error occurred")
      }
    }
  },
  persist: true
})
