<template>
  <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-12">
      <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-primary-500" />
    </div>

    <!-- Article Content -->
    <article v-else-if="news" class="bg-white dark:bg-gray-800 rounded-lg shadow-sm overflow-hidden">
      <!-- Header -->
      <div class="p-8 border-b border-gray-200 dark:border-gray-700">
        <!-- Breadcrumb -->
        <div class="mb-4">
          <button
            @click="$router.back()"
            class="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400 hover:text-primary-500"
          >
            <Icon icon="mdi:arrow-left" class="h-4 w-4" />
            {{ $t('news.detail.back') }}
          </button>
        </div>

        <!-- Badges -->
        <div class="flex flex-wrap items-center gap-2 mb-4">
          <!-- Type Badge -->
          <span
            :class="[
              'px-3 py-1 rounded-full text-sm font-medium',
              getTypeBadgeClass(news.type)
            ]"
          >
            {{ $t(`news.types.${news.type}`) }}
          </span>

          <!-- Priority Badge -->
          <span
            v-if="news.priority === 'urgent'"
            class="px-3 py-1 rounded-full text-sm font-medium bg-red-100 text-red-700 dark:bg-red-900 dark:text-red-300"
          >
            <Icon icon="mdi:alert" class="inline h-4 w-4 mr-1" />
            {{ $t('news.priority.urgent') }}
          </span>
          <span
            v-else-if="news.priority === 'important'"
            class="px-3 py-1 rounded-full text-sm font-medium bg-orange-100 text-orange-700 dark:bg-orange-900 dark:text-orange-300"
          >
            <Icon icon="mdi:star" class="inline h-4 w-4 mr-1" />
            {{ $t('news.priority.important') }}
          </span>

          <!-- Pinned Badge -->
          <span
            v-if="news.is_pinned"
            class="px-3 py-1 rounded-full text-sm font-medium bg-yellow-100 text-yellow-700 dark:bg-yellow-900 dark:text-yellow-300"
          >
            <Icon icon="mdi:pin" class="inline h-4 w-4 mr-1" />
            {{ $t('news.detail.pinned') }}
          </span>

          <!-- Category -->
          <span
            v-if="news.category"
            class="px-3 py-1 rounded-full text-sm font-medium bg-gray-100 text-gray-700 dark:bg-gray-700 dark:text-gray-300"
          >
            <Icon :icon="news.category.icon || 'mdi:folder'" class="inline h-4 w-4 mr-1" :style="{ color: news.category.color }" />
            {{ news.category.name }}
          </span>
        </div>

        <!-- Title -->
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
          {{ news.title }}
        </h1>

        <!-- Summary -->
        <p class="text-xl text-gray-600 dark:text-gray-400 mb-6">
          {{ news.summary }}
        </p>

        <!-- Metadata -->
        <div class="flex flex-wrap items-center gap-6 text-sm text-gray-600 dark:text-gray-400">
          <div class="flex items-center gap-2">
            <Icon icon="mdi:account" class="h-5 w-5" />
            <span>{{ news.author?.full_name || news.author?.email }}</span>
          </div>
          <div class="flex items-center gap-2">
            <Icon icon="mdi:calendar" class="h-5 w-5" />
            <span>{{ formatDate(news.published_at || news.created_at) }}</span>
          </div>
          <div class="flex items-center gap-2">
            <Icon icon="mdi:eye" class="h-5 w-5" />
            <span>{{ news.view_count || 0 }} {{ $t('news.detail.views') }}</span>
          </div>
        </div>

        <!-- Tags -->
        <div v-if="news.tags && news.tags.length > 0" class="flex flex-wrap gap-2 mt-4">
          <span
            v-for="tag in news.tags"
            :key="tag.id"
            class="px-3 py-1 rounded-full text-sm font-medium bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300"
            :style="{ backgroundColor: tag.color + '20', color: tag.color }"
          >
            #{{ tag.name }}
          </span>
        </div>
      </div>

      <!-- Content -->
      <div class="p-8">
        <div class="prose prose-lg dark:prose-invert max-w-none">
          <TiptapRenderer :content="news.content" />
        </div>
      </div>

      <!-- Reactions Section -->
      <div class="p-8 border-t border-gray-200 dark:border-gray-700">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">
          {{ $t('news.detail.reactions') }}
        </h3>

        <div class="flex flex-wrap gap-3">
          <button
            v-for="reactionType in availableReactions"
            :key="reactionType"
            @click="toggleReaction(reactionType)"
            :class="[
              'flex items-center gap-2 px-4 py-2 rounded-lg border-2 transition-all',
              hasUserReacted(reactionType)
                ? 'border-primary-500 bg-primary-50 dark:bg-primary-900/20'
                : 'border-gray-300 dark:border-gray-600 hover:border-primary-300 dark:hover:border-primary-700'
            ]"
          >
            <span class="text-2xl">{{ getReactionEmoji(reactionType) }}</span>
            <span class="text-sm font-medium text-gray-700 dark:text-gray-300">
              {{ getReactionCount(reactionType) }}
            </span>
          </button>
        </div>
      </div>

      <!-- Edit Button (for editors/admins) -->
      <div v-if="canEdit" class="p-8 border-t border-gray-200 dark:border-gray-700">
        <button
          @click="editNews"
          class="px-4 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors flex items-center gap-2"
        >
          <Icon icon="mdi:pencil" class="h-5 w-5" />
          {{ $t('news.detail.edit') }}
        </button>
      </div>
    </article>

    <!-- Error State -->
    <div v-else class="text-center py-12">
      <Icon icon="mdi:alert-circle" class="h-16 w-16 mx-auto text-red-500 mb-4" />
      <p class="text-gray-600 dark:text-gray-400 text-lg">
        {{ $t('news.detail.notFound') }}
      </p>
      <button
        @click="$router.push({ name: 'NewsCenter' })"
        class="mt-4 px-6 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors"
      >
        {{ $t('news.detail.backToList') }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { newsService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'
import TiptapRenderer from '@/components/news/TiptapRenderer.vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

// Data
const loading = ref(false)
const news = ref(null)
const reactionCounts = ref({})
const userReaction = ref(null)

const availableReactions = ['thumbs_up', 'heart', 'celebrate']

// Computed
const canEdit = computed(() => {
  const userRole = authStore.user?.role
  return userRole === 'admin' || userRole === 'editor'
})

// Methods
const fetchNews = async () => {
  loading.value = true
  try {
    const slug = route.params.slug
    news.value = await newsService.getNewsBySlug(slug)

    // Increment view count
    await newsService.incrementView(news.value.id)

    // Fetch reactions
    await fetchReactions()
  } catch (error) {
    console.error('Error fetching news:', error)
    news.value = null
  } finally {
    loading.value = false
  }
}

const fetchReactions = async () => {
  try {
    const response = await newsService.getReactions(news.value.id)
    reactionCounts.value = response.counts || {}
    // user_reaction est un objet avec reaction_type, on extrait juste le type
    userReaction.value = response.user_reaction?.reaction_type || null
  } catch (error) {
    console.error('Error fetching reactions:', error)
  }
}

const toggleReaction = async (reactionType) => {
  try {
    if (hasUserReacted(reactionType)) {
      // Remove reaction
      await newsService.removeReaction(news.value.id)
      userReaction.value = null
    } else {
      // Add or change reaction
      await newsService.addReaction(news.value.id, reactionType)
      userReaction.value = reactionType
    }

    // Refresh reactions
    await fetchReactions()
  } catch (error) {
    console.error('Error toggling reaction:', error)
  }
}

const hasUserReacted = (reactionType) => {
  return userReaction.value === reactionType
}

const getReactionCount = (reactionType) => {
  return reactionCounts.value[reactionType] || 0
}

const getReactionEmoji = (type) => {
  const emojis = {
    thumbs_up: '👍',
    heart: '❤️',
    celebrate: '🎉'
  }
  return emojis[type] || '👍'
}

const getTypeBadgeClass = (type) => {
  const classes = {
    article: 'bg-blue-100 text-blue-700 dark:bg-blue-900 dark:text-blue-300',
    tutorial: 'bg-green-100 text-green-700 dark:bg-green-900 dark:text-green-300',
    announcement: 'bg-purple-100 text-purple-700 dark:bg-purple-900 dark:text-purple-300',
    faq: 'bg-yellow-100 text-yellow-700 dark:bg-yellow-900 dark:text-yellow-300'
  }
  return classes[type] || classes.article
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  return date.toLocaleDateString('fr-FR', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const editNews = () => {
  router.push({ name: 'AdminNewsEdit', params: { id: news.value.id } })
}

// Lifecycle
onMounted(() => {
  fetchNews()
})
</script>
