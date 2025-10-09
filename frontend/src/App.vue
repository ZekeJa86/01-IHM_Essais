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

/* Alternative 1: Slide moderne avec scale (décommenter pour utiliser) */
/*
.scale-fade-enter-active, 
.scale-fade-leave-active {
  transition: all 0.5s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.scale-fade-enter-active {
  position: relative;
}

.scale-fade-leave-active {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
}

.scale-fade-enter-from {
  opacity: 0;
  transform: translateX(100%) scale(0.8) rotate(5deg);
}

.scale-fade-enter-to {
  opacity: 1;
  transform: translateX(0) scale(1) rotate(0deg);
}

.scale-fade-leave-from {
  opacity: 1;
  transform: translateX(0) scale(1) rotate(0deg);
}

.scale-fade-leave-to {
  opacity: 0;
  transform: translateX(-100%) scale(0.8) rotate(-5deg);
}
*/

/* Alternative 2: Flip 3D moderne (décommenter pour utiliser) */
/*
.transition-container {
  perspective: 1200px;
}

.scale-fade-enter-active, 
.scale-fade-leave-active {
  transition: all 0.6s cubic-bezier(0.68, -0.55, 0.27, 1.55);
}

.scale-fade-enter-active {
  position: relative;
}

.scale-fade-leave-active {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
}

.scale-fade-enter-from {
  opacity: 0;
  transform: rotateY(90deg) scale(0.9);
  transform-origin: left center;
}

.scale-fade-enter-to {
  opacity: 1;
  transform: rotateY(0deg) scale(1);
  transform-origin: left center;
}

.scale-fade-leave-from {
  opacity: 1;
  transform: rotateY(0deg) scale(1);
  transform-origin: right center;
}

.scale-fade-leave-to {
  opacity: 0;
  transform: rotateY(-90deg) scale(0.9);
  transform-origin: right center;
}
*/

/* Alternative 3: Zoom élégant (décommenter pour utiliser) */
/*
.scale-fade-enter-active {
  transition: all 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  position: relative;
}

.scale-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.55, 0.09, 0.68, 0.53);
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
}

.scale-fade-enter-from {
  opacity: 0;
  transform: scale(1.2);
  filter: blur(20px);
}

.scale-fade-enter-to {
  opacity: 1;
  transform: scale(1);
  filter: blur(0);
}

.scale-fade-leave-from {
  opacity: 1;
  transform: scale(1);
  filter: blur(0);
}

.scale-fade-leave-to {
  opacity: 0;
  transform: scale(0.8);
  filter: blur(20px);
}
*/
</style>