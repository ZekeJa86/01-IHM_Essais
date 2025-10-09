import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'
import Profil from './components/Profil.vue'
import Privacy from './components/Privacy.vue'
import Contact from './components/Contact.vue'
import Portfolio from './components/Portfolio.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/profil', component: Profil },
  { path: '/portfolio', component: Portfolio },
  { path: '/privacy', component: Privacy },
  { path: '/contact', component: Contact }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
