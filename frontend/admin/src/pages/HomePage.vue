<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios, { AxiosError } from 'axios'
import { useDarkMode } from '@/composables/useDarkMode'

// Interface buat response API
interface LoginResponse {
  data: {
    token: string
    username: string
  }
}

const { setTheme, currentTheme } = useDarkMode()
const router = useRouter()

// Reaktif variabel
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
    const response = await axios.post<LoginResponse>(`${import.meta.env.VITE_API_BASE_URL}/api/auth/login`, {
      username: username.value,
      password: password.value,
    })

    const token = response.data.data.token
    sessionStorage.setItem('auth_token', token)
    router.push('/dashboard')
  } catch (err) {
    const error = err as AxiosError<{ message?: string }>
    errorMessage.value = error.response?.data?.message || 'Gagal login, cek koneksi!'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <main class="login-page">
    <div class="login-card">
      <div class="card-header">
        <h2>Welcome Back!</h2>
        <p>Silahkan login dulu ya...</p>
      </div>

      <div class="form-group">
        <div class="input-wrapper">
          <label class="input-label">Username</label>
          <input
            v-model="username"
            type="text"
            placeholder="Username kamu"
            class="input-field"
            @keyup.enter="handleLogin"
          />
        </div>

        <div class="input-wrapper">
          <label class="input-label">Password</label>
          <input
            v-model="password"
            type="password"
            placeholder="Password kamu"
            class="input-field"
            @keyup.enter="handleLogin"
          />
        </div>
      </div>

      <p v-if="errorMessage" class="error-text-display">
        {{ errorMessage }}
      </p>

      <button class="login-btn" @click="handleLogin" :disabled="isLoading">
        {{ isLoading ? 'Loading...' : 'LOGIN' }}
      </button>

      <div class="theme-options">
        <button @click="setTheme('light')" :class="{ 'is-active': currentTheme === 'light' }">☀️</button>
        <button @click="setTheme('dark')" :class="{ 'is-active': currentTheme === 'dark' }">🌙</button>
        <button @click="setTheme('system')" :class="{ 'is-active': currentTheme === 'system' }">💻</button>
      </div>
    </div>
  </main>
</template>

<style lang="scss" scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--bg-color);
  padding: 1rem;
}

.login-card {
  background: var(--bg-color);
  padding: 40px;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 400px;
  border: 2px solid var(--primary-color);
  text-align: center;
}

.card-header {
  margin-bottom: 25px;
  h2 {
    font-size: 24px;
    color: var(--text-color);
    margin-bottom: 8px;
  }
  p {
    font-size: 14px;
    color: var(--text-color);
    opacity: 0.7;
  }
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-bottom: 20px;
  text-align: left;

  .input-label {
    font-size: 14px;
    font-weight: 600;
    margin-bottom: 5px;
    display: block;
  }

  .input-field {
    width: 100%;
    padding: 12px 15px;
    font-size: 16px;
    border-radius: 8px;
    border: 2px solid #ddd;
    background: #fff;
    color: #333;
    outline: none;

    &:focus {
      border-color: var(--primary-color);
    }
  }
}

.error-text-display {
  color: #ff4d4d;
  font-size: 14px;
  margin-bottom: 15px;
  font-weight: bold;
}

.login-btn {
  width: 100%;
  padding: 14px;
  background-color: var(--primary-color);
  color: white;
  border: none;
  border-radius: 10px;
  font-weight: bold;
  cursor: pointer;
  font-size: 16px;

  &:disabled {
    opacity: 0.5;
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
</style>
