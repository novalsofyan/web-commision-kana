<script setup lang="ts">
import { ref } from 'vue'
import NavbarDashboard from '@/components/NavbarDashboard.vue'
import ModalAddProduct from '@/components/ModalAddProduct.vue'
import { FilePlus } from 'lucide-vue-next'

const isModalOpen = ref(false)

const handleRefresh = () => {
  console.log('Data perlu di-refresh ke API...')
}
</script>

<template>
  <div class="layout-wrapper">
    <main class="dashboard-container">
      <h1>Dashboard Utama</h1>
      <p style="padding-bottom: 1rem">Selamat datang!</p>

      <div class="products-dashboard">
        <div class="product-title">
          <p>Manajemen Produk</p>
          <FilePlus class="plus-product" @click="isModalOpen = true" />
        </div>
      </div>
    </main>

    <NavbarDashboard />

    <ModalAddProduct :is-open="isModalOpen" @close="isModalOpen = false" @refresh="handleRefresh" />
  </div>
</template>

<style lang="scss" scoped>
.layout-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  max-width: 1200px;
  margin: 0 auto;
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
}

.product-title {
  font-size: 1.6rem;
  font-weight: bold;
  display: flex;
  justify-content: space-between;

  .plus-product {
    color: var(--primary-color);
    cursor: pointer;
    &:hover {
      transform: scale(1.1);
    }
  }
}
</style>
