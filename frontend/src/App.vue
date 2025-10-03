<template>
  <!-- Loader pendant le chargement -->
  <Loader v-if="loading" />
  
<div v-else class="min-h-screen bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900 text-white">
  <Header />
  <Transition name="slide-right" mode="out-in">
    <RouterView :key="$route.fullPath" />
  </Transition>
  <Footer />
</div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Loader from './components/Loader.vue'
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import { RouterView } from 'vue-router'

const loading = ref(true)

onMounted(() => {
  // Simule un temps de chargement de 2 secondes
  setTimeout(() => {
    loading.value = false
  }, 2000)
})
</script>

<style>
/* Transition slide vers la droite */
.slide-right-enter-active, .slide-right-leave-active {
  transition: all 0.5s ease;
}

.slide-right-enter-from {
  transform: translateX(100%);
  opacity: 0;
}

.slide-right-enter-to {
  transform: translateX(0);
  opacity: 1;
}

.slide-right-leave-from {
  transform: translateX(0);
  opacity: 1;
}

.slide-right-leave-to {
  transform: translateX(100%);
  opacity: 0;
}
</style>
