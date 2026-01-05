<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { useProductStore } from '@/stores/product'
import { storeToRefs } from 'pinia'

export default defineComponent({
  setup() {
    const productStore = useProductStore()
    const { products, loading } = storeToRefs(productStore)

    onMounted(() => {
      if (products.value.length === 0) {
        productStore.fetchProducts()
      }
    })

    const handleEdit = (id: number) => {
      console.log('Edit product', id)
    }

    const handleDelete = async (id: number) => {
      await productStore.deleteProduct(id)
    }

    const formatPrice = (value: number) =>
      new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR',
        minimumFractionDigits: 2,
      }).format(value)

    return { products, loading, handleEdit, handleDelete, formatPrice }
  },
})
</script>

<template>
  <div class="pl-container">
    <div v-if="loading">Loading...</div>
    <table v-else class="pl-table">
      <thead>
        <tr>
          <th class="name-col">Product Name</th>
          <th class="price-col">Price</th>
          <th class="actions-col">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.product_id">
          <td class="name-col" data-label="Product Name">
            <span class="value">{{ product.product_name }}</span>
          </td>
          <td class="price-col" data-label="Price">
            <span class="value price">{{ formatPrice(product.product_price) }}</span>
          </td>
          <td class="actions-col" data-label="Actions">
            <div class="pl-actions">
              <button @click="handleEdit(product.product_id)" class="btn btn--edit">Edit</button>
              <button @click="handleDelete(product.product_id)" class="btn btn--delete">Delete</button>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style lang="scss">
.pl-container {
  background-color: var(--card-bg-color);
  width: 100%;
  border-radius: 2rem;
  padding: 2rem;
  box-sizing: border-box;
  backdrop-filter: blur(8px);
  margin-bottom: 9rem;
  border: 0.1rem solid var(--primary-color);

  .pl-table {
    width: 100%;
    border-collapse: collapse;

    /* MOBILE FIRST */
    thead {
      display: none;
    }

    tr {
      display: flex;
      flex-direction: column;
      gap: 1rem;
      padding: 2rem 1rem;
      border-bottom: 0.1rem solid gray;

      &:last-child {
        border-bottom: none;
      }
    }

    td {
      display: flex;
      flex-direction: column;
      gap: 0.5rem;
      border: none;

      &::before {
        content: attr(data-label);
        font-size: 1.2rem;
        font-weight: 700;
        text-transform: uppercase;
        color: var(--text-color);
      }

      .value {
        font-size: 1.6rem;
        font-weight: 500;
        color: var(--text-color);

        &.price {
          color: var(--primary-color);
          font-size: 1.8rem;
          font-weight: 700;
        }
      }
    }

    .pl-actions {
      display: flex;
      gap: 1rem;

      .btn {
        flex: 1;
        padding: 0.7rem;
        border-radius: 1rem;
        border: none;
        font-weight: 700;
        font-size: 1.4rem;
        cursor: pointer;

        &--edit {
          background-color: var(--primary-color);
          color: white;
        }

        &--delete {
          background-color: #ef4444;
          color: white;
        }
      }
    }

    /* DESKTOP */
    @media (min-width: 1024px) {
      thead {
        display: table-header-group;

        th {
          padding: 1rem;
          border-bottom: 0.2rem solid var(--primary-color);
          font-size: 1.6rem;
          font-weight: 800;
          text-transform: uppercase;
          color: var(--text-color);
          text-align: left;
        }

        .actions-col {
          text-align: center;
        }
      }

      tr {
        display: table-row;
      }

      td {
        display: table-cell;
        padding: 1rem;
        vertical-align: middle;
        &::before {
          display: none;
        }
      }

      .pl-actions {
        margin-top: 0;
        justify-content: center;

        .btn {
          flex: 0 1 auto;
          width: 9rem;
          padding: 0.6rem;
          font-size: 1.5rem;
        }
      }

      .value {
        font-size: 1.5rem;
      }
    }
  }
}
</style>
