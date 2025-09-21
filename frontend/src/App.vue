<!-- Architecture page header -> main variable -> footer -->
<template>
  <Header />
  <main>
    <router-view /> <!-- Ici le contenu change -->
  </main>
  <Footer />
</template>

<script setup>
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
</script>

<style lang="postcss">
html.dark { 
  color-scheme: dark;
}

body {
  @apply bg-slate-50 dark:bg-slate-800 dark:text-slate-50;
}
</style>
