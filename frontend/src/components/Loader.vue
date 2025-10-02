<template>
  <div class="fixed inset-0 flex items-center justify-center bg-black z-50">
    <img 
      src="../assets/images/KDEV_light.png" 
      alt="Logo KDEV" 
      class="h-36 w-auto animate-bounce"
    />
  </div>
</template>

<script setup>
// Optionnel : tu peux gérer l'affichage via une prop ou un state
</script>

<style>
/* Si tu veux une animation custom différente de bounce */
@keyframes spin-slow {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
