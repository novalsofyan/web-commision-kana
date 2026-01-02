import { createRouter, createWebHistory } from 'vue-router'
import axios from 'axios'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/login',
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/pages/HomePage.vue'),
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('@/pages/DashboardPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/settings',
      name: 'Settings',
      component: () => import('@/pages/SettingsPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/pages/NotFound.vue'),
    },
  ],
})

// Middleware Global (Navigation Guard)
router.beforeEach(async (to, from, next) => {
  const token = sessionStorage.getItem('auth_token')

  // 1. Cek apakah halaman yang dituju (atau parent-nya) butuh auth
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // Kalau nggak ada token, langsung tendang ke Login
    if (!token) {
      return next({ name: 'Login' })
    }

    try {
      // Validasi token ke backend
      const baseUrl = import.meta.env.VITE_API_BASE_URL || 'http://192.168.1.6:8080'
      await axios.get(`${baseUrl}/api/auth/me`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      next()
    } catch (err) {
      console.error('Auth Validation Failed', err)
      sessionStorage.removeItem('auth_token')
      next({ name: 'Login' })
    }
  }

  // 2. Kalau user sudah login (punya token) tapi iseng mau buka halaman login
  else if (to.name === 'Login' && token) {
    next({ name: 'Dashboard' })
  }

  // 3. Halaman umum (Login atau 404)
  else {
    next()
  }
})

export default router
