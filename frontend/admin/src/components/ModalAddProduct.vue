<script setup lang="ts">
import { ref } from 'vue'
import { useProductStore } from '@/stores/product'

defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits(['close', 'created'])
const productStore = useProductStore()

const productName = ref('')
const productPrice = ref<number | null>(null)

const closeModal = () => {
  productName.value = ''
  productPrice.value = null
  emit('close')
}

const addProduct = async () => {
  if (!productName.value || productPrice.value === null) return

  await productStore.createProduct({
    product_name: productName.value,
    product_price: productPrice.value,
  })

  emit('created')
  closeModal()
}
</script>

<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="isOpen" class="modal-overlay" @click.self="closeModal">
        <div class="modal-content">
          <div class="modal-header">
            <h2>Tambah Produk Baru</h2>
            <button class="close-btn" @click="closeModal">&times;</button>
          </div>

          <div class="modal-body">
            <div class="input-group">
              <label>Nama Produk</label>
              <input v-model="productName" type="text" placeholder="Masukkan nama produk..." />
            </div>
            <div class="input-group">
              <label>Harga (IDR)</label>
              <input v-model="productPrice" type="number" placeholder="Contoh: 50000" />
            </div>
          </div>

          <div class="modal-footer">
            <button class="btn-secondary" @click="closeModal">Batal</button>
            <button class="btn-primary" @click="addProduct" :disabled="!productName || !productPrice">
              Simpan Produk
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style lang="scss">
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.modal-content {
  background: var(--bg-color);
  width: 90%;
  max-width: 450px;
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;

  h2 {
    font-size: 1.25rem;
    font-weight: 700;
    color: var(--primary-color);
  }

  .close-btn {
    background: none;
    border: none;
    font-size: 3rem;
    cursor: pointer;
    color: rgb(255, 100, 100);
  }
}

.modal-body .input-group {
  margin-bottom: 1.2rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;

  label {
    font-size: 1.4rem;
    font-weight: 600;
    opacity: 0.8;
  }

  input {
    padding: 0.8rem;
    border-radius: 8px;
    border: 1px solid #ddd;
    background: #f9f9f9;
    font-family: inherit;
    &:focus {
      outline: 2px solid var(--primary-color);
      border-color: transparent;
    }
  }
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1rem;

  button {
    padding: 0.7rem 1.2rem;
    border-radius: 8px;
    cursor: pointer;
    transition: 0.3s;
  }

  .btn-secondary {
    background: #eee;
    border: none;
    color: #666;
    &:hover {
      background: #ddd;
    }
  }

  .btn-primary {
    background: var(--primary-color);
    border: none;
    color: white;
    &:disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }
    &:hover:not(:disabled) {
      filter: brightness(1.1);
    }
  }
}

// Transition
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.1s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
