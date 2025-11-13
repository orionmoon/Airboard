import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import i18n from '@/i18n'

// Styles CSS
import './assets/css/main.css'

// Services
import { setupInterceptors } from './services/api'

// Configuration de l'application
const app = createApp(App)
const pinia = createPinia()

// Plugins
app.use(pinia)
app.use(router)
app.use(i18n)

// Configuration des intercepteurs API
setupInterceptors(router)

// Montage de l'application
app.mount('#app')

// Configuration globale pour le mode debug
if (import.meta.env.DEV) {
  app.config.globalProperties.$log = console.log
  window.__VUE_APP__ = app
}