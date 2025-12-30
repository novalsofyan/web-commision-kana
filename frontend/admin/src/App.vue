<script setup lang="ts">
import { RouterView } from 'vue-router'
import { useDarkMode } from '@/composables/useDarkMode'
import { onMounted } from 'vue'

const { initTheme } = useDarkMode()

// Inisialisasi tema global di root aplikasi
onMounted(() => {
  initTheme()
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
}
</style>
