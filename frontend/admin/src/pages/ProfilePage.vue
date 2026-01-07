<script setup lang="ts">
import NavbarDashboard from '@/components/NavbarDashboard.vue'
</script>

<template>
  <main class="profile-container">
    <h1 class="profile-title">Update Profile</h1>
    <div class="profile-content">
      <div class="profile-card">
        <h2>Profil</h2>
      </div>
    </div>
    <NavbarDashboard />
  </main>
</template>

<style lang="scss">
.profile-container {
  color: var(--text-color);
  text-align: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.profile-title {
  font-size: 4rem;
  font-weight: bold;
  padding-bottom: 1rem;
  border-bottom: 1px solid var(--primary-color);
}

.profile-content {
  margin-top: 2rem;
}

.profile-card {
  width: 100%;
  border: 1px solid var(--primary-color);
  border-radius: 1rem;
  padding: 2rem;
}
</style>
