<script setup>
const navLinkClass = 
"px-7 py-2 text-white rounded-lg font-bold transform hover:-translate-y-1 transition duration-400 hover:text-blue-600";
</script>
<template>
  <header>
    <div class="fixed top-0 left-0 w-full flex flex-row dark:bg-slate-900/80 h-16 justify-start items-center gap-4 px-4 z-50">
      <img src="../assets/images/KDEV_light.png" alt="Logo KDEV" class="h-20 w-auto" />
      <router-link to="/" :class="navLinkClass">Home</router-link>
      <router-link to="/profil" :class="navLinkClass">Profil</router-link>
      <router-link to="/monitoring" :class="navLinkClass">Monitoring</router-link>
      <router-link to="/parameters" :class="navLinkClass">Parameters</router-link>
      <router-link to="/contact" :class="navLinkClass">Contact</router-link>
      <DarkMode />
    </div>
  </header>
  <!-- Décalage du contenu pour ne pas être caché par le header -->
  <main class="pt-16">
    <!-- ton contenu principal ici -->
  </main>
</template>
