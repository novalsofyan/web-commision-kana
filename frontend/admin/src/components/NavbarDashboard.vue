<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { House, User, Settings } from 'lucide-vue-next'

const menuItems = [
  { label: 'Home', path: '/dashboard', iconImage: House },
  { label: 'Profile', path: '/profile', iconImage: User },
  { label: 'Settings', path: '/settings', iconImage: Settings },
]
</script>

<template>
  <nav class="nav-bottom">
    <div class="nav-wrapper">
      <div class="scroll-container">
        <RouterLink v-for="(item, index) in menuItems" :key="index" :to="item.path" class="nav-item">
          <component :is="item.iconImage" class="icon-style" />
          <span class="label-text">{{ item.label }}</span>
        </RouterLink>
      </div>
    </div>
  </nav>
</template>

<style scoped lang="scss">
.nav-bottom {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  padding: 0 1rem 1rem 1rem;
  z-index: 100;
}

/* Wrapper navbar */
.nav-wrapper {
  pointer-events: auto;
  display: flex;
  width: 100%;
  max-width: 1200px;
  height: 70px;
  background: var(--bg-color);
  border: 1px solid var(--primary-color);
  border-radius: 15px;
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.scroll-container {
  display: flex;
  align-items: center;
  width: 100%;
  overflow-x: auto;
  padding: 0 1.5rem;
  gap: 1.5rem;

  /* Sembunyikan Scrollbar */
  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  scrollbar-width: none;
  justify-content: space-around;
}

/* Styling per item menu */
.nav-item {
  flex: 0 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-decoration: none;
  color: var(--text-color);
  transition: all 0.3s ease;
  min-width: 65px;

  .label-text {
    font-size: 1.2rem;
    color: var(--text-color);
  }

  /* Style saat link aktif/diklik */
  &.router-link-active {
    color: var(--primary-color);
  }

  &:active {
    transform: scale(0.95);
  }
}

/* Label text menu */
.label-text {
  font-size: 10px;
  white-space: nowrap;
}
</style>
