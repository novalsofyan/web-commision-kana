<template>
  <div style="padding: 2rem; text-align: center">
    <h1>🚀 Dashboard Utama</h1>
    <button @click="logout" style="padding: 0.5rem 1rem; cursor: pointer">
      Logout
    </button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

const logout = async () => {
  try {
    const token = sessionStorage.getItem('auth_token')
    if (token) {
      await axios.post(
        `${import.meta.env.VITE_API_BASE_URL}/api/auth/logout`,
        {},
        {
          headers: { Authorization: `Bearer ${token}` },
        },
      )
    }
  } catch (error) {
    console.error('Logout backend gagal:', error)
  } finally {
    sessionStorage.removeItem('auth_token')
    router.push('/login')
  }
}
</script>
