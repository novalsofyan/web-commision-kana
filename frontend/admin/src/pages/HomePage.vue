<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios, { AxiosError } from 'axios'
import { useDarkMode } from '@/composables/useDarkMode'

interface LoginResponse {
  data: {
    token: string
    username: string
  }
}

const { setTheme, currentTheme } = useDarkMode()
const router = useRouter()

const username = ref<string>('')
const password = ref<string>('')
const isLoading = ref<boolean>(false)
const errorMessage = ref<string>('')

const handleLogin = async (): Promise<void> => {
  if (!username.value || !password.value) {
    errorMessage.value = 'Isi dulu username & password-nya ya!'
    return
  }

  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await axios.post<LoginResponse>(
      `${import.meta.env.VITE_API_BASE_URL}/api/auth/login`,
      {
        username: username.value,
        password: password.value,
      },
    )

    const token = response.data.data.token
    sessionStorage.setItem('auth_token', token)

    // Redirect (misal ke dashboard)
    router.push('/dashboard')
  } catch (err) {
    const error = err as AxiosError<{ message?: string }>
    errorMessage.value =
      error.response?.data?.message || 'Gagal login, cek koneksi atau akun!'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <main>
    <div class="login-card">
      <h2>Welcome Back!</h2>
      <p>Silahkan login dulu ya...</p>

      <div class="form-group">
        <input
          v-model="username"
          type="text"
          placeholder="Username"
          @keyup.enter="handleLogin"
        />
        <input
          v-model="password"
          type="password"
          placeholder="Password"
          @keyup.enter="handleLogin"
        />
      </div>

      <p
        v-if="errorMessage"
        style="color: #ff4d4d; font-size: 0.8rem; margin-bottom: 10px"
      >
        {{ errorMessage }}
      </p>

      <button class="login-btn" @click="handleLogin" :disabled="isLoading">
        {{ isLoading ? 'Loading...' : 'Login' }}
      </button>

      <div class="theme-options">
        <button
          @click="setTheme('light')"
          :class="{ active: currentTheme === 'light' }"
          title="Light Mode"
        >
          ☀️
        </button>
        <button
          @click="setTheme('dark')"
          :class="{ active: currentTheme === 'dark' }"
          title="Dark Mode"
        >
          🌙
        </button>
        <button
          @click="setTheme('system')"
          :class="{ active: currentTheme === 'system' }"
          title="System Mode"
        >
          💻
        </button>
      </div>
    </div>
  </main>
</template>

<style lang="scss" scoped>
main {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--bg-color);
  color: var(--text-color);
  transition: all 0.3s ease;
}

.login-card {
  background: rgba(var(--text-color), 0.05);
  padding: 2.5rem;
  border-radius: 16px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  text-align: center;
  border: 1px solid var(--primary-color);
}

.theme-options {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  gap: 15px;

  button {
    background: transparent;
    border: 1px solid transparent;
    font-size: 1.2rem;
    cursor: pointer;
    padding: 5px;
    border-radius: 8px;
    transition: all 0.2s;
    opacity: 0.5;

    &.active {
      opacity: 1;
      border-color: var(--primary-color);
      background: rgba(var(--primary-color), 0.1);
    }

    &:hover {
      opacity: 1;
      transform: scale(1.1);
    }
  }
}

.form-group {
  margin: 2rem 0;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  input {
    padding: 0.8rem;
    border-radius: 8px;
    border: 1px solid #ccc;
    background: transparent;
    color: var(--text-color);
  }
}

.login-btn {
  width: 100%;
  padding: 0.8rem;
  background-color: var(--primary-color);
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: bold;

  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
}
</style>
