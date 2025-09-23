import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'
import Profil from './components/Profil.vue'
import Monitoring from './components/Monitoring.vue'
import Logs from './components/Logs.vue'
import Parameters from './components/Parameters.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/profil', component: Profil },
  { path: '/parameters', component: Parameters },
  { path: '/monitoring', component: Monitoring },
  { path: '/logs', component: Logs }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
