<template>
  <div class="table-view-container">
    <!-- Filters Bar -->
    <div class="filters-bar">
      <div class="flex flex-col sm:flex-row gap-4 mb-6">
        <!-- Search Input -->
        <div class="flex-1">
          <div class="relative">
            <Icon icon="mdi:magnify" class="absolute left-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
            <input
              v-model="filters.search"
              type="text"
              :placeholder="$t('dashboard.searchApps')"
              class="w-full pl-10 pr-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            />
          </div>
        </div>

        <!-- Category Filter -->
        <div class="w-full sm:w-64">
          <select
            v-model="filters.category"
            class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="">{{ $t('dashboard.allCategories') }}</option>
            <option v-for="category in availableCategories" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
        </div>

        <!-- Sort By -->
        <div class="w-full sm:w-48">
          <select
            v-model="filters.sortBy"
            class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg bg-white dark:bg-gray-800 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          >
            <option value="name">{{ $t('dashboard.sortByName') }}</option>
            <option value="category">{{ $t('dashboard.sortByCategory') }}</option>
            <option value="recent">{{ $t('dashboard.sortByRecent') }}</option>
          </select>
        </div>

        <!-- Reset Filters Button -->
        <button
          v-if="hasActiveFilters"
          @click="resetFilters"
          class="px-4 py-2 bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 rounded-lg text-gray-700 dark:text-gray-300 transition-colors flex items-center gap-2"
        >
          <Icon icon="mdi:filter-off" class="h-5 w-5" />
          <span class="hidden sm:inline">{{ $t('dashboard.resetFilters') }}</span>
        </button>
      </div>
    </div>

    <!-- Results Count -->
    <div class="mb-4 text-sm text-gray-600 dark:text-gray-400">
      {{ filteredApplications.length }} {{ $t('dashboard.applicationsFound') }}
    </div>

    <!-- Table -->
    <div class="overflow-x-auto bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 shadow-sm">
      <table class="w-full">
        <thead class="bg-gray-50 dark:bg-gray-900 border-b border-gray-200 dark:border-gray-700">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-700 dark:text-gray-300 uppercase tracking-wider">
              <Icon icon="mdi:star" class="h-4 w-4 inline" />
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-700 dark:text-gray-300 uppercase tracking-wider">
              {{ $t('dashboard.application') }}
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-700 dark:text-gray-300 uppercase tracking-wider">
              {{ $t('dashboard.description') }}
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-700 dark:text-gray-300 uppercase tracking-wider">
              {{ $t('dashboard.category') }}
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-700 dark:text-gray-300 uppercase tracking-wider">
              {{ $t('dashboard.actions') }}
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200 dark:divide-gray-700">
          <tr
            v-for="app in filteredApplications"
            :key="app.id"
            class="hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors cursor-pointer"
            @click="openApplication(app)"
          >
            <!-- Favorite -->
            <td class="px-6 py-4 whitespace-nowrap">
              <button
                @click.stop="toggleFavorite(app)"
                class="p-1 hover:scale-110 transition-transform"
              >
                <Icon
                  :icon="isFavorite(app.id) ? 'mdi:star' : 'mdi:star-outline'"
                  class="h-5 w-5"
                  :class="isFavorite(app.id) ? 'text-yellow-500' : 'text-gray-400'"
                />
              </button>
            </td>

            <!-- Application Name & Icon -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center gap-3">
                <div
                  class="h-10 w-10 rounded-lg flex items-center justify-center flex-shrink-0"
                  :style="{ backgroundColor: app.color || '#6366f1' }"
                >
                  <Icon :icon="app.icon || 'mdi:application'" class="h-5 w-5 text-white" />
                </div>
                <div class="flex items-center gap-2">
                  <span class="font-medium text-gray-900 dark:text-white">{{ app.name }}</span>
                  <Icon v-if="app.open_in_new_tab" icon="mdi:open-in-new" class="h-4 w-4 text-gray-500 dark:text-gray-400" />
                </div>
              </div>
            </td>

            <!-- Description -->
            <td class="px-6 py-4">
              <span class="text-sm text-gray-600 dark:text-gray-400 line-clamp-2">
                {{ app.description || '-' }}
              </span>
            </td>

            <!-- Category -->
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center gap-2">
                <div
                  class="h-6 w-6 rounded flex items-center justify-center flex-shrink-0"
                  :style="{ backgroundColor: app.app_group?.color || '#10b981' }"
                >
                  <Icon :icon="app.app_group?.icon || 'mdi:folder'" class="h-3 w-3 text-white" />
                </div>
                <span class="text-sm text-gray-700 dark:text-gray-300">
                  {{ app.app_group?.name || '-' }}
                </span>
              </div>
            </td>

            <!-- Actions -->
            <td class="px-6 py-4 whitespace-nowrap">
              <button
                @click.stop="openApplication(app)"
                class="inline-flex items-center gap-2 px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium rounded-lg transition-colors"
              >
                <Icon icon="mdi:open-in-new" class="h-4 w-4" />
                {{ $t('dashboard.open') }}
              </button>
            </td>
          </tr>

          <!-- Empty State -->
          <tr v-if="filteredApplications.length === 0">
            <td colspan="5" class="px-6 py-12 text-center">
              <Icon icon="mdi:application-off" class="h-12 w-12 mx-auto text-gray-400 mb-2" />
              <p class="text-gray-500 dark:text-gray-400">{{ $t('dashboard.noApplicationsFound') }}</p>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Icon } from '@iconify/vue'
import { useFavoritesStore } from '@/stores/favorites'
import { analyticsService } from '@/services/api'

const props = defineProps({
  appGroups: {
    type: Array,
    required: true
  }
})

// No emits needed - favorites store handles reactivity

const favoritesStore = useFavoritesStore()

// Filters state
const filters = ref({
  search: '',
  category: '',
  sortBy: 'name'
})

// Computed: All applications from all groups
const allApplications = computed(() => {
  const apps = []
  props.appGroups.forEach(group => {
    group.applications?.forEach(app => {
      apps.push({
        ...app,
        app_group: {
          id: group.id,
          name: group.name,
          icon: group.icon,
          color: group.color
        }
      })
    })
  })
  return apps
})

// Computed: Available categories
const availableCategories = computed(() => {
  return props.appGroups || []
})

// Computed: Check if filters are active
const hasActiveFilters = computed(() => {
  return filters.value.search !== '' ||
         filters.value.category !== '' ||
         filters.value.sortBy !== 'name'
})

// Computed: Filtered applications
const filteredApplications = computed(() => {
  let apps = [...allApplications.value]

  // Filter by search
  if (filters.value.search) {
    const search = filters.value.search.toLowerCase()
    apps = apps.filter(app =>
      app.name.toLowerCase().includes(search) ||
      app.description?.toLowerCase().includes(search)
    )
  }

  // Filter by category
  if (filters.value.category) {
    apps = apps.filter(app => app.app_group.id === parseInt(filters.value.category))
  }

  // Sort
  switch (filters.value.sortBy) {
    case 'name':
      apps.sort((a, b) => a.name.localeCompare(b.name))
      break
    case 'category':
      apps.sort((a, b) => a.app_group.name.localeCompare(b.app_group.name))
      break
    case 'recent':
      // If you have a created_at or updated_at field, sort by that
      // For now, reverse the order (most recent first)
      apps.reverse()
      break
  }

  return apps
})

// Methods
const resetFilters = () => {
  filters.value = {
    search: '',
    category: '',
    sortBy: 'name'
  }
}

const isFavorite = (appId) => {
  return favoritesStore.isFavorite(appId)
}

const toggleFavorite = async (app) => {
  try {
    await favoritesStore.toggleFavorite(app.id)
  } catch (error) {
    console.error('Error toggling favorite:', error)
  }
}

const openApplication = async (app) => {
  // Track click for analytics
  analyticsService.trackClick(app.id)

  // Open application
  if (app.open_in_new_tab) {
    window.open(app.url, '_blank', 'noopener,noreferrer')
  } else {
    window.location.href = app.url
  }

  console.log(`Application opened: ${app.name} - ${app.url}`)
}

// Save filters to localStorage
watch(filters, (newFilters) => {
  localStorage.setItem('dashboard-table-filters', JSON.stringify(newFilters))
}, { deep: true })

// Load filters from localStorage
const loadFilters = () => {
  const saved = localStorage.getItem('dashboard-table-filters')
  if (saved) {
    try {
      filters.value = JSON.parse(saved)
    } catch (error) {
      console.error('Error loading filters:', error)
    }
  }
}

// Load filters on mount
loadFilters()
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
