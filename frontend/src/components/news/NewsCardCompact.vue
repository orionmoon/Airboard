<template>
  <div
    @click="handleClick"
    class="group cursor-pointer p-3 rounded-lg border border-gray-200 dark:border-gray-700 hover:border-primary-300 dark:hover:border-primary-700 hover:shadow-md transition-all duration-200 bg-white dark:bg-gray-800"
  >
    <div class="flex items-center gap-3">
      <!-- Priority & Pin Indicator -->
      <div class="flex-shrink-0 flex flex-col items-center gap-1">
        <div
          v-if="news.priority === 'urgent'"
          class="w-2 h-2 rounded-full bg-red-500 animate-pulse"
          title="Urgent"
        ></div>
        <div
          v-else-if="news.priority === 'important'"
          class="w-2 h-2 rounded-full bg-orange-500"
          title="Important"
        ></div>
        <Icon
          v-if="news.is_pinned"
          icon="mdi:pin"
          class="h-3 w-3 text-yellow-500"
          title="Pinned"
        />
      </div>

      <!-- Content -->
      <div class="flex-1 min-w-0">
        <div class="flex items-start justify-between gap-2 mb-1">
          <h3 class="font-medium text-sm text-gray-900 dark:text-white line-clamp-1 flex-1">
            {{ news.title }}
          </h3>
          <span
            :class="[
              'px-2 py-0.5 rounded text-xs font-medium flex-shrink-0',
              getTypeBadgeClass(news.type)
            ]"
          >
            {{ $t(`news.types.${news.type}`) }}
          </span>
        </div>

        <div class="flex items-center gap-3 text-xs text-gray-500 dark:text-gray-400">
          <span class="flex items-center gap-1">
            <Icon icon="mdi:calendar" class="h-3 w-3" />
            {{ formatDate(news.published_at || news.created_at) }}
          </span>

          <span v-if="news.category" class="flex items-center gap-1">
            <Icon :icon="news.category.icon || 'mdi:folder'" class="h-3 w-3" :style="{ color: news.category.color }" />
            {{ news.category.name }}
          </span>

          <span class="flex items-center gap-1 ml-auto">
            <Icon icon="mdi:eye" class="h-3 w-3" />
            {{ news.views || 0 }}
          </span>

          <span v-if="news.reactions && news.reactions.length > 0" class="flex items-center gap-1">
            <Icon icon="mdi:heart" class="h-3 w-3" />
            {{ news.reactions.length }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const props = defineProps({
  news: {
    type: Object,
    required: true
  }
})

const handleClick = () => {
  router.push({ name: 'NewsDetail', params: { slug: props.news.slug } })
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
</script>
