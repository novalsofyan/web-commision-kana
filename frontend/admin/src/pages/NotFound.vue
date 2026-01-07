<template>
  <main class="not-found">
    <h1>404</h1>
    <h2>Waduh! Halamannya Gak Ada...</h2>
    <p>Kayaknya kamu nyasar atau link-nya udah kedaluwarsa.</p>
    <router-link to="/" class="back-btn">Balik ke Home</router-link>
  </main>
</template>

<style scoped>
.not-found {
  color: var(--text-color);
  background: var(--bg-color);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 80vh;
  font-family: sans-serif;
  text-align: center;
}
h1 {
  font-size: 6rem;
  margin: 0;
  color: #ff4757;
}
h2 {
  margin: 1rem 0;
}
.back-btn {
  margin-top: 1.5rem;
  padding: 0.8rem 1.5rem;
  background: #2f3542;
  color: white;
  text-decoration: none;
  border-radius: 8px;
}
.back-btn:hover {
  background: #57606f;
}
</style>
