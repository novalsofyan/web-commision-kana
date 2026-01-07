<script setup lang="ts">
import NavbarDashboard from '@/components/NavbarDashboard.vue'
import { ref, computed } from 'vue'
import { useAuthStore } from '@/stores/auth'
import axios from 'axios'

const username = ref<string>('')
const password = ref<string>('')
const authStore = useAuthStore()
const isLoading = ref<boolean>(false)
const successMessage = ref<string>('')
const errorMessage = ref<string>('')

const isDisabled = computed(() => {
  const hasUsername = username.value.trim().length > 0
  const hasPassword = password.value.trim().length > 0

  return isLoading.value || (!hasUsername && !hasPassword)
})

const handleUpdate = async () => {
  successMessage.value = ''
  errorMessage.value = ''
  isLoading.value = true

  try {
    const payload: { username?: string; password?: string } = {}

    if (username.value.trim()) {
      payload.username = username.value.trim()
    }

    if (password.value.trim()) {
      payload.password = password.value.trim()
    }

    await authStore.updateProfile(payload)

    successMessage.value = 'Profil berhasil diupdate!'

    username.value = ''
    password.value = ''
  } catch (error) {
    if (axios.isAxiosError(error)) {
      errorMessage.value = error.response?.data?.message || 'Gagal update profil. Coba lagi ya!'
    } else if (error instanceof Error) {
      errorMessage.value = error.message
    } else {
      errorMessage.value = 'Terjadi kesalahan yang tidak diketahui'
    }
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <main class="profile-container">
    <h1 class="profile-title">Profile</h1>
    <div class="profile-content">
      <div class="profile-card">
        <h2 class="card-title">Ubah Profil</h2>

        <div v-if="successMessage" class="alert alert-success">
          {{ successMessage }}
        </div>
        <div v-if="errorMessage" class="alert alert-error">
          {{ errorMessage }}
        </div>

        <form class="form-update-profile" @submit.prevent="handleUpdate">
          <div class="form-group">
            <label for="name">Username <span class="optional">(opsional)</span></label>
            <input
              v-model="username"
              type="text"
              id="name"
              placeholder="Kosongkan jika tidak ingin mengubah"
              :disabled="isLoading"
            />
          </div>

          <div class="form-group">
            <label for="password">Password <span class="optional">(opsional)</span></label>
            <input
              v-model="password"
              type="password"
              id="password"
              placeholder="Kosongkan jika tidak ingin mengubah"
              :disabled="isLoading"
            />
          </div>

          <button type="submit" class="btn-update" :disabled="isDisabled">
            {{ isLoading ? 'Menyimpan...' : 'Simpan Perubahan' }}
          </button>
        </form>
      </div>
    </div>
    <NavbarDashboard />
  </main>
</template>

<style lang="scss">
.profile-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.profile-title {
  color: var(--text-color);
  text-align: center;
  font-size: 4rem;
  font-weight: bold;
  padding-bottom: 1rem;
  border-bottom: 1px solid var(--primary-color);
}

.profile-content {
  margin-top: 2rem;
  display: flex;
  justify-content: center;
}

.profile-card {
  width: 100%;
  border: 1px solid var(--primary-color);
  border-radius: 1rem;
  padding: 2rem;
  box-sizing: border-box;

  .card-title {
    color: var(--text-color);
    font-size: 1.6rem;
    text-transform: uppercase;
    font-weight: 800;
    margin-bottom: 1rem;
    text-align: center;
  }
}

.form-update-profile {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0 auto;
  width: 100%;
}

.form-group {
  color: var(--text-color);
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
  width: 100%;
}

.form-group label {
  font-size: 1.6rem;
  font-weight: 800;
  text-transform: uppercase;
  margin-bottom: 0.5rem;

  .optional {
    font-size: 1.2rem;
    font-weight: normal;
    text-transform: none;
  }
}

.form-group input {
  padding: 0.5rem;
  border: 1px solid var(--primary-color);
  border-radius: 0.5rem;
  font-size: 1.6rem;
  margin-bottom: 1rem;
  background: white;

  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
}

.btn-update {
  background-color: var(--primary-color);
  color: whitesmoke;
  width: 100%;
  padding: 0.7rem;
  border: none;
  border-radius: 1rem;
  font-size: 1.4rem;
  font-weight: 700;
  cursor: pointer;
  margin-top: 1rem;
  transition: all 0.3s ease;

  &:hover:not(:disabled) {
    opacity: 0.8;
    transform: translateY(-2px);
  }

  &:disabled {
    background-color: #cccccc;
    color: #666666;
    cursor: not-allowed;
    border: 1px solid #999999;
  }
}

/* DESKTOP SIZING SYNC */
@media (min-width: 768px) {
  .form-group label {
    font-size: 1.6rem;
  }

  .form-group input {
    font-size: 1.5rem;
  }

  .btn-update {
    width: auto;
    padding: 0.6rem 2rem;
    font-size: 1.5rem;
  }
}

// Alert styles
.alert {
  padding: 1rem;
  border-radius: 0.5rem;
  margin-bottom: 1rem;
  text-align: center;
  font-size: 1.4rem;
  font-weight: 500;
}

.alert-success {
  background-color: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.alert-error {
  background-color: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}
</style>
