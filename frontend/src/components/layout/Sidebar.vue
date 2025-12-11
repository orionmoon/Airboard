<template>
  <aside :class="['sidebar', { 'sidebar-open': appStore.sidebarOpen }]">
    <!-- Header -->
    <div class="sidebar-header">
      <div class="flex-1">
        <div class="flex items-center gap-3">
          <div class="h-8 w-8 bg-white rounded-lg flex items-center justify-center">
            <Icon :icon="appStore.appSettings?.app_icon || 'mdi:view-dashboard'" class="h-5 w-5 text-black" />
          </div>
          <div class="flex-1">
            <h1 class="sidebar-brand">{{ appStore.appSettings?.app_name || 'Airboard' }}</h1>
            <!-- Version info with update badge -->
            <div class="flex items-center gap-2 mt-0.5">
              <span class="text-xs text-gray-400">
                v{{ versionStore.versionInfo.version }}
              </span>
              <button
                v-if="versionStore.shouldShowUpdateBadge()"
                @click="showUpdateModal = true"
                class="relative inline-flex items-center gap-1 px-2 py-0.5 text-xs font-medium rounded-full bg-red-500 text-white hover:bg-red-600 transition-colors animate-pulse"
                title="Nouvelle version disponible"
              >
                <Icon icon="mdi:update" class="h-3 w-3" />
                <span>Mise à jour</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Mobile close button -->
      <button
        @click="appStore.toggleSidebar()"
        class="lg:hidden btn-ghost p-2"
      >
        <Icon icon="mdi:close" class="h-5 w-5" />
      </button>
    </div>

    <!-- Navigation -->
    <nav class="sidebar-nav">
      <!-- Dashboard -->
      <router-link
        to="/dashboard"
        :class="getLinkClasses('/dashboard')"
      >
        <Icon icon="mdi:view-dashboard" class="h-4 w-4" />
        <span>{{ $t('common.dashboard') }}</span>
      </router-link>

      <router-link
        to="/news"
        :class="getLinkClasses('/news')"
      >
        <Icon icon="mdi:newspaper" class="h-4 w-4" />
        <span>{{ $t('common.newsHub') }}</span>
      </router-link>

      <!-- Admin section -->
      <div v-if="authStore.isAdmin" class="sidebar-section">
        <div class="sidebar-section-title">
          {{ $t('common.administration') }}
        </div>
        
        <router-link
          to="/admin"
          :class="getLinkClasses('/admin', true)"
        >
          <Icon icon="mdi:chart-box" class="h-4 w-4" />
          <span>{{ $t('common.overview') }}</span>
        </router-link>
        
        <router-link
          to="/admin/app-groups"
          :class="getLinkClasses('/admin/app-groups')"
        >
          <Icon icon="mdi:folder-multiple" class="h-4 w-4" />
          <span>{{ $t('common.appGroups') }}</span>
        </router-link>
        
        <router-link
          to="/admin/applications"
          :class="getLinkClasses('/admin/applications')"
        >
          <Icon icon="mdi:apps" class="h-4 w-4" />
          <span>{{ $t('common.applications') }}</span>
        </router-link>
        
        <router-link
          to="/admin/users"
          :class="getLinkClasses('/admin/users')"
        >
          <Icon icon="mdi:account-multiple" class="h-4 w-4" />
          <span>{{ $t('common.users') }}</span>
        </router-link>
        
        <router-link
          to="/admin/groups"
          :class="getLinkClasses('/admin/groups')"
        >
          <Icon icon="mdi:account-group" class="h-4 w-4" />
          <span>{{ $t('common.groups') }}</span>
        </router-link>
        
        <router-link
          to="/admin/settings"
          :class="getLinkClasses('/admin/settings')"
        >
          <Icon icon="mdi:cog" class="h-4 w-4" />
          <span>{{ $t('common.settings') }}</span>
        </router-link>

        <router-link
          to="/admin/oauth"
          :class="getLinkClasses('/admin/oauth')"
        >
          <Icon icon="mdi:shield-key" class="h-4 w-4" />
          <span>{{ $t('common.oauth') }}</span>
        </router-link>

        <router-link
          to="/admin/analytics"
          :class="getLinkClasses('/admin/analytics')"
        >
          <Icon icon="mdi:chart-line" class="h-4 w-4" />
          <span>{{ $t('common.analytics') }}</span>
        </router-link>

        <router-link
          to="/admin/announcements"
          :class="getLinkClasses('/admin/announcements')"
        >
          <Icon icon="mdi:bullhorn" class="h-4 w-4" />
          <span>{{ $t('common.announcements') }}</span>
        </router-link>
      </div>

      <!-- Editor section (for admin and editor roles) -->
      <div v-if="authStore.user?.role === 'admin' || authStore.user?.role === 'editor'" class="sidebar-section">
        <div class="sidebar-section-title">
          News Hub
        </div>

        <router-link
          to="/admin/news"
          :class="getLinkClasses('/admin/news')"
        >
          <Icon icon="mdi:newspaper" class="h-4 w-4" />
          <span>News Articles</span>
        </router-link>
      </div>

      <!-- Group Admin section (for admin and group_admin roles) -->
      <div v-if="authStore.isGroupAdmin || authStore.isAdmin" class="sidebar-section">
        <div class="sidebar-section-title">
          Administration de Groupe
        </div>

        <router-link
          to="/group-admin"
          :class="getLinkClasses('/group-admin', true)"
        >
          <Icon icon="mdi:shield-account" class="h-4 w-4" />
          <span>Tableau de bord</span>
        </router-link>

        <router-link
          to="/group-admin/app-groups"
          :class="getLinkClasses('/group-admin/app-groups')"
        >
          <Icon icon="mdi:folder-multiple" class="h-4 w-4" />
          <span>Mes AppGroups</span>
        </router-link>

        <router-link
          to="/group-admin/applications"
          :class="getLinkClasses('/group-admin/applications')"
        >
          <Icon icon="mdi:apps" class="h-4 w-4" />
          <span>Mes Applications</span>
        </router-link>

        <router-link
          to="/group-admin/news"
          :class="getLinkClasses('/group-admin/news')"
        >
          <Icon icon="mdi:newspaper-variant" class="h-4 w-4" />
          <span>News de Groupe</span>
        </router-link>
      </div>
    </nav>

    <!-- Footer -->
    <div class="mt-auto border-t border-gray-700 p-4">
      <!-- Zoom Control -->
      <div class="flex items-center justify-between mb-4">
        <span class="text-sm text-gray-400">{{ $t('common.zoom') }}</span>
        <ZoomControl v-model="appStore.zoomLevel" @update:modelValue="appStore.setZoomLevel" />
      </div>

      <!-- Theme toggle -->
      <div class="flex items-center justify-between mb-4">
        <span class="text-sm text-gray-400">{{ $t('common.darkMode') }}</span>
        <button
          @click="appStore.toggleDarkMode()"
          class="theme-toggle"
        >
          <Icon
            :icon="appStore.isDarkMode ? 'mdi:weather-night' : 'mdi:weather-sunny'"
            class="h-4 w-4"
          />
        </button>
      </div>

      <!-- Language selector -->
      <div class="mb-4">
        <label class="block text-xs text-gray-400 mb-1">{{ $t('common.language') }}</label>
        <select
          class="form-select w-full"
          :value="appStore.locale"
          @change="onChangeLocale($event.target.value)"
        >
          <option value="ar">{{ $t('common.language_ar') }}</option>
          <option value="en">{{ $t('common.language_en') }}</option>
          <option value="es">{{ $t('common.language_es') }}</option>
          <option value="fr">{{ $t('common.language_fr') }}</option>
        </select>
      </div>

      <!-- User profile -->
      <div class="flex items-center space-x-3 p-3 rounded-lg hover:bg-gray-800 transition-colors">
        <div class="flex-shrink-0">
          <div class="h-8 w-8 rounded-full bg-gray-700 flex items-center justify-center">
            <span class="text-sm font-medium text-white">
              {{ authStore.userInitials }}
            </span>
          </div>
        </div>
        <div class="flex-1 min-w-0">
          <p class="text-sm font-medium text-white truncate">
            {{ authStore.userDisplayName }}
          </p>
          <p class="text-xs text-gray-400 truncate">
            {{ authStore.user?.email }}
          </p>
        </div>
        <button
          @click="handleLogout"
          class="btn-ghost p-1"
          :title="$t('common.logout')"
        >
          <Icon icon="mdi:logout" class="h-4 w-4" />
        </button>
      </div>
    </div>

    <!-- Update Modal -->
    <Teleport to="body">
      <div
        v-if="showUpdateModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black bg-opacity-50"
        @click.self="showUpdateModal = false"
      >
        <div class="bg-white dark:bg-gray-800 rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-hidden">
          <!-- Header -->
          <div class="px-6 py-4 border-b border-gray-200 dark:border-gray-700 bg-gradient-to-r from-blue-500 to-blue-600">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-3">
                <div class="h-10 w-10 bg-white rounded-lg flex items-center justify-center">
                  <Icon icon="mdi:update" class="h-6 w-6 text-blue-600" />
                </div>
                <div>
                  <h3 class="text-lg font-semibold text-white">Nouvelle version disponible</h3>
                  <p class="text-sm text-blue-100">Une mise à jour est prête à être installée</p>
                </div>
              </div>
              <button
                @click="showUpdateModal = false"
                class="text-white hover:text-gray-200 transition-colors"
              >
                <Icon icon="mdi:close" class="h-6 w-6" />
              </button>
            </div>
          </div>

          <!-- Content -->
          <div class="px-6 py-4 overflow-y-auto max-h-[60vh]">
            <div v-if="versionStore.updateDetails" class="space-y-4">
              <!-- Version comparison -->
              <div class="flex items-center justify-center gap-4 py-4">
                <div class="text-center">
                  <p class="text-sm text-gray-500 dark:text-gray-400 mb-1">Version actuelle</p>
                  <div class="px-4 py-2 bg-gray-100 dark:bg-gray-700 rounded-lg">
                    <p class="text-xl font-bold text-gray-900 dark:text-white">
                      {{ versionStore.updateDetails.currentVersion }}
                    </p>
                  </div>
                </div>
                <Icon icon="mdi:arrow-right" class="h-8 w-8 text-gray-400" />
                <div class="text-center">
                  <p class="text-sm text-gray-500 dark:text-gray-400 mb-1">Nouvelle version</p>
                  <div class="px-4 py-2 bg-green-100 dark:bg-green-900 rounded-lg">
                    <p class="text-xl font-bold text-green-600 dark:text-green-400">
                      {{ versionStore.updateDetails.latestVersion }}
                    </p>
                  </div>
                </div>
              </div>

              <!-- Release date -->
              <div v-if="versionStore.updateDetails.releaseDate" class="text-center text-sm text-gray-500 dark:text-gray-400">
                Publiée le {{ formatDate(versionStore.updateDetails.releaseDate) }}
              </div>

              <!-- Release notes -->
              <div v-if="versionStore.updateDetails.releaseNotes" class="mt-6">
                <h4 class="text-sm font-semibold text-gray-700 dark:text-gray-300 mb-2">Notes de version</h4>
                <div class="bg-gray-50 dark:bg-gray-900 rounded-lg p-4 prose prose-sm dark:prose-invert max-w-none">
                  <pre class="whitespace-pre-wrap text-sm text-gray-700 dark:text-gray-300">{{ versionStore.updateDetails.releaseNotes }}</pre>
                </div>
              </div>
            </div>
          </div>

          <!-- Footer -->
          <div class="px-6 py-4 border-t border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-900 flex items-center justify-between">
            <button
              @click="dismissUpdate"
              class="btn-ghost text-sm"
            >
              Me le rappeler plus tard
            </button>
            <div class="flex gap-3">
              <button
                @click="showUpdateModal = false"
                class="btn-secondary"
              >
                Fermer
              </button>
              <a
                v-if="versionStore.updateDetails?.releaseURL"
                :href="versionStore.updateDetails.releaseURL"
                target="_blank"
                class="btn-primary inline-flex items-center gap-2"
              >
                <Icon icon="mdi:github" class="h-5 w-5" />
                Voir sur GitHub
              </a>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </aside>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'
import { useVersionStore } from '@/stores/version'
import ZoomControl from '@/components/dashboard/ZoomControl.vue'

const authStore = useAuthStore()
const appStore = useAppStore()
const versionStore = useVersionStore()
const route = useRoute()
const router = useRouter()

// State for update modal
const showUpdateModal = ref(false)

// Load app settings and version when component mounts
onMounted(async () => {
  if (!appStore.appSettings && authStore.isAuthenticated && authStore.isAdmin) {
    try {
      await appStore.refreshAppSettings()
    } catch (error) {
      console.error('Failed to load app settings in sidebar:', error)
    }
  }

  // Load version info
  try {
    await versionStore.fetchVersion()
    // Initialize from cache
    versionStore.initializeFromCache()
    // Start periodic update checks
    versionStore.startPeriodicUpdateCheck()
  } catch (error) {
    console.error('Failed to load version info:', error)
  }
})

// Watch for settings changes
watch(() => appStore.settingsLastUpdated, async () => {
  if (authStore.isAuthenticated && authStore.isAdmin && !appStore.appSettings) {
    try {
      await appStore.refreshAppSettings()
    } catch (error) {
      console.error('Failed to reload app settings:', error)
    }
  }
})

// Functions
const getLinkClasses = (path, exact = false) => {
  const isActive = exact 
    ? route.path === path 
    : route.path.startsWith(path)
  
  return [
    'nav-link',
    isActive ? 'nav-link-active' : 'nav-link-inactive'
  ]
}

const handleLogout = async () => {
  try {
    authStore.logout()
    appStore.showInfo('')
    router.push('/auth/login')
  } catch (error) {
    console.error('Logout error:', error)
    appStore.showError('')
  }
}

const onChangeLocale = (loc) => {
  appStore.setAppLocale(loc)
}

const dismissUpdate = () => {
  versionStore.dismissUpdate()
  showUpdateModal.value = false
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('fr-FR', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}
</script>