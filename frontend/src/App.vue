<template>
  <!-- Loader pendant le chargement -->
  <Loader v-if="loading" />
  
  <div v-else class="min-h-screen bg-gradient-to-b from-slate-900 via-slate-800 to-slate-900 text-white flex flex-col">
    <Header />
    <div class="transition-container flex-1">
      <RouterView v-slot="{ Component, route }">
        <Transition name="scale-fade" mode="out-in">
          <div :key="route.fullPath" class="view-wrapper">
            <component :is="Component" />
          </div>
        </Transition>
      </RouterView>
    </div>
    <Footer />
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import Loader from './components/Loader.vue'
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
import { RouterView, useRouter } from 'vue-router'

const loading = ref(true)
const router = useRouter()

onMounted(() => {
  // Simule un temps de chargement de 2 secondes
  setTimeout(() => {
    loading.value = false
  }, 2000)
})

// Fonction pour scroller en haut
const scrollToTop = () => {
  nextTick(() => {
    window.scrollTo({ top: 0, left: 0, behavior: 'instant' })
    document.documentElement.scrollTop = 0
    document.body.scrollTop = 0
  })
}

// Écouter les changements de route
router.afterEach(() => {
  scrollToTop()
})
</script>

<style>
/* Empêcher le débordement sur tout le document */
html, body {
  overflow-x: hidden;
  max-width: 100vw;
}

/* Container pour les transitions */
.transition-container {
  position: relative;
  width: 100%;
  display: flex;
  flex-direction: column;
  min-height: 0;
}

/* Wrapper pour chaque vue */
.view-wrapper {
  width: 100%;
  flex: 1;
  display: flex;
  flex-direction: column;
}

/* Transition moderne : Scale + Fade avec blur */
.scale-fade-enter-active {
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

.scale-fade-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.6, 1);
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
}

.scale-fade-enter-from {
  opacity: 0;
  transform: scale(0.95) translateY(20px);
  filter: blur(10px);
}

.scale-fade-enter-to {
  opacity: 1;
  transform: scale(1) translateY(0);
  filter: blur(0);
}

.scale-fade-leave-from {
  opacity: 1;
  transform: scale(1) translateY(0);
  filter: blur(0);
}

.scale-fade-leave-to {
  opacity: 0;
  transform: scale(1.05) translateY(-20px);
  filter: blur(10px);
}

/* Empêcher le débordement sur tout le document */
html, body {
  overflow-x: hidden;
  max-width: 100vw;
}

/* Masquer la scrollbar tout en gardant le scroll fonctionnel */
/* Pour Firefox */
* {
  scrollbar-width: none;
}

/* Pour Chrome, Safari et Edge */
*::-webkit-scrollbar {
  display: none;
}

/* Pour IE et Edge legacy */
* {
  -ms-overflow-style: none;
}
</style>