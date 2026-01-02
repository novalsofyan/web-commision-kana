<script setup lang="ts">
import NavbarDashboard from '@/components/NavbarDashboard.vue'
import { FilePlus } from 'lucide-vue-next';
</script>

<template>
  <div class="layout-wrapper">
    <main class="dashboard-container">
      <h1>Dashboard Utama</h1>
      <p style="padding-bottom: 1rem;">Selamat datang!</p>

      <div class="products-dashboard">
        <div class="product-title">
          <p>Manajemen Produk</p>
          <FilePlus class="plus-product"/>
        </div>
      </div>
    </main>

    <NavbarDashboard />
  </div>
</template>

<style lang="scss" scoped>
.layout-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  max-width: 1200px;
  margin: 0 auto
}

.dashboard-container {
  flex: 1;
  padding: 2rem;
  text-align: center;
}

.products-dashboard {
  border-top: 0.1rem solid var(--primary-color);
  text-align: left;
  color: var(--text-color);
  padding-top: 1rem;
  padding-bottom: 100px;
}

.product-title {
  font-size: 1.6rem;
  font-weight: bold;
  display: flex;
  justify-content: space-between;

  .plus-product {
    color: var(--primary-color);
    cursor: pointer;
  }
}
</style>
