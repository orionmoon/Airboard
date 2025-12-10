<template>
  <aside :class="['sidebar', { 'sidebar-open': appStore.sidebarOpen }]">
    <!-- Header -->
    <div class="sidebar-header">
      <div class="flex items-center gap-3">
        <div class="h-8 w-8 bg-white rounded-lg flex items-center justify-center">
          <Icon :icon="appStore.appSettings?.app_icon || 'mdi:view-dashboard'" class="h-5 w-5 text-black" />
        </div>
        <h1 class="sidebar-brand">{{ appStore.appSettings?.app_name || 'Airboard' }}</h1>
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
  </aside>
</template>

<script setup>
import { computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { useAuthStore } from '@/stores/auth'
import { useAppStore } from '@/stores/app'
import ZoomControl from '@/components/dashboard/ZoomControl.vue'

const authStore = useAuthStore()
const appStore = useAppStore()
const route = useRoute()
const router = useRouter()

// Load app settings when component mounts
onMounted(async () => {
  if (!appStore.appSettings && authStore.isAuthenticated && authStore.isAdmin) {
    try {
      await appStore.refreshAppSettings()
    } catch (error) {
      console.error('Failed to load app settings in sidebar:', error)
    }
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
</script>