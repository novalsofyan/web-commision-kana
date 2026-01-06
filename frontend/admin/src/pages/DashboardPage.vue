<script setup lang="ts">
import { ref } from 'vue'
import NavbarDashboard from '@/components/NavbarDashboard.vue'
import ModalAddProduct from '@/components/ModalAddProduct.vue'
import ProductList from '@/components/ProductList.vue'
import { FilePlus } from 'lucide-vue-next'

const isModalOpen = ref(false)
</script>

<template>
  <div class="layout-wrapper">
    <main class="dashboard-container">
      <h1 class="dashboard-title">Dashboard Utama</h1>
      <p class="dashboard-subtitle">Selamat datang!</p>

      <div class="products-dashboard">
        <div class="product-title">
          <p>Manajemen Produk</p>
          <FilePlus class="plus-product" @click="isModalOpen = true" />
        </div>
        <ProductList />
      </div>
      <NavbarDashboard />
    </main>

    <ModalAddProduct :is-open="isModalOpen" @close="isModalOpen = false" />
  </div>
</template>

<style lang="scss" scoped>
.layout-wrapper {
  margin: 0 auto;
  color: var(--text-color);
  display: flex;
  flex-direction: column;
  max-width: 1200px;
}

.dashboard-container {
  padding: 2rem 2rem 10rem;

  .dashboard-title,
  .dashboard-subtitle {
    text-align: center;
  }

  .dashboard-subtitle {
    padding-bottom: 1rem;
  }
}

.products-dashboard {
  border-top: 0.1rem solid var(--primary-color);
  padding-top: 1rem;
}

.product-title {
  font-size: 1.6rem;
  font-weight: bold;
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;

  @media (min-width: 768px) {
    font-size: 2.4rem;
  }

  .plus-product {
    color: var(--primary-color);
    cursor: pointer;

    &:hover {
      transform: scale(1.1);
    }

    @media (min-width: 768px) {
      width: 3.4rem;
      height: 3.4rem;
    }
  }
}
</style>
