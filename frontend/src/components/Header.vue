<script setup>
import { ref } from 'vue';

const navLinkClass = 
  "px-7 py-2 text-white rounded-lg font-bold transform hover:-translate-y-1 transition duration-400 hover:text-blue-600";

const mobileMenuOpen = ref(false);

const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value;
};

const closeMobileMenu = () => {
  mobileMenuOpen.value = false;
};
</script>

<template>
  <header>
    <!-- Header Desktop et Mobile -->
    <div class="fixed top-0 left-0 w-full flex flex-row h-16 justify-between items-center gap-4 px-4 z-50 bg-gray-900/50 backdrop-blur-sm shadow-md">
      <!-- Logo -->
      <img src="../assets/images/KDEV_light.png" alt="Logo KDEV" class="h-20 w-auto" />
      
      <!-- Navigation Desktop (cachée sur mobile) -->
      <nav class="hidden md:flex flex-row items-center gap-4">
        <router-link to="/" :class="navLinkClass">Home</router-link>
        <router-link to="/profil" :class="navLinkClass">Profil</router-link>
        <router-link to="/portfolio" :class="navLinkClass">Portfolio</router-link>
        <router-link to="/contact" :class="navLinkClass">Contact</router-link>
        <DarkMode />
      </nav>

      <!-- Bouton Hamburger (visible uniquement sur mobile) -->
      <button 
        @click="toggleMobileMenu"
        class="md:hidden text-white p-2 focus:outline-none hover:text-blue-600 transition duration-300"
        aria-label="Toggle menu"
      >
        <svg 
          class="w-7 h-7 transform transition-transform duration-300" 
          :class="mobileMenuOpen ? 'rotate-90' : ''"
          fill="none" 
          stroke="currentColor" 
          viewBox="0 0 24 24"
        >
          <path 
            v-if="!mobileMenuOpen"
            stroke-linecap="round" 
            stroke-linejoin="round" 
            stroke-width="2" 
            d="M4 6h16M4 12h16M4 18h16"
          />
          <path 
            v-else
            stroke-linecap="round" 
            stroke-linejoin="round" 
            stroke-width="2" 
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>
    </div>

    <!-- Menu Mobile -->
    <Transition name="slide-fade">
      <div 
        v-if="mobileMenuOpen"
        class="fixed top-16 right-0 w-64 h-screen bg-gray-900 bg-opacity-98 z-40 md:hidden shadow-2xl"
      >
        <nav class="flex flex-col p-6 gap-4">
          <router-link 
            to="/" 
            :class="navLinkClass"
            @click="closeMobileMenu"
          >
            Home
          </router-link>
          <router-link 
            to="/profil" 
            :class="navLinkClass"
            @click="closeMobileMenu"
          >
            Profil
          </router-link>
          <router-link 
            to="/portfolio" 
            :class="navLinkClass"
            @click="closeMobileMenu"
          >
            Portfolio
          </router-link>
          <router-link 
            to="/contact" 
            :class="navLinkClass"
            @click="closeMobileMenu"
          >
            Contact
          </router-link>
          <div class="mt-4">
            <DarkMode />
          </div>
        </nav>
      </div>
    </Transition>

    <!-- Overlay pour fermer le menu en cliquant à l'extérieur -->
    <Transition name="fade">
      <div 
        v-if="mobileMenuOpen"
        @click="closeMobileMenu"
        class="fixed inset-0 bg-black bg-opacity-50 z-30 md:hidden"
      />
    </Transition>
  </header>
</template>

<style scoped>
/* Animation slide-fade combinée pour le menu mobile */
.slide-fade-enter-active {
  animation: slideAndFade 0.4s cubic-bezier(0.25, 0.8, 0.25, 1);
}

.slide-fade-leave-active {
  animation: slideAndFade 0.3s cubic-bezier(0.25, 0.8, 0.25, 1) reverse;
}

@keyframes slideAndFade {
  0% {
    transform: translateX(100%) scale(0.95);
    opacity: 0;
  }
  100% {
    transform: translateX(0) scale(1);
    opacity: 1;
  }
}

/* Animation fade pour l'overlay */
.fade-enter-active {
  animation: fadeIn 0.3s ease;
}

.fade-leave-active {
  animation: fadeIn 0.2s ease reverse;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>