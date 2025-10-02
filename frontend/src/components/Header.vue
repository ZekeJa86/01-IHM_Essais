<script setup>
const navLinkClass = 
  "px-4 py-2 rounded-xl text-white font-medium bg-slate-900 hover:bg-blue-700 transition-colors duration-300 shadow-md hover:shadow-lg";
console.log("Header.vue monté ✅");

</script>

<template>
  <header>
    <div class="flex flex-row bg-slate-600 h-16 justify-start items-center gap-4 px-4">
      <img src="../assets/images/KDEV_light.png" alt="Logo KDEV" class="h-20 w-auto" />

      <router-link to="/" :class="navLinkClass">Home</router-link>
      <router-link to="/profil" :class="navLinkClass">Profil</router-link>
      <router-link to="/monitoring" :class="navLinkClass">Monitoring</router-link>
      <router-link to="/parameters" :class="navLinkClass">Parameters</router-link>
      <router-link to="/logs" :class="navLinkClass">Logs</router-link>

      <DarkMode />
    </div>
  </header>
</template>



