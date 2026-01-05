<script setup lang="ts">
import { ref } from 'vue'
import NavbarDashboard from '@/components/NavbarDashboard.vue'
import ModalAddProduct from '@/components/ModalAddProduct.vue'
import ProductList from '@/components/ProductList.vue'
import { FilePlus } from 'lucide-vue-next'
import { useProductStore } from '@/stores/product'

const isModalOpen = ref(false)
const productStore = useProductStore()

const handleCreated = async () => {
  await productStore.fetchProducts()
}
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
    </main>

    <NavbarDashboard />

    <ModalAddProduct :is-open="isModalOpen" @close="isModalOpen = false" @created="handleCreated" />
  </div>
</template>

<style lang="scss" scoped>
.layout-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  max-width: 1200px;
  padding-bottom: 8rem;
  margin: 0 auto;
}

.dashboard-container {
  flex: 1;
  padding: 2rem;

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
  margin-bottom: 1rem;
}

.product-title {
  font-size: 1.6rem;
  font-weight: bold;
  display: flex;
  justify-content: space-between;
  margin-bottom: 1rem;

  .plus-product {
    color: var(--primary-color);
    cursor: pointer;

    &:hover {
      transform: scale(1.1);
    }
  }
}
</style>
