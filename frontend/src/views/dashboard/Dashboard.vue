<template>
  <div class="content-area">
    <!-- Header avec paramètres dynamiques -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">{{ appSettings.dashboard_title || 'Dashboard' }}</h1>
          <p class="page-subtitle">
            {{ appSettings.welcome_message || 'Welcome to your application portal' }}
          </p>
        </div>
        
        <!-- Bouton menu mobile -->
        <button
          @click="appStore.toggleSidebar()"
          class="lg:hidden btn btn-secondary"
        >
          <Icon icon="mdi:menu" class="h-5 w-5" />
        </button>
      </div>
    </div>

    <!-- Statistiques (Admin seulement) -->
    <div v-if="authStore.isAdmin && dashboard?.stats" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="card hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="h-10 w-10 bg-green-500 dark:bg-green-900 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:folder-multiple" class="h-5 w-5 text-white" />
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600 dark:text-gray-400">App Groups</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ dashboard.stats.total_app_groups }}
            </p>
          </div>
        </div>
      </div>

      <div class="card hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="h-10 w-10 bg-blue-500 dark:bg-blue-900 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:application" class="h-5 w-5 text-white" />
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600 dark:text-gray-400">Applications</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ dashboard.stats.total_applications }}
            </p>
          </div>
        </div>
      </div>

      <div class="card hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="h-10 w-10 bg-purple-500 dark:bg-purple-900 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:account-multiple" class="h-5 w-5 text-white" />
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600 dark:text-gray-400">Users</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ dashboard.stats.total_users }}
            </p>
          </div>
        </div>
      </div>

      <div class="card hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="h-10 w-10 bg-yellow-500 dark:bg-yellow-900 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:account-group" class="h-5 w-5 text-white" />
            </div>
          </div>
          <div class="ml-4">
            <p class="text-sm font-medium text-gray-600 dark:text-gray-400">Groups</p>
            <p class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ dashboard.stats.total_groups }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center py-12">
      <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-gray-400" />
    </div>

    <!-- Empty state -->
    <div v-else-if="!dashboard?.app_groups?.length" class="empty-state">
      <Icon icon="mdi:application" class="empty-state-icon" />
      <h3 class="empty-state-title">No Applications Available</h3>
      <p class="empty-state-description">
        {{ authStore.isAdmin 
          ? 'Start by creating application groups and applications.' 
          : 'Contact your administrator to get access to applications.' 
        }}
      </p>
      <div v-if="authStore.isAdmin" class="mt-6">
        <router-link to="/admin/app-groups" class="btn btn-primary">
          <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
          Create App Group
        </router-link>
      </div>
    </div>

    <!-- Groupes d'applications -->
    <div v-else class="space-y-8">
      <div
        v-for="appGroup in dashboard.app_groups"
        :key="appGroup.id"
        class="fade-in"
      >
        <!-- En-tête du groupe -->
        <div class="flex items-center space-x-3 mb-6">
          <div 
            class="h-10 w-10 rounded-lg flex items-center justify-center shadow-sm"
            :style="{ backgroundColor: appGroup.color || '#10b981' }"
          >
            <Icon :icon="appGroup.icon || 'mdi:folder'" class="h-5 w-5 text-white" />
          </div>
          <div>
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
              {{ appGroup.name }}
            </h2>
            <p v-if="appGroup.description" class="text-sm text-gray-600 dark:text-gray-400">
              {{ appGroup.description }}
            </p>
          </div>
        </div>

        <!-- Applications du groupe - version simple et sobre -->
        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div
            v-for="app in appGroup.applications"
            :key="app.id"
            class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg p-3 cursor-pointer transition-all duration-150 hover:bg-gray-50 dark:hover:bg-gray-800 hover:shadow-md hover:-translate-y-0.5"
            @click="openApplication(app)"
          >
            <!-- Entête: icône + titre alignés -->
            <div class="flex items-center gap-2">
              <div 
                class="h-10 w-10 rounded-lg flex items-center justify-center"
                :style="{ backgroundColor: app.color || '#6366f1' }"
              >
                <Icon :icon="app.icon || 'mdi:application'" class="h-5 w-5 text-white" />
              </div>
              <h3 class="font-semibold text-gray-900 dark:text-white text-sm leading-tight flex-1 truncate">
                {{ app.name }}
              </h3>
              <Icon v-if="app.open_in_new_tab" icon="mdi:open-in-new" class="h-4 w-4 text-gray-500 dark:text-gray-400" />
            </div>

            <!-- Description sous le titre -->
            <p v-if="app.description" class="text-xs text-gray-600 dark:text-gray-400 mt-2">
              {{ app.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { Icon } from '@iconify/vue'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'
import { dashboardService, adminService } from '@/services/api'

const authStore = useAuthStore()
const appStore = useAppStore()

// État local
const dashboard = ref(null)
const appSettings = ref({})
const isLoading = ref(false)

// Fonctions
const uniqueById = (items) => {
  if (!Array.isArray(items)) return []
  const seen = new Set()
  const out = []
  for (const item of items) {
    const id = item?.id ?? JSON.stringify(item)
    if (!seen.has(id)) {
      seen.add(id)
      out.push(item)
    }
  }
  return out
}

const loadDashboard = async () => {
  try {
    isLoading.value = true
    const data = await dashboardService.getDashboard()
    // Déduplication côté client par sécurité
    if (data?.app_groups) {
      data.app_groups = uniqueById(data.app_groups).map(g => ({
        ...g,
        applications: uniqueById(g.applications || [])
      }))
    }
    dashboard.value = data
  } catch (error) {
    console.error('Error loading dashboard:', error)
    appStore.showError('Failed to load dashboard')
  } finally {
    isLoading.value = false
  }
}

const loadAppSettings = async () => {
  try {
    // Utiliser le store si dispo
    if (appStore.appSettings) {
      appSettings.value = appStore.appSettings
      return
    }
    // Ne pas appeler l'API si l'utilisateur n'est pas admin (évite 403)
    if (!authStore.isAdmin) {
      appSettings.value = {
        dashboard_title: 'Dashboard',
        welcome_message: 'Welcome to your application portal'
      }
      return
    }
    const settings = await adminService.getAppSettings()
    appSettings.value = settings
    appStore.setAppSettings(settings)
  } catch (error) {
    console.error('Error loading app settings:', error)
    // Use defaults if settings can't be loaded
    appSettings.value = {
      dashboard_title: 'Dashboard',
      welcome_message: 'Welcome to your application portal'
    }
  }
}

const openApplication = (app) => {
  if (app.open_in_new_tab) {
    window.open(app.url, '_blank', 'noopener,noreferrer')
  } else {
    window.location.href = app.url
  }
  
  // Analytics ou logging (optionnel)
  console.log(`Application opened: ${app.name} - ${app.url}`)
}

// 🔧 FIX: Watcher pour les changements de settings
watch(() => appStore.settingsLastUpdated, async () => {
  console.log('Settings updated, reloading dashboard settings...')
  await loadAppSettings()
}, { immediate: false })

// Lifecycle
onMounted(async () => {
  await Promise.all([
    loadDashboard(),
    loadAppSettings()
  ])
})
</script>

<style scoped>
/* Cartes d'applications améliorées */
.app-card {
  @apply relative bg-white dark:bg-gray-800 border-2 border-gray-200 dark:border-gray-700 rounded-xl p-4 transition-all duration-200 hover:border-green-500 dark:hover:border-green-500 hover:shadow-xl hover:-translate-y-1 flex flex-col;
  min-height: 140px;
}

.app-icon-sm {
  @apply h-10 w-10 rounded-lg flex items-center justify-center shadow-md transition-transform duration-200;
}

.app-title-sm {
  @apply font-semibold text-gray-900 dark:text-white text-sm leading-tight flex-1;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.app-description {
  @apply text-xs text-gray-600 dark:text-gray-400 leading-relaxed;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Effet de transparence légère */
.app-card:hover {
  background-color: rgba(241, 245, 249, 0.8);
}

.dark .app-card:hover {
  background-color: rgba(55, 65, 81, 0.9);
}

/* Couleur hover pour les stats cards - now handled by Tailwind dark: classes */

/* Animation d'apparition */
.fade-in {
  animation: fadeIn 0.5s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Style pour les descriptions tronquées */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>