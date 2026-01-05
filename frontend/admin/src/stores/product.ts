import { defineStore } from 'pinia'
import api from '@/lib/axios'
import type { Product } from '@/types/product'

export const useProductStore = defineStore('products', {
  state: () => ({
    products: [] as Product[],
    loading: false,
  }),

  actions: {
    async fetchProducts() {
      this.loading = true
      try {
        const res = await api.get('/api/admin/products')
        this.products = res.data.data
      } catch (err) {
        console.error(err)
      } finally {
        this.loading = false
      }
    },

    async createProduct(payload: Omit<Product, 'product_id'>) {
      await api.post('/api/admin/products', payload)
      await this.fetchProducts()
    },

    async updateProduct(id: number, payload: Omit<Product, 'product_id'>) {
      await api.put(`/api/admin/products/${id}`, payload)
      await this.fetchProducts()
    },

    async deleteProduct(id: number) {
      await api.delete(`/api/admin/products/${id}`)
      await this.fetchProducts()
    },
  },
})
