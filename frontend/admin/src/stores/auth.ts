import { defineStore } from 'pinia'
import axios from 'axios'

interface User {
  username: string
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
    token: sessionStorage.getItem('auth_token') || null,
    isValidating: false,
    isInitialized: false,
  }),

  getters: {
    isAuthenticated: (state) => !!state.user && !!state.token,
  },

  actions: {
    // Initialize - cek token saat app pertama load
    async initialize() {
      if (this.isInitialized) return

      if (this.token) {
        await this.validateAuth()
      }

      this.isInitialized = true
    },

    // Validasi token ke backend
    async validateAuth() {
      if (!this.token) {
        this.user = null
        return false
      }

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
          headers: { Authorization: `Bearer ${this.token}` },
        })

        this.user = response.data.data || { username: 'User' }
        return true
      } catch (error) {
        if (axios.isAxiosError(error)) {
          if (error.response?.status === 401) {
            console.warn('Token invalid, clearing auth')
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
    setAuth(token: string, user?: User) {
      this.token = token
      this.user = user || null
      sessionStorage.setItem('auth_token', token)
    },

    // Clear auth (logout)
    clearAuth() {
      this.user = null
      this.token = null
      sessionStorage.removeItem('auth_token')
    },

    // update profile
    async updateProfile(payload: { username?: string; password?: string }) {
      if (!this.token) {
        throw new Error('Unauthorized')
      }

      this.isValidating = true
      try {
        const baseUrl = import.meta.env.VITE_API_BASE_URL
        const response = await axios.patch(`${baseUrl}/api/admin/me`, payload, {
          headers: { Authorization: `Bearer ${this.token}` },
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
