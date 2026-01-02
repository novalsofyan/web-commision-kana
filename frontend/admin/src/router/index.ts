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
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/pages/NotFound.vue'),
    },
  ],
})

// Middleware Global
router.beforeEach(async (to, from, next) => {
  const token = sessionStorage.getItem('auth_token')

  // Cek apakah halaman yang dituju butuh login
  if (to.meta.requiresAuth) {
    if (!token) {
      return next({ name: 'Login' })
    }

    try {
      await axios.get('http://192.168.1.6:8080/api/auth/me', {
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

  // Kalau user sudah login tapi mau buka halaman login
  else if (to.name === 'Login' && token) {
    next({ name: 'Dashboard' })
  } else {
    next()
  }
})

export default router
