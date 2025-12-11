<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">News Hub</h1>
          <p class="page-subtitle">Manage your news articles and updates</p>
        </div>
        <router-link to="/admin/news/new" class="btn btn-primary">
          <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
          New Article
        </router-link>
      </div>
    </div>

    <!-- Filters -->
    <div class="card mb-6">
      <div class="flex flex-col sm:flex-row gap-4">
        <!-- Search -->
        <div class="flex-1 min-w-[200px]">
          <input
            v-model="filters.search"
            type="text"
            placeholder="Search news..."
            class="input"
          />
        </div>

        <!-- Category filter -->
        <select v-model="filters.category" class="input w-full sm:w-48">
          <option value="">All Categories</option>
          <option v-for="cat in categories" :key="cat.id" :value="cat.id">
            {{ cat.name }}
          </option>
        </select>

        <!-- Status filter -->
        <select v-model="filters.status" class="input w-full sm:w-40">
          <option value="">All Status</option>
          <option value="published">Published</option>
          <option value="draft">Draft</option>
        </select>

        <!-- Priority filter -->
        <select v-model="filters.priority" class="input w-full sm:w-40">
          <option value="">All Priority</option>
          <option value="urgent">Urgent</option>
          <option value="important">Important</option>
          <option value="normal">Normal</option>
        </select>
      </div>
    </div>

    <!-- Loading -->
    <div v-if="isLoading" class="flex justify-center py-12">
      <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-gray-400" />
    </div>

    <!-- News List -->
    <div v-else-if="newsList.length > 0" class="space-y-4">
      <div
        v-for="news in newsList"
        :key="news.id"
        class="card hover:shadow-lg transition-shadow"
      >
        <div class="flex items-start gap-4">
          <!-- Status indicator -->
          <div class="flex-shrink-0 pt-1">
            <div
              :class="news.is_published ? 'bg-green-500' : 'bg-gray-400'"
              class="h-3 w-3 rounded-full"
              :title="news.is_published ? 'Published' : 'Draft'"
            ></div>
          </div>

          <!-- Content -->
          <div class="flex-1 min-w-0">
            <div class="flex items-start justify-between gap-4">
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <!-- Pin indicator -->
                  <Icon
                    v-if="news.is_pinned"
                    icon="mdi:pin"
                    class="h-4 w-4 text-yellow-500"
                  />

                  <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    {{ news.title }}
                  </h3>

                  <!-- Priority badge -->
                  <span
                    v-if="news.priority !== 'normal'"
                    :class="{
                      'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200': news.priority === 'urgent',
                      'bg-orange-100 text-orange-800 dark:bg-orange-900 dark:text-orange-200': news.priority === 'important'
                    }"
                    class="px-2 py-0.5 rounded text-xs font-medium"
                  >
                    {{ news.priority }}
                  </span>
                </div>

                <p v-if="news.summary" class="text-sm text-gray-600 dark:text-gray-400 mb-2">
                  {{ news.summary }}
                </p>

                <div class="flex flex-wrap items-center gap-4 text-xs text-gray-500 dark:text-gray-400">
                  <!-- Category -->
                  <div v-if="news.category" class="flex items-center gap-1">
                    <Icon :icon="news.category.icon" class="h-4 w-4" />
                    <span>{{ news.category.name }}</span>
                  </div>

                  <!-- Author -->
                  <div class="flex items-center gap-1">
                    <Icon icon="mdi:account" class="h-4 w-4" />
                    <span>{{ news.author?.username || 'Unknown' }}</span>
                  </div>

                  <!-- Date -->
                  <div class="flex items-center gap-1">
                    <Icon icon="mdi:calendar" class="h-4 w-4" />
                    <span>{{ formatDate(news.published_at || news.created_at) }}</span>
                  </div>

                  <!-- Views -->
                  <div class="flex items-center gap-1">
                    <Icon icon="mdi:eye" class="h-4 w-4" />
                    <span>{{ news.view_count }}</span>
                  </div>

                  <!-- Tags -->
                  <div v-if="news.tags && news.tags.length > 0" class="flex items-center gap-1">
                    <Icon icon="mdi:tag" class="h-4 w-4" />
                    <div class="flex gap-1">
                      <span
                        v-for="tag in news.tags.slice(0, 2)"
                        :key="tag.id"
                        class="px-1.5 py-0.5 bg-gray-100 dark:bg-gray-700 rounded text-xs"
                      >
                        {{ tag.name }}
                      </span>
                      <span v-if="news.tags.length > 2" class="text-xs">
                        +{{ news.tags.length - 2 }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Actions -->
              <div class="flex items-center gap-2">
                <router-link
                  :to="`/admin/news/${news.slug}/edit`"
                  class="btn btn-secondary btn-sm"
                  title="Edit"
                >
                  <Icon icon="mdi:pencil" class="h-4 w-4" />
                </router-link>

                <button
                  v-if="authStore.isAdmin"
                  @click="togglePin(news)"
                  class="btn btn-secondary btn-sm"
                  :title="news.is_pinned ? 'Unpin' : 'Pin'"
                >
                  <Icon :icon="news.is_pinned ? 'mdi:pin-off' : 'mdi:pin'" class="h-4 w-4" />
                </button>

                <button
                  @click="confirmDelete(news)"
                  class="btn btn-danger btn-sm"
                  title="Delete"
                >
                  <Icon icon="mdi:delete" class="h-4 w-4" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div v-if="pagination.total_pages > 1" class="flex justify-center gap-2 mt-6">
        <button
          @click="changePage(pagination.page - 1)"
          :disabled="pagination.page === 1"
          class="btn btn-secondary btn-sm"
        >
          <Icon icon="mdi:chevron-left" class="h-4 w-4" />
        </button>

        <span class="px-4 py-2 text-sm text-gray-700 dark:text-gray-300">
          Page {{ pagination.page }} of {{ pagination.total_pages }}
        </span>

        <button
          @click="changePage(pagination.page + 1)"
          :disabled="pagination.page === pagination.total_pages"
          class="btn btn-secondary btn-sm"
        >
          <Icon icon="mdi:chevron-right" class="h-4 w-4" />
        </button>
      </div>
    </div>

    <!-- Empty state -->
    <div v-else class="empty-state">
      <Icon icon="mdi:newspaper" class="empty-state-icon" />
      <h3 class="empty-state-title">No News Articles</h3>
      <p class="empty-state-description">
        Start by creating your first news article.
      </p>
      <router-link to="/admin/news/new" class="btn btn-primary mt-4">
        <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
        Create Article
      </router-link>
    </div>

    <!-- Delete Confirmation Modal -->
    <div
      v-if="newsToDelete"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      @click.self="newsToDelete = null"
    >
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 max-w-md w-full mx-4">
        <h3 class="text-lg font-semibold mb-4 text-gray-900 dark:text-white">
          Delete News Article
        </h3>
        <p class="text-gray-600 dark:text-gray-400 mb-6">
          Are you sure you want to delete "{{ newsToDelete.title }}"? This action cannot be undone.
        </p>
        <div class="flex justify-end gap-3">
          <button @click="newsToDelete = null" class="btn btn-secondary">
            Cancel
          </button>
          <button @click="deleteNews" class="btn btn-danger">
            Delete
          </button>
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
import { newsService } from '@/services/api'

const authStore = useAuthStore()
const appStore = useAppStore()

const isLoading = ref(false)
const newsList = ref([])
const categories = ref([])
const newsToDelete = ref(null)

const filters = ref({
  search: '',
  category: '',
  status: '',
  priority: '',
  page: 1,
  page_size: 10
})

const pagination = ref({
  page: 1,
  page_size: 10,
  total: 0,
  total_pages: 0
})

// Load news
const loadNews = async () => {
  try {
    isLoading.value = true

    const params = {
      page: filters.value.page,
      page_size: filters.value.page_size
    }

    if (filters.value.search) params.search = filters.value.search
    if (filters.value.category) params.category_id = filters.value.category
    if (filters.value.priority) params.priority = filters.value.priority

    // For status filter
    if (filters.value.status === 'published') {
      // Will be handled by backend to show only published
    } else if (filters.value.status === 'draft') {
      // Will be handled by backend to show only drafts
    }

    const response = await newsService.getNews(params)
    newsList.value = response.news || []
    pagination.value = {
      page: response.page,
      page_size: response.page_size,
      total: response.total,
      total_pages: response.total_pages
    }
  } catch (error) {
    console.error('Error loading news:', error)
    appStore.showError('Failed to load news')
  } finally {
    isLoading.value = false
  }
}

// Load categories
const loadCategories = async () => {
  try {
    categories.value = await newsService.getCategories()
  } catch (error) {
    console.error('Error loading categories:', error)
  }
}

// Toggle pin
const togglePin = async (news) => {
  try {
    await newsService.togglePin(news.id)
    appStore.showSuccess(news.is_pinned ? 'News unpinned' : 'News pinned')
    loadNews()
  } catch (error) {
    console.error('Error toggling pin:', error)
    appStore.showError('Failed to update news')
  }
}

// Confirm delete
const confirmDelete = (news) => {
  newsToDelete.value = news
}

// Delete news
const deleteNews = async () => {
  try {
    await newsService.deleteNews(newsToDelete.value.id)
    appStore.showSuccess('News deleted successfully')
    newsToDelete.value = null
    loadNews()
  } catch (error) {
    console.error('Error deleting news:', error)
    appStore.showError('Failed to delete news')
  }
}

// Change page
const changePage = (page) => {
  if (page >= 1 && page <= pagination.value.total_pages) {
    filters.value.page = page
  }
}

// Format date
const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

// Watch filters
watch(
  () => [filters.value.search, filters.value.category, filters.value.status, filters.value.priority, filters.value.page],
  () => {
    loadNews()
  }
)

// Lifecycle
onMounted(() => {
  loadNews()
  loadCategories()
})
</script>
