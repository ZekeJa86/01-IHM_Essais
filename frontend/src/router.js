import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'
import Profil from './components/Profil.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/profil', component: Profil }
]

export default createRouter({
  history: createWebHistory(),
  routes
})