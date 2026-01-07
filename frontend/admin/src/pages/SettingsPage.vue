<script setup lang="ts">
import { useRouter } from 'vue-router'
import axios from 'axios'
import NavbarDashboard from '@/components/NavbarDashboard.vue'
import { useAuthStore } from '@/stores/auth'
import { useDarkMode } from '@/composables/useDarkMode'

const { setTheme, currentTheme } = useDarkMode()

const router = useRouter()
const authStore = useAuthStore()

const logout = async () => {
  try {
    const token = authStore.token
    if (token) {
      await axios.post(
        `${import.meta.env.VITE_API_BASE_URL}/api/auth/logout`,
        {},
        {
          headers: { Authorization: `Bearer ${token}` },
        },
      )

      authStore.clearAuth()
      router.push('/login')
    }
  } catch (error) {
    console.error('Logout backend gagal:', error)
  }
}
</script>

<template>
  <div class="layout-wrapper">
    <main class="dashboard-container">
      <h1 class="dashboard-title">Settings</h1>

      <button class="btn-logout" @click="logout">Logout</button>
      <div class="theme-options">
        <button @click="setTheme('light')" :class="{ 'is-active': currentTheme === 'light' }">☀️</button>
        <button @click="setTheme('dark')" :class="{ 'is-active': currentTheme === 'dark' }">🌙</button>
        <button @click="setTheme('system')" :class="{ 'is-active': currentTheme === 'system' }">💻</button>
      </div>

      <div class="content-placeholder"></div>
    </main>

    <NavbarDashboard />
  </div>
</template>

<style lang="scss" scoped>
.layout-wrapper {
  color: var(--text-color);
  display: flex;
  flex-direction: column;
}

.dashboard-container {
  flex: 1;
  padding: 2rem;
  padding-bottom: 100px;
  text-align: center;
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;

  .dashboard-title {
    font-size: 4rem;
    font-weight: bold;
    padding-bottom: 1rem;
    border-bottom: 1px solid var(--primary-color);
  }
}

.btn-logout {
  margin-top: 2rem;
  padding: 0.6rem 1.2rem;
  background: #ff4d4f;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: opacity 0.3s;

  &:hover {
    opacity: 0.8;
  }
}

.theme-options {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  gap: 10px;

  button {
    background: rgba(var(--text-color), 0.1);
    border: 2px solid transparent;
    padding: 8px;
    border-radius: 10px;
    cursor: pointer;
    font-size: 18px;

    &.is-active {
      border-color: var(--primary-color);
    }
  }
}

.content-placeholder {
  margin-top: 2rem;
}
</style>
