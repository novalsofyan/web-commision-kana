<script setup lang="ts">
import NavbarDashboard from '@/components/NavbarDashboard.vue'
</script>

<template>
  <div class="layout-wrapper">
    <main class="dashboard-container">
      <h1>Dashboard Utama</h1>
      <p>Selamat datang!</p>

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

.content-placeholder {
  margin-top: 2rem;
}
</style>
