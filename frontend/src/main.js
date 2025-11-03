import { inject } from '@vercel/analytics'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router.js'
import './style.css'


inject()

createApp(App)
    .use(router)
    .mount('#app')
    