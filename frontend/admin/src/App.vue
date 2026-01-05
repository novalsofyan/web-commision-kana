<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useDarkMode } from '@/composables/useDarkMode'
import { onMounted, onUnmounted } from 'vue'
import { useAuthStore } from '@/stores/auth'

const { initTheme } = useDarkMode()
const authStore = useAuthStore()
let refreshInterval: ReturnType<typeof setInterval> | null = null

onMounted(() => {
  initTheme()
  // Cek auth tiap 5 menit (300.000 ms)
  refreshInterval = setInterval(
    () => {
      if (authStore.token) {
        authStore.silentRefresh()
      }
    },
    5 * 60 * 1000,
  )
})

onUnmounted(() => {
  // Biar gak memory leak
  if (refreshInterval) clearInterval(refreshInterval)
})
</script>

<template>
  <RouterView />
</template>

<style lang="scss">
body {
  background-color: var(--bg-color);
  color: var(--text-color);
  transition:
    background-color 0.3s ease,
    color 0.3s ease;
  margin: 0;
}

#app {
  color: var(--text-color);
}

h1,
h2,
h3,
p,
span,
div {
  color: inherit;
}
</style>
