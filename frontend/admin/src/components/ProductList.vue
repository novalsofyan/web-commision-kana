<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup() {
    const products = ref([
      { id: 1, name: 'Kana Commission - Basic Pack', price: 150000 },
      { id: 2, name: 'Fullstack Mastery Web', price: 500000 },
    ])

    const handleEdit = (id: number) => {
      console.log(id)
    }

    const handleDelete = (id: number) => {
      console.log(id)
    }

    const formatPrice = (value: number) => {
      return new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR',
        minimumFractionDigits: 2,
      }).format(value)
    }

    return { products, handleEdit, handleDelete, formatPrice }
  },
})
</script>

<template>
  <div class="pl-container">
    <table class="pl-table">
      <thead>
        <tr>
          <th class="name-col">Product Name</th>
          <th class="price-col">Price</th>
          <th class="actions-col">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.id">
          <td data-label="Product Name" class="name-col">
            <span class="value">{{ product.name }}</span>
          </td>
          <td data-label="Price" class="price-col">
            <span class="value price">{{ formatPrice(product.price) }}</span>
          </td>
          <td data-label="Actions" class="actions-col">
            <div class="pl-actions">
              <button @click="handleEdit(product.id)" class="btn btn--edit">Edit</button>
              <button @click="handleDelete(product.id)" class="btn btn--delete">Delete</button>
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

    // Mobile First
    thead {
      display: none;
    }

    tr {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      padding: 2.5rem 1rem;
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
        color: var(--text-color);
        font-size: 1.6rem;
        font-weight: 500;

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
      margin-top: 0.8rem;

      .btn {
        flex: 1;
        padding: 1.2rem;
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

    // Desktop
    @media (min-width: 1024px) {
      thead {
        display: table-header-group;

        th {
          padding: 0 1rem 1.5rem 1rem;
          border-bottom: 0.2rem solid var(--primary-color);
          font-size: 1.6rem;
          font-weight: 800;
          text-transform: uppercase;
          color: var(--text-color);
          text-align: left;
          letter-spacing: 0.1rem;
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
        padding: 2rem 1rem;
        vertical-align: middle;

        &::before {
          display: none;
        }

        .value {
          font-size: 1.5rem;
        }
      }

      .actions-col {
        width: 20rem;
      }

      .pl-actions {
        justify-content: center;
        margin-top: 0;

        .btn {
          flex: 0 1 auto;
          width: 9rem;
          padding: 0.8rem;
          font-size: 1.2rem;
        }
      }
    }
  }
}
</style>
