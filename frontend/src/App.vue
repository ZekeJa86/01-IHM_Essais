<!-- filepath: c:\Users\sylva\myproject\frontend\src\App.vue -->
<template>
  <Home />
</template>

<script setup>
import Home from './components/Home.vue'
</script>

<style lang="postcss">
html.dark { 
  color-scheme: dark;
}

body {
  @apply bg-slate-50 dark:bg-slate-800 dark:text-slate-50;
}
</style>
