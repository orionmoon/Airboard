<template>
  <div class="w-full p-6">
    <!-- Header -->
    <div class="mb-8">
      <div class="flex items-start justify-between gap-4 mb-4">
        <div>
          <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
            {{ $t('news.center.title') }}
          </h1>
          <p class="text-gray-600 dark:text-gray-400">
            {{ $t('news.center.subtitle') }}
          </p>
        </div>

        <!-- View Mode Selector -->
        <ViewModeSelector v-model="viewMode" />
      </div>
    </div>

    <!-- Filters Bar -->
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm p-4 mb-6">
      <div class="grid grid-cols-1 md:grid-cols-5 gap-4">
        <!-- Search -->
        <div class="md:col-span-2">
          <div class="relative">
            <Icon icon="mdi:magnify" class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-5 w-5" />
            <input
              v-model="filters.search"
              type="text"
              :placeholder="$t('news.center.searchPlaceholder')"
              class="w-full pl-10 pr-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
              @input="debouncedSearch"
            />
          </div>
        </div>

        <!-- Category Filter -->
        <div>
          <select
            v-model="filters.category_id"
            class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
            @change="fetchNews"
          >
            <option :value="null">{{ $t('news.center.allCategories') }}</option>
            <option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
        </div>

        <!-- Type Filter -->
        <div>
          <select
            v-model="filters.type"
            class="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
            @change="fetchNews"
          >
            <option :value="null">{{ $t('news.center.allTypes') }}</option>
            <option value="article">{{ $t('news.types.article') }}</option>
            <option value="tutorial">{{ $t('news.types.tutorial') }}</option>
            <option value="announcement">{{ $t('news.types.announcement') }}</option>
            <option value="faq">{{ $t('news.types.faq') }}</option>
          </select>
        </div>

        <!-- Sort Selector -->
        <div>
          <SortSelector v-model="sortBy" @update:modelValue="handleSortChange" />
        </div>
      </div>

      <!-- Tags Filter -->
      <div v-if="tags.length > 0" class="mt-4 flex flex-wrap gap-2">
        <button
          v-for="tag in tags"
          :key="tag.id"
          @click="toggleTag(tag.id)"
          :class="[
            'px-3 py-1 rounded-full text-sm font-medium transition-colors',
            selectedTags.includes(tag.id)
              ? 'bg-primary-500 text-white'
              : 'bg-gray-200 dark:bg-gray-700 text-gray-700 dark:text-gray-300 hover:bg-gray-300 dark:hover:bg-gray-600'
          ]"
        >
          {{ tag.name }}
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center py-12">
      <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-primary-500" />
    </div>

    <!-- News List - Grid View -->
    <transition-group
      v-if="newsList.length > 0 && viewMode === 'grid'"
      name="fade-scale"
      tag="div"
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
    >
      <NewsCard
        v-for="news in newsList"
        :key="news.id"
        :news="news"
        @click="viewNews(news)"
      />
    </transition-group>

    <!-- News List - List View -->
    <transition-group
      v-else-if="newsList.length > 0 && viewMode === 'list'"
      name="fade-slide"
      tag="div"
      class="space-y-4"
    >
      <!-- Pinned News -->
      <div v-for="news in pinnedNews" :key="'pinned-' + news.id" class="relative">
        <div class="absolute -top-2 -left-2 z-10">
          <div class="bg-yellow-500 text-white rounded-full p-2 shadow-lg">
            <Icon icon="mdi:pin" class="h-4 w-4" />
          </div>
        </div>
        <NewsCard :news="news" @click="viewNews(news)" />
      </div>

      <!-- Regular News -->
      <NewsCard
        v-for="news in regularNews"
        :key="news.id"
        :news="news"
        @click="viewNews(news)"
      />
    </transition-group>

    <!-- News List - Compact View -->
    <transition-group
      v-else-if="newsList.length > 0 && viewMode === 'compact'"
      name="fade-slide"
      tag="div"
      class="space-y-2"
    >
      <NewsCardCompact
        v-for="news in newsList"
        :key="news.id"
        :news="news"
      />
    </transition-group>

    <!-- Empty State -->
    <div v-else class="text-center py-12">
      <Icon icon="mdi:newspaper-variant-outline" class="h-16 w-16 mx-auto text-gray-400 mb-4" />
      <p class="text-gray-600 dark:text-gray-400 text-lg">
        {{ $t('news.center.noNews') }}
      </p>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="mt-8 flex justify-center">
      <nav class="flex items-center gap-2">
        <button
          @click="changePage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100 dark:hover:bg-gray-700"
        >
          <Icon icon="mdi:chevron-left" class="h-5 w-5" />
        </button>

        <button
          v-for="page in visiblePages"
          :key="page"
          @click="changePage(page)"
          :class="[
            'px-4 py-2 rounded-lg border',
            page === currentPage
              ? 'bg-primary-500 text-white border-primary-500'
              : 'border-gray-300 dark:border-gray-600 hover:bg-gray-100 dark:hover:bg-gray-700'
          ]"
        >
          {{ page }}
        </button>

        <button
          @click="changePage(currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-600 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100 dark:hover:bg-gray-700"
        >
          <Icon icon="mdi:chevron-right" class="h-5 w-5" />
        </button>
      </nav>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { newsService } from '@/services/api'
import NewsCard from '@/components/news/NewsCard.vue'
import NewsCardCompact from '@/components/news/NewsCardCompact.vue'
import ViewModeSelector from '@/components/news/ViewModeSelector.vue'
import SortSelector from '@/components/news/SortSelector.vue'

const router = useRouter()

// View Mode
const viewMode = ref(localStorage.getItem('news-view-mode') || 'grid')

// Watch view mode and save to localStorage
watch(viewMode, (newMode) => {
  localStorage.setItem('news-view-mode', newMode)
})

// Sort
const sortBy = ref(localStorage.getItem('news-sort-by') || 'date-desc')

// Watch sort and save to localStorage
watch(sortBy, (newSort) => {
  localStorage.setItem('news-sort-by', newSort)
})

// Data
const loading = ref(false)
const newsList = ref([])
const categories = ref([])
const tags = ref([])
const selectedTags = ref([])

// Filters
const filters = ref({
  search: '',
  category_id: null,
  type: null
})

// Pagination
const currentPage = ref(1)
const totalPages = ref(1)
const pageSize = 12

// Computed
const pinnedNews = computed(() => newsList.value.filter(n => n.is_pinned))
const regularNews = computed(() => newsList.value.filter(n => !n.is_pinned))

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, currentPage.value + 2)

  for (let i = start; i <= end; i++) {
    pages.push(i)
  }

  return pages
})

// Methods
const fetchNews = async () => {
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      limit: pageSize,
      ...filters.value
    }

    if (selectedTags.value.length > 0) {
      params.tags = selectedTags.value.join(',')
    }

    const response = await newsService.getNews(params)
    newsList.value = response.news || []
    totalPages.value = Math.ceil((response.total || 0) / pageSize)

    // Apply sorting after fetching
    sortNewsList()
  } catch (error) {
    console.error('Error fetching news:', error)
  } finally {
    loading.value = false
  }
}

const fetchCategories = async () => {
  try {
    categories.value = await newsService.getCategories()
  } catch (error) {
    console.error('Error fetching categories:', error)
  }
}

const fetchTags = async () => {
  try {
    tags.value = await newsService.getTags()
  } catch (error) {
    console.error('Error fetching tags:', error)
  }
}

const toggleTag = (tagId) => {
  const index = selectedTags.value.indexOf(tagId)
  if (index > -1) {
    selectedTags.value.splice(index, 1)
  } else {
    selectedTags.value.push(tagId)
  }
  currentPage.value = 1
  fetchNews()
}

const changePage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    fetchNews()
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const viewNews = (news) => {
  router.push({ name: 'NewsDetail', params: { slug: news.slug } })
}

// Debounced search
let searchTimeout = null
const debouncedSearch = () => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(() => {
    currentPage.value = 1
    fetchNews()
  }, 500)
}

// Sort handler
const handleSortChange = () => {
  sortNewsList()
}

const sortNewsList = () => {
  const [field, order] = sortBy.value.split('-')

  newsList.value.sort((a, b) => {
    let aValue, bValue

    switch (field) {
      case 'date':
        aValue = new Date(a.published_at || a.created_at).getTime()
        bValue = new Date(b.published_at || b.created_at).getTime()
        break
      case 'views':
        aValue = a.views || 0
        bValue = b.views || 0
        break
      case 'reactions':
        aValue = (a.reactions && a.reactions.length) || 0
        bValue = (b.reactions && b.reactions.length) || 0
        break
      case 'title':
        aValue = a.title.toLowerCase()
        bValue = b.title.toLowerCase()
        break
      default:
        return 0
    }

    if (order === 'asc') {
      return aValue > bValue ? 1 : aValue < bValue ? -1 : 0
    } else {
      return aValue < bValue ? 1 : aValue > bValue ? -1 : 0
    }
  })
}

// Lifecycle
onMounted(() => {
  fetchNews()
  fetchCategories()
  fetchTags()
})
</script>

<style scoped>
/* Fade Scale Animation for Grid View */
.fade-scale-enter-active,
.fade-scale-leave-active {
  transition: all 0.3s ease;
}

.fade-scale-enter-from {
  opacity: 0;
  transform: scale(0.95);
}

.fade-scale-leave-to {
  opacity: 0;
  transform: scale(0.95);
}

.fade-scale-move {
  transition: transform 0.3s ease;
}

/* Fade Slide Animation for List/Compact Views */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateX(-20px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateX(20px);
}

.fade-slide-move {
  transition: transform 0.3s ease;
}
</style>
