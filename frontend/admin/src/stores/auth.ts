import { defineStore } from 'pinia'
import axios from 'axios'

interface User {
  username: string
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
    isValidating: false,
    isInitialized: false,
  }),

  getters: {
    isAuthenticated: (state) => !!state.user,
  },

  actions: {
    // Initialize - cek auth saat app pertama load
    async initialize() {
      if (this.isInitialized) return

      await this.validateAuth()
      this.isInitialized = true
    },

    // Validasi auth ke backend (cookie-based)
    async validateAuth() {
      if (this.user) return true

      if (this.isValidating) {
        return new Promise((resolve) => {
          const interval = setInterval(() => {
            if (!this.isValidating) {
              clearInterval(interval)
              resolve(!!this.user)
            }
          }, 50)
        })
      }

      this.isValidating = true
      try {
        const baseUrl = import.meta.env.VITE_API_BASE_URL
        const response = await axios.get(`${baseUrl}/api/auth/me`, {
          withCredentials: true,
        })

        this.user = response.data.data || { username: 'User' }
        return true
      } catch (error) {
        if (axios.isAxiosError(error)) {
          if (error.response?.status === 401) {
            console.warn('Session invalid, clearing auth')
            this.clearAuth()
            return false
          }

          console.error('error:', error.message)
          this.user = { username: 'User (Offline)' }
          return true
        }

        this.clearAuth()
        return false
      } finally {
        this.isValidating = false
      }
    },

    // Set auth setelah login
    setAuth(user: User) {
      this.user = user
    },

    // Clear auth (logout)
    clearAuth() {
      this.user = null
    },

    // update profile
    async updateProfile(payload: { username?: string; password?: string }) {
      this.isValidating = true
      try {
        const baseUrl = import.meta.env.VITE_API_BASE_URL
        const response = await axios.patch(`${baseUrl}/api/admin/me`, payload, {
          withCredentials: true,
        })

        if (payload.username && this.user) {
          this.user = { ...this.user, username: payload.username }
        }

        return { success: true, data: response.data.data }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          throw new Error(error.response?.data?.message || 'Update gagal')
        }
        throw error
      } finally {
        this.isValidating = false
      }
    },
  },
})
