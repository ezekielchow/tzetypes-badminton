import { defineStore } from "pinia"

export const useSessionStore = defineStore('session', {
  state: () => ({
    toRedirectToUrl: "",
  }),
  persist: true
})