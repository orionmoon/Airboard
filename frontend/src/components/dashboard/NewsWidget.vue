<template>
  <div class="card">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
        <Icon icon="mdi:newspaper" class="h-5 w-5 text-primary-500" />
        {{ $t('news.widget.title') }}
        <span v-if="unreadCount > 0" class="ml-2 px-2 py-1 bg-red-500 text-white text-xs rounded-full">
          {{ unreadCount }}
        </span>
      </h2>
      <router-link
        to="/news"
        class="text-sm text-primary-500 hover:text-primary-600 font-medium flex items-center gap-1"
      >
        {{ $t('news.widget.viewAll') }}
        <Icon icon="mdi:arrow-right" class="h-4 w-4" />
      </router-link>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-8">
      <Icon icon="mdi:loading" class="h-6 w-6 animate-spin text-primary-500" />
    </div>

    <!-- News List -->
    <div v-else-if="newsList.length > 0" class="space-y-3">
      <div
        v-for="news in newsList"
        :key="news.id"
        @click="viewNews(news)"
        class="group cursor-pointer p-3 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-primary-300 dark:hover:border-primary-700 hover:shadow-sm transition-all"
      >
        <!-- Header -->
        <div class="flex items-start justify-between mb-2">
          <div class="flex items-center gap-2 flex-1 min-w-0">
            <!-- Priority Indicator -->
            <div
              v-if="news.priority === 'urgent'"
              class="flex-shrink-0 w-2 h-2 rounded-full bg-red-500 animate-pulse"
            ></div>
            <div
              v-else-if="news.priority === 'important'"
              class="flex-shrink-0 w-2 h-2 rounded-full bg-orange-500"
            ></div>

            <!-- Title -->
            <h3 class="font-medium text-gray-900 dark:text-white text-sm line-clamp-1 flex-1">
              {{ news.title }}
            </h3>

            <!-- Pinned Icon -->
            <Icon
              v-if="news.is_pinned"
              icon="mdi:pin"
              class="h-4 w-4 text-yellow-500 flex-shrink-0"
            />
          </div>

          <!-- Type Badge -->
          <span
            :class="[
              'ml-2 px-2 py-0.5 rounded text-xs font-medium flex-shrink-0',
              getTypeBadgeClass(news.type)
            ]"
          >
            {{ $t(`news.types.${news.type}`) }}
          </span>
        </div>

        <!-- Summary -->
        <p class="text-sm text-gray-600 dark:text-gray-400 line-clamp-2 mb-2">
          {{ news.summary }}
        </p>

        <!-- Footer -->
        <div class="flex items-center justify-between text-xs text-gray-500 dark:text-gray-500">
          <div class="flex items-center gap-3">
            <span class="flex items-center gap-1">
              <Icon icon="mdi:calendar" class="h-3 w-3" />
              {{ formatDate(news.published_at || news.created_at) }}
            </span>
            <span v-if="news.category" class="flex items-center gap-1">
              <Icon :icon="news.category.icon || 'mdi:folder'" class="h-3 w-3" :style="{ color: news.category.color }" />
              {{ news.category.name }}
            </span>
          </div>

          <!-- Reactions count -->
          <div v-if="news.reactions && news.reactions.length > 0" class="flex items-center gap-1">
            <Icon icon="mdi:heart" class="h-3 w-3" />
            <span>{{ news.reactions.length }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-8">
      <Icon icon="mdi:newspaper-variant-outline" class="h-12 w-12 mx-auto text-gray-400 mb-2" />
      <p class="text-sm text-gray-600 dark:text-gray-400">
        {{ $t('news.widget.noNews') }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { newsService } from '@/services/api'

const router = useRouter()

// Data
const loading = ref(false)
const newsList = ref([])
const unreadCount = ref(0)

// Methods
const fetchTopNews = async () => {
  loading.value = true
  try {
    const response = await newsService.getNews({
      limit: 3,
      is_pinned: true // Get pinned news first, or latest
    })
    newsList.value = response.news || []
  } catch (error) {
    console.error('Error fetching top news:', error)
  } finally {
    loading.value = false
  }
}

const fetchUnreadCount = async () => {
  try {
    const response = await newsService.getUnreadCount()
    unreadCount.value = response.count || 0
  } catch (error) {
    console.error('Error fetching unread count:', error)
  }
}

const viewNews = (news) => {
  router.push({ name: 'NewsDetail', params: { slug: news.slug } })
}

const getTypeBadgeClass = (type) => {
  const classes = {
    article: 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300',
    tutorial: 'bg-green-100 text-green-700 dark:bg-green-900/30 dark:text-green-300',
    announcement: 'bg-purple-100 text-purple-700 dark:bg-purple-900/30 dark:text-purple-300',
    faq: 'bg-yellow-100 text-yellow-700 dark:bg-yellow-900/30 dark:text-yellow-300'
  }
  return classes[type] || classes.article
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  const now = new Date()
  const diffInDays = Math.floor((now - date) / (1000 * 60 * 60 * 24))

  if (diffInDays === 0) return 'Aujourd\'hui'
  if (diffInDays === 1) return 'Hier'
  if (diffInDays < 7) return `Il y a ${diffInDays}j`

  return date.toLocaleDateString('fr-FR', { day: 'numeric', month: 'short' })
}

// Lifecycle
onMounted(() => {
  fetchTopNews()
  fetchUnreadCount()
})
</script>
