<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-6">
    <div class="w-full">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
          {{ $t('analytics.title') }}
        </h1>
        <p class="mt-1 text-gray-600 dark:text-gray-400">
          {{ $t('analytics.subtitle') }}
        </p>
      </div>

      <!-- Loading state -->
      <div v-if="isLoading" class="flex justify-center py-12">
        <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-gray-400" />
      </div>

      <!-- Analytics Dashboard -->
      <div v-else-if="analytics">
        <!-- KPI Cards -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <!-- Total Clicks -->
          <div class="card p-6">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="h-12 w-12 bg-blue-100 dark:bg-blue-900 rounded-lg flex items-center justify-center">
                  <Icon icon="mdi:cursor-default-click" class="h-6 w-6 text-blue-600 dark:text-blue-400" />
                </div>
              </div>
              <div class="ml-4">
                <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('analytics.totalClicks') }}</p>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">
                  {{ analytics.total_clicks.toLocaleString() }}
                </p>
              </div>
            </div>
          </div>

          <!-- Unique Users -->
          <div class="card p-6">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="h-12 w-12 bg-green-100 dark:bg-green-900 rounded-lg flex items-center justify-center">
                  <Icon icon="mdi:account-multiple" class="h-6 w-6 text-green-600 dark:text-green-400" />
                </div>
              </div>
              <div class="ml-4">
                <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('analytics.uniqueUsers') }}</p>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">
                  {{ analytics.total_unique_users }}
                </p>
              </div>
            </div>
          </div>

          <!-- Clicks Last 7 Days -->
          <div class="card p-6">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="h-12 w-12 bg-purple-100 dark:bg-purple-900 rounded-lg flex items-center justify-center">
                  <Icon icon="mdi:calendar-week" class="h-6 w-6 text-purple-600 dark:text-purple-400" />
                </div>
              </div>
              <div class="ml-4">
                <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('analytics.last7Days') }}</p>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">
                  {{ analytics.clicks_last_7_days.toLocaleString() }}
                </p>
              </div>
            </div>
          </div>

          <!-- Clicks Last 30 Days -->
          <div class="card p-6">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="h-12 w-12 bg-orange-100 dark:bg-orange-900 rounded-lg flex items-center justify-center">
                  <Icon icon="mdi:calendar-month" class="h-6 w-6 text-orange-600 dark:text-orange-400" />
                </div>
              </div>
              <div class="ml-4">
                <p class="text-sm font-medium text-gray-600 dark:text-gray-400">{{ $t('analytics.last30Days') }}</p>
                <p class="text-2xl font-bold text-gray-900 dark:text-white">
                  {{ analytics.clicks_last_30_days.toLocaleString() }}
                </p>
              </div>
            </div>
          </div>
        </div>

        <!-- Charts Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
          <!-- Top Applications -->
          <div class="card p-6">
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-6">
              {{ $t('analytics.topApplications') }}
            </h2>
            <div v-if="analytics.top_applications && analytics.top_applications.length > 0" class="space-y-4">
              <div
                v-for="(app, index) in analytics.top_applications"
                :key="app.application_id"
                class="flex items-center gap-4"
              >
                <div class="flex-shrink-0 w-8 text-center">
                  <span class="text-lg font-bold text-gray-400">#{{ index + 1 }}</span>
                </div>
                <div
                  class="h-10 w-10 rounded-lg flex items-center justify-center flex-shrink-0"
                  :style="{ backgroundColor: app.color || '#6366f1' }"
                >
                  <Icon :icon="app.icon || 'mdi:application'" class="h-5 w-5 text-white" />
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-medium text-gray-900 dark:text-white truncate">{{ app.application_name }}</p>
                  <p class="text-xs text-gray-500 dark:text-gray-400">
                    {{ app.click_count }} {{ $t('analytics.clicks') }} • {{ app.unique_users }} {{ $t('analytics.users') }}
                  </p>
                </div>
                <div class="flex-shrink-0">
                  <div class="w-32 bg-gray-200 dark:bg-gray-700 rounded-full h-2">
                    <div
                      class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full transition-all duration-500"
                      :style="{ width: (app.click_count / analytics.top_applications[0].click_count * 100) + '%' }"
                    ></div>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="text-center py-8 text-gray-500 dark:text-gray-400">
              {{ $t('analytics.noData') }}
            </div>
          </div>

          <!-- Top Users -->
          <div class="card p-6">
            <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-6">
              {{ $t('analytics.topUsers') }}
            </h2>
            <div v-if="analytics.top_users && analytics.top_users.length > 0" class="space-y-4">
              <div
                v-for="(user, index) in analytics.top_users"
                :key="user.user_id"
                class="flex items-center gap-4"
              >
                <div class="flex-shrink-0 w-8 text-center">
                  <span class="text-lg font-bold text-gray-400">#{{ index + 1 }}</span>
                </div>
                <div class="h-10 w-10 rounded-full bg-gradient-to-br from-green-400 to-blue-500 flex items-center justify-center flex-shrink-0">
                  <span class="text-sm font-bold text-white">{{ getUserInitials(user) }}</span>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-medium text-gray-900 dark:text-white truncate">
                    {{ user.first_name }} {{ user.last_name }}
                  </p>
                  <p class="text-xs text-gray-500 dark:text-gray-400">
                    {{ user.click_count }} {{ $t('analytics.clicks') }} • {{ user.unique_apps }} {{ $t('analytics.apps') }}
                  </p>
                </div>
                <div class="flex-shrink-0">
                  <div class="w-32 bg-gray-200 dark:bg-gray-700 rounded-full h-2">
                    <div
                      class="bg-gradient-to-r from-green-500 to-blue-500 h-2 rounded-full transition-all duration-500"
                      :style="{ width: (user.click_count / analytics.top_users[0].click_count * 100) + '%' }"
                    ></div>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="text-center py-8 text-gray-500 dark:text-gray-400">
              {{ $t('analytics.noData') }}
            </div>
          </div>
        </div>

        <!-- News Hub Statistics -->
        <div v-if="newsStats" class="card p-6 mb-8">
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-6">
            📰 Statistiques News Hub
          </h2>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div class="text-center">
              <p class="text-3xl font-bold text-blue-600 dark:text-blue-400">{{ newsStats.total_news }}</p>
              <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Articles totaux</p>
            </div>
            <div class="text-center">
              <p class="text-3xl font-bold text-green-600 dark:text-green-400">{{ newsStats.published_news }}</p>
              <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Publiés</p>
            </div>
            <div class="text-center">
              <p class="text-3xl font-bold text-purple-600 dark:text-purple-400">{{ newsStats.total_views }}</p>
              <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Vues totales</p>
            </div>
            <div class="text-center">
              <p class="text-3xl font-bold text-red-600 dark:text-red-400">{{ newsStats.total_reactions }}</p>
              <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">Réactions totales</p>
            </div>
          </div>

          <!-- Top News -->
          <div v-if="newsStats.top_news && newsStats.top_news.length > 0" class="mt-6">
            <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Articles les plus populaires</h3>
            <div class="space-y-3">
              <div
                v-for="(article, index) in newsStats.top_news"
                :key="article.id"
                class="flex items-center gap-4 p-3 bg-gray-50 dark:bg-gray-700 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors cursor-pointer"
                @click="viewArticle(article.slug)"
              >
                <div class="flex-shrink-0 w-8 text-center">
                  <span class="text-lg font-bold text-gray-400">#{{ index + 1 }}</span>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-medium text-gray-900 dark:text-white truncate">{{ article.title }}</p>
                  <div class="flex gap-3 mt-1">
                    <span class="text-xs text-gray-500 dark:text-gray-400">
                      <Icon icon="mdi:eye" class="inline h-3 w-3" /> {{ article.view_count }}
                    </span>
                    <span class="text-xs text-gray-500 dark:text-gray-400">
                      <Icon icon="mdi:heart" class="inline h-3 w-3" /> {{ article.reaction_count }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Daily Activity Chart -->
        <div class="card p-6">
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-6">
            {{ $t('analytics.dailyActivity') }}
          </h2>
          <div v-if="analytics.daily_activity && analytics.daily_activity.length > 0" class="space-y-4">
            <!-- Simple Bar Chart -->
            <div class="flex items-end justify-between gap-2 h-64">
              <div
                v-for="day in analytics.daily_activity"
                :key="day.date"
                class="flex-1 flex flex-col items-center gap-2"
              >
                <div class="w-full flex items-end justify-center h-full">
                  <div
                    class="w-full bg-gradient-to-t from-blue-500 to-purple-500 rounded-t-lg transition-all duration-500 hover:opacity-80 relative group"
                    :style="{ height: getBarHeight(day.click_count) + '%' }"
                    :title="`${day.date}: ${day.click_count} clicks`"
                  >
                    <div class="absolute -top-8 left-1/2 transform -translate-x-1/2 opacity-0 group-hover:opacity-100 transition-opacity bg-gray-900 dark:bg-gray-700 text-white text-xs px-2 py-1 rounded whitespace-nowrap">
                      {{ day.click_count }} {{ $t('analytics.clicks') }}
                    </div>
                  </div>
                </div>
                <span class="text-xs text-gray-500 dark:text-gray-400 rotate-45 origin-left whitespace-nowrap">
                  {{ formatDate(day.date) }}
                </span>
              </div>
            </div>
          </div>
          <div v-else class="text-center py-12 text-gray-500 dark:text-gray-400">
            {{ $t('analytics.noData') }}
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="empty-state">
        <Icon icon="mdi:chart-bar" class="empty-state-icon" />
        <h3 class="empty-state-title">{{ $t('analytics.noAnalytics') }}</h3>
        <p class="empty-state-description">{{ $t('analytics.noAnalyticsText') }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { analyticsService, newsService } from '@/services/api'

const router = useRouter()
const analytics = ref(null)
const newsStats = ref(null)
const isLoading = ref(false)

const loadAnalytics = async () => {
  try {
    isLoading.value = true
    analytics.value = await analyticsService.getDashboard()
  } catch (error) {
    console.error('Error loading analytics:', error)
  } finally {
    isLoading.value = false
  }
}

const loadNewsStats = async () => {
  try {
    newsStats.value = await newsService.getAnalytics()
  } catch (error) {
    console.error('Error loading news stats:', error)
  }
}

const getUserInitials = (user) => {
  const first = user.first_name?.[0] || ''
  const last = user.last_name?.[0] || user.username?.[0] || '?'
  return (first + last).toUpperCase()
}

const getBarHeight = (clickCount) => {
  if (!analytics.value?.daily_activity || analytics.value.daily_activity.length === 0) return 0
  const maxClicks = Math.max(...analytics.value.daily_activity.map(d => d.click_count))
  if (maxClicks === 0) return 0
  return Math.max((clickCount / maxClicks * 100), 5) // Minimum 5% for visibility
}

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('fr-FR', { day: '2-digit', month: 'short' })
}

const viewArticle = (slug) => {
  router.push({ name: 'NewsDetail', params: { slug } })
}

onMounted(() => {
  loadAnalytics()
  loadNewsStats()
})
</script>

<style scoped>
.card {
  @apply bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700;
}

.empty-state {
  @apply flex flex-col items-center justify-center py-12;
}

.empty-state-icon {
  @apply h-16 w-16 text-gray-400 dark:text-gray-600 mb-4;
}

.empty-state-title {
  @apply text-xl font-semibold text-gray-900 dark:text-white mb-2;
}

.empty-state-description {
  @apply text-gray-600 dark:text-gray-400 text-center max-w-md;
}
</style>
