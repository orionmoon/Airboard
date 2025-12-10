<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">News de Groupe</h1>
          <p class="page-subtitle">Gérez les actualités pour vos groupes</p>
        </div>
        <router-link to="/group-admin/news/new" class="btn btn-primary">
          <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
          Nouvel Article
        </router-link>
      </div>
    </div>

    <!-- Info Alert -->
    <div class="mb-6 p-4 bg-blue-900/20 border border-blue-600 rounded-lg">
      <div class="flex items-start">
        <Icon icon="mdi:information" class="h-5 w-5 text-blue-500 mr-3 mt-0.5" />
        <div>
          <h3 class="text-blue-500 font-medium">Gestion des actualités de groupe</h3>
          <p class="text-sm text-blue-200 mt-1">
            En tant que Group Admin, vous pouvez créer des actualités ciblées pour vos groupes administrés.
            Ces actualités seront visibles uniquement par les membres des groupes que vous sélectionnez.
          </p>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="card mb-6">
      <div class="flex flex-wrap gap-4">
        <!-- Search -->
        <div class="flex-1 min-w-[200px]">
          <input
            v-model="filters.search"
            type="text"
            placeholder="Rechercher des articles..."
            class="input"
          />
        </div>

        <!-- Category filter -->
        <select v-model="filters.category" class="input w-48">
          <option value="">Toutes les catégories</option>
          <option v-for="cat in categories" :key="cat.id" :value="cat.id">
            {{ cat.name }}
          </option>
        </select>

        <!-- Status filter -->
        <select v-model="filters.status" class="input w-40">
          <option value="">Tous les statuts</option>
          <option value="published">Publié</option>
          <option value="draft">Brouillon</option>
        </select>

        <!-- Priority filter -->
        <select v-model="filters.priority" class="input w-40">
          <option value="">Toutes les priorités</option>
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
              :title="news.is_published ? 'Publié' : 'Brouillon'"
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
                    <span>{{ news.author?.username || 'Inconnu' }}</span>
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

                  <!-- Target Groups -->
                  <div v-if="news.target_groups && news.target_groups.length > 0" class="flex items-center gap-1">
                    <Icon icon="mdi:account-group" class="h-4 w-4" />
                    <div class="flex gap-1">
                      <span
                        v-for="group in news.target_groups.slice(0, 2)"
                        :key="group.id"
                        class="px-1.5 py-0.5 bg-blue-100 dark:bg-blue-900 rounded text-xs"
                      >
                        {{ group.name }}
                      </span>
                      <span v-if="news.target_groups.length > 2" class="text-xs">
                        +{{ news.target_groups.length - 2 }}
                      </span>
                    </div>
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
                  :to="`/group-admin/news/${news.slug}/edit`"
                  class="btn btn-secondary btn-sm"
                  title="Modifier"
                >
                  <Icon icon="mdi:pencil" class="h-4 w-4" />
                </router-link>

                <button
                  @click="confirmDelete(news)"
                  class="btn btn-danger btn-sm"
                  title="Supprimer"
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
          Page {{ pagination.page }} sur {{ pagination.total_pages }}
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

    <!-- Empty State -->
    <div v-else class="empty-state">
      <Icon icon="mdi:newspaper-variant" class="empty-state-icon" />
      <h3 class="empty-state-title">Aucune actualité</h3>
      <p class="empty-state-description">
        {{ filters.search || filters.category || filters.status || filters.priority
          ? 'Aucun article ne correspond à vos critères.'
          : 'Créez votre première actualité pour vos groupes.'
        }}
      </p>
      <router-link v-if="!filters.search && !filters.category" to="/group-admin/news/new" class="btn btn-primary">
        <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
        Créer un article
      </router-link>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="modal-overlay">
      <div class="modal-container">
        <div class="modal-panel sm:max-w-lg sm:w-full">
          <div class="modal-header">
            <div class="flex items-start gap-3">
              <div class="flex-shrink-0">
                <Icon icon="mdi:alert" class="h-6 w-6 text-red-500" />
              </div>
              <div>
                <h3 class="modal-title">Supprimer l'article</h3>
                <p class="modal-subtitle">
                  Êtes-vous sûr de vouloir supprimer "<strong>{{ newsToDelete?.title }}</strong>" ?
                  Cette action est irréversible.
                </p>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button @click="closeDeleteModal" class="btn btn-secondary w-full sm:w-auto">
              Annuler
            </button>
            <button
              @click="deleteNews"
              :disabled="deleteLoading"
              class="btn btn-danger w-full sm:w-auto"
            >
              <Icon v-if="deleteLoading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              Supprimer
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { Icon } from '@iconify/vue'
import { newsService, groupAdminService } from '@/services/api'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

// State
const newsList = ref([])
const categories = ref([])
const isLoading = ref(false)
const showDeleteModal = ref(false)
const newsToDelete = ref(null)
const deleteLoading = ref(false)

const filters = ref({
  search: '',
  category: '',
  status: '',
  priority: ''
})

const pagination = ref({
  page: 1,
  limit: 20,
  total: 0,
  total_pages: 0
})

// Methods
const loadNews = async () => {
  try {
    isLoading.value = true
    const params = {
      page: pagination.value.page,
      page_size: pagination.value.limit,
      search: filters.value.search || undefined,
      category_id: filters.value.category || undefined,
      priority: filters.value.priority || undefined
    }

    const response = await groupAdminService.getNews(params)
    newsList.value = response.news || []
    pagination.value = {
      ...pagination.value,
      page: response.page || 1,
      total: response.total || 0,
      total_pages: response.total_pages || 1
    }
  } catch (error) {
    console.error('Error loading news:', error)
    newsList.value = []
  } finally {
    isLoading.value = false
  }
}

const loadCategories = async () => {
  try {
    const data = await newsService.getCategories()
    categories.value = Array.isArray(data) ? data : (data.data || [])
  } catch (error) {
    console.error('Error loading categories:', error)
    categories.value = []
  }
}

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  const date = new Date(dateString)
  return date.toLocaleDateString('fr-FR', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const changePage = (page) => {
  if (page >= 1 && page <= pagination.value.total_pages) {
    pagination.value.page = page
    loadNews()
  }
}

const confirmDelete = (news) => {
  newsToDelete.value = news
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  newsToDelete.value = null
}

const deleteNews = async () => {
  if (!newsToDelete.value) return

  try {
    deleteLoading.value = true
    await groupAdminService.deleteNews(newsToDelete.value.id)
    closeDeleteModal()
    await loadNews()
  } catch (error) {
    console.error('Error deleting news:', error)
    alert('Erreur lors de la suppression de l\'article')
  } finally {
    deleteLoading.value = false
  }
}

// Watch filters
watch(
  () => [filters.value.search, filters.value.category, filters.value.status, filters.value.priority],
  () => {
    pagination.value.page = 1
    loadNews()
  }
)

onMounted(async () => {
  await loadCategories()
  await loadNews()
})
</script>
