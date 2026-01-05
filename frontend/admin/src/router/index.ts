import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

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

// Middleware Global
router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  // Initialize auth store dulu
  if (!authStore.isInitialized) {
    await authStore.initialize()
  }

  // Cek apakah halaman butuh auth
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    // Validasi auth
    const isValid = await authStore.validateAuth()

    if (!isValid) {
      return next({ name: 'Login' })
    }

    next()
  }

  // Kalau user sudah login tapi mau ke halaman login atau home
  else if ((to.name === 'Login' || to.name === 'Home') && authStore.isAuthenticated) {
    next({ name: 'Dashboard' })
  }

  // Halaman umum
  else {
    next()
  }
})

export default router
