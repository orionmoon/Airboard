<template>
  <div class="app-layout">
    <!-- Sidebar (non zoomée) -->
    <Sidebar v-if="isAuthenticated && !isAuthPage" />

    <!-- Overlay mobile -->
    <div
      v-if="isAuthenticated && !isAuthPage && appStore.sidebarOpen"
      @click="appStore.setSidebarOpen(false)"
      class="fixed inset-0 z-40 bg-black bg-opacity-50 lg:hidden"
    />

    <!-- Main content -->
    <div :class="mainContentClasses">
      <!-- Zoom wrapper (contenu zoomé) -->
      <div v-if="isAuthenticated && !isAuthPage" class="zoom-wrapper" :style="zoomWrapperStyle">
        <!-- Loading global -->
        <LoadingOverlay v-if="appStore.isLoading" />

        <!-- Router view -->
        <router-view v-slot="{ Component, route }">
          <transition name="page" mode="out-in">
            <component :is="Component" :key="route.path" />
          </transition>
        </router-view>
      </div>

      <!-- Auth pages (non zoomées) -->
      <template v-else>
        <!-- Loading global -->
        <LoadingOverlay v-if="appStore.isLoading" />

        <!-- Router view -->
        <router-view v-slot="{ Component, route }">
          <transition name="page" mode="out-in">
            <component :is="Component" :key="route.path" />
          </transition>
        </router-view>
      </template>
    </div>

    <!-- Notifications -->
    <NotificationContainer />
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'

// Components
import Sidebar from '@/components/layout/Sidebar.vue'
import LoadingOverlay from '@/components/ui/LoadingOverlay.vue'
import NotificationContainer from '@/components/ui/NotificationContainer.vue'

// Stores
const authStore = useAuthStore()
const appStore = useAppStore()
const route = useRoute()
const router = useRouter()

// Computed
const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAuthPage = computed(() => route.path.startsWith('/auth'))

const mainContentClasses = computed(() => {
  if (isAuthenticated.value && !isAuthPage.value) {
    return 'main-content'
  } else {
    return 'min-h-screen bg-gray-900'
  }
})

const zoomWrapperStyle = computed(() => {
  // Utiliser la propriété zoom CSS au lieu de transform scale
  // Cela préserve les coordonnées pour Vue DevTools
  return {
    zoom: appStore.zoomScale
  }
})

// Lifecycle
onMounted(async () => {
  // Charger les préférences depuis le localStorage
  const hasStoredAuth = authStore.loadFromStorage()
  appStore.loadFromStorage()

  // Si pas d'authentification stockée, tenter SSO auto-login
  // SAUF si on est sur un callback OAuth (Microsoft/Google)
  const isOAuthCallback = route.path.includes('/auth/oauth/') && route.query.code

  if (!hasStoredAuth && !isOAuthCallback) {
    try {
      const ssoResult = await authStore.autoLoginSSO()
      if (ssoResult) {
        console.log('✅ SSO auto-login réussi')

        // Rediriger vers le dashboard après SSO réussi
        // Toujours rediriger si on est sur une page d'auth
        if (route.path.startsWith('/auth')) {
          console.log('🚀 Redirection vers le dashboard après SSO')
          const redirectPath = route.query.redirect || '/dashboard'
          await router.push(redirectPath)
        } else if (route.path === '/') {
          // Si on est à la racine, rediriger vers dashboard
          await router.push('/dashboard')
        }
      }
    } catch (error) {
      // Échec SSO silencieux (mode classique)
      console.log('ℹ️ SSO non disponible, mode classique activé')
    }
  }

  // Initialiser le watcher du thème système
  const cleanupThemeWatcher = appStore.initSystemThemeWatcher()

  // Charger les settings de l'application (admin uniquement)
  if (authStore.isAuthenticated && authStore.isAdmin) {
    try {
      await appStore.refreshAppSettings()
    } catch (error) {
      console.error('Failed to load app settings on startup:', error)
    }
  }

  // Nettoyer au démontage
  onUnmounted(() => {
    if (cleanupThemeWatcher) {
      cleanupThemeWatcher()
    }
  })
})

</script>

<style>
/* Zoom wrapper */
.zoom-wrapper {
  /* Avec la propriété zoom CSS, pas besoin de compensation de taille */
  width: 100%;
  min-height: 100vh;
}

/* Transitions globales */
.page-enter-active,
.page-leave-active {
  transition: all 0.3s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateX(10px);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-10px);
}

/* Loading overlay */
.loading-overlay {
  backdrop-filter: blur(2px);
}
</style>