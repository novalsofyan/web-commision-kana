<script setup lang="ts">
import { useRouter } from 'vue-router'
import axios from 'axios'
import NavbarDashboard from '@/components/NavbarDashboard.vue'

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

<template>
  <div class="layout-wrapper">
    <main class="dashboard-container">
      <h1>⚙️ Settings</h1>

      <button class="btn-logout" @click="logout">Logout</button>

      <div class="content-placeholder"></div>
    </main>

    <NavbarDashboard />
  </div>
</template>

<style lang="scss" scoped>
.layout-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color, #f9fafb);
}

.dashboard-container {
  flex: 1;
  padding: 2rem;
  padding-bottom: 100px;
  text-align: center;
}

.btn-logout {
  margin-top: 1rem;
  padding: 0.6rem 1.2rem;
  background-color: #ff4d4f;
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

.content-placeholder {
  margin-top: 2rem;
}
</style>
