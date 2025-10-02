<template>
  <div class="app-layout flex flex-col min-h-screen">
    <Header />
    <main class="content-area flex flex-col flex-1 min-h-0">
      <RouterView />
    </main>
    <Footer />
  </div>
</template>


<script setup>
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import { RouterView } from 'vue-router';
</script>

<style>
/* Layout global */
.app-layout {
  @apply h-screen flex flex-col;
}
</style>
