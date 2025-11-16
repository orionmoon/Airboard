<template>
  <div
    class="bg-white dark:bg-gray-800 rounded-lg shadow-sm hover:shadow-md transition-shadow cursor-pointer overflow-hidden"
    @click="$emit('click')"
  >
    <div class="p-6">
      <!-- Header with Type Badge and Priority -->
      <div class="flex items-start justify-between mb-3">
        <div class="flex items-center gap-2">
          <!-- Type Badge -->
          <span
            :class="[
              'px-3 py-1 rounded-full text-xs font-medium',
              getTypeBadgeClass(news.type)
            ]"
          >
            {{ $t(`news.types.${news.type}`) }}
          </span>

          <!-- Priority Badge -->
          <span
            v-if="news.priority === 'urgent'"
            class="px-3 py-1 rounded-full text-xs font-medium bg-red-100 text-red-700 dark:bg-red-900 dark:text-red-300"
          >
            <Icon icon="mdi:alert" class="inline h-3 w-3 mr-1" />
            {{ $t('news.priority.urgent') }}
          </span>
          <span
            v-else-if="news.priority === 'important'"
            class="px-3 py-1 rounded-full text-xs font-medium bg-orange-100 text-orange-700 dark:bg-orange-900 dark:text-orange-300"
          >
            <Icon icon="mdi:star" class="inline h-3 w-3 mr-1" />
            {{ $t('news.priority.important') }}
          </span>
        </div>

        <!-- Category -->
        <div v-if="news.category" class="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400">
          <Icon :icon="news.category.icon || 'mdi:folder'" class="h-4 w-4" :style="{ color: news.category.color }" />
          <span>{{ news.category.name }}</span>
        </div>
      </div>

      <!-- Title -->
      <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-2 line-clamp-2">
        {{ news.title }}
      </h3>

      <!-- Summary -->
      <p class="text-gray-600 dark:text-gray-400 mb-4 line-clamp-3">
        {{ news.summary }}
      </p>

      <!-- Tags -->
      <div v-if="news.tags && news.tags.length > 0" class="flex flex-wrap gap-2 mb-4">
        <span
          v-for="tag in news.tags.slice(0, 5)"
          :key="tag.id"
          class="px-2 py-1 rounded text-xs font-medium bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300"
          :style="{ backgroundColor: tag.color + '20', color: tag.color }"
        >
          #{{ tag.name }}
        </span>
        <span
          v-if="news.tags.length > 5"
          class="px-2 py-1 rounded text-xs font-medium text-gray-500 dark:text-gray-400"
        >
          +{{ news.tags.length - 5 }}
        </span>
      </div>

      <!-- Footer with Metadata and Reactions -->
      <div class="flex items-center justify-between pt-4 border-t border-gray-200 dark:border-gray-700">
        <!-- Author and Date -->
        <div class="flex items-center gap-4 text-sm text-gray-600 dark:text-gray-400">
          <div class="flex items-center gap-1">
            <Icon icon="mdi:account" class="h-4 w-4" />
            <span>{{ news.author?.full_name || news.author?.email }}</span>
          </div>
          <div class="flex items-center gap-1">
            <Icon icon="mdi:calendar" class="h-4 w-4" />
            <span>{{ formatDate(news.published_at || news.created_at) }}</span>
          </div>
        </div>

        <!-- Reactions -->
        <div class="flex items-center gap-3">
          <!-- View Count -->
          <div class="flex items-center gap-1 text-sm text-gray-600 dark:text-gray-400">
            <Icon icon="mdi:eye" class="h-4 w-4" />
            <span>{{ news.view_count || 0 }}</span>
          </div>

          <!-- Reaction Counts -->
          <div v-if="news.reactions && news.reactions.length > 0" class="flex items-center gap-2">
            <button
              v-for="reaction in aggregateReactions(news.reactions)"
              :key="reaction.type"
              class="flex items-center gap-1 text-sm hover:scale-110 transition-transform"
              @click.stop="handleReaction(reaction.type)"
            >
              <span class="text-lg">{{ getReactionEmoji(reaction.type) }}</span>
              <span class="text-gray-600 dark:text-gray-400">{{ reaction.count }}</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'

const props = defineProps({
  news: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['click'])

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
  const now = new Date()
  const diffInDays = Math.floor((now - date) / (1000 * 60 * 60 * 24))

  if (diffInDays === 0) return 'Aujourd\'hui'
  if (diffInDays === 1) return 'Hier'
  if (diffInDays < 7) return `Il y a ${diffInDays} jours`
  if (diffInDays < 30) return `Il y a ${Math.floor(diffInDays / 7)} semaines`

  return date.toLocaleDateString('fr-FR', { day: 'numeric', month: 'short', year: 'numeric' })
}

const aggregateReactions = (reactions) => {
  if (!reactions) return []

  const counts = {}
  reactions.forEach(r => {
    counts[r.reaction_type] = (counts[r.reaction_type] || 0) + 1
  })

  return Object.entries(counts).map(([type, count]) => ({ type, count }))
}

const getReactionEmoji = (type) => {
  const emojis = {
    thumbs_up: '👍',
    heart: '❤️',
    celebrate: '🎉'
  }
  return emojis[type] || '👍'
}

const handleReaction = (type) => {
  // This will be handled by the parent component or NewsDetail page
  console.log('Reaction clicked:', type)
}
</script>
