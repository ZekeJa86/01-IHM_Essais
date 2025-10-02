<template>
  <Loader v-if="loading" />
  <div v-else>
    <Header />
    <RouterView />
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
