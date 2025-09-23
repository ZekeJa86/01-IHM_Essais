<script setup>
import DarkMode from './DarkMode.vue' // chemin relatif depuis Header.vue
console.log("Header.vue monté ✅");
</script>

<template>
  <header>
    <div class="nav">
      <img src="../assets/images/KDEV_light.png" alt="Logo KDEV" class="logo" />
      <router-link to="/">Home</router-link>
      <router-link to="/profil">Profil</router-link>
      <router-link to="/monitoring">Monitoring</router-link>
      <router-link to="/parameters">Parameters</router-link>
      <router-link to="/logs">Logs</router-link>
      <DarkMode />
    </div>
  </header>
</template>


<style>
.nav {
  display: flex;              /* active Flexbox */
  align-items: center;        /* centre verticalement */
  gap: 20px;                  /* espace entre les éléments */
  background-color:  rgba(6, 6, 6, 0.8);
  padding: 0 10px;            /* padding horizontal */
  height: 70px;              /* hauteur de la nav */
}

.nav a {
  color: #f2f2f2;
  text-decoration: none;
  font-size: 20px;
  padding: 10px 15px;
  transition: all 0.3s;
}

.nav a:hover {
  background-color: #cfcfcf;
  color: #000;
  border-radius: 4px;
}

.logo {
  height: 80px;               /* taille du logo */
  width: auto;
}
</style>


