<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">{{ isEditMode ? 'Edit Article' : 'New Article' }}</h1>
          <p class="page-subtitle">{{ isEditMode ? 'Update your news article' : 'Create a new news article' }}</p>
        </div>
        <router-link to="/admin/news" class="btn btn-secondary">
          <Icon icon="mdi:arrow-left" class="h-4 w-4 mr-2" />
          Back to List
        </router-link>
      </div>
    </div>

    <!-- Form -->
    <form @submit.prevent="saveNews" class="space-y-6">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Main content -->
        <div class="lg:col-span-2 space-y-6">
          <!-- Title -->
          <div class="card">
            <div class="mb-3">
              <label class="block text-sm font-semibold text-gray-900 dark:text-white mb-2 flex items-center gap-2">
                <Icon icon="mdi:format-title" class="h-5 w-5 text-primary-500" />
                Title *
              </label>
              <input
                v-model="form.title"
                type="text"
                required
                placeholder="Enter a compelling title for your article..."
                class="w-full px-4 py-3 text-lg font-medium border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white transition-colors"
              />
            </div>
          </div>

          <!-- Summary -->
          <div class="card">
            <div class="mb-3">
              <label class="block text-sm font-semibold text-gray-900 dark:text-white mb-2 flex items-center gap-2">
                <Icon icon="mdi:text-short" class="h-5 w-5 text-primary-500" />
                Summary
              </label>
              <textarea
                v-model="form.summary"
                rows="3"
                maxlength="300"
                placeholder="Write a brief summary that will appear in listings and previews..."
                class="w-full px-4 py-3 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white resize-none transition-colors"
              ></textarea>
              <div class="flex justify-between items-center mt-2">
                <p class="text-xs text-gray-500 dark:text-gray-400">
                  Brief description for previews and search
                </p>
                <p class="text-xs font-medium" :class="form.summary.length >= 300 ? 'text-red-500' : 'text-gray-500 dark:text-gray-400'">
                  {{ form.summary.length }}/300
                </p>
              </div>
            </div>
          </div>

          <!-- Content -->
          <div class="card">
            <div class="mb-3">
              <label class="block text-sm font-semibold text-gray-900 dark:text-white mb-3 flex items-center gap-2">
                <Icon icon="mdi:text-box" class="h-5 w-5 text-primary-500" />
                Article Content *
              </label>
              <div class="border-2 border-gray-200 dark:border-gray-700 rounded-lg overflow-hidden">
                <RichTextEditor
                  v-model="form.content"
                  placeholder="Write your article content here. Use the toolbar to format your text, add links, code blocks, and more..."
                />
              </div>
            </div>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="space-y-6">
          <!-- Publish settings -->
          <div class="card bg-gradient-to-br from-white to-gray-50 dark:from-gray-800 dark:to-gray-900">
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
              <Icon icon="mdi:publish" class="h-5 w-5 text-primary-500" />
              Publication
            </h3>

            <div class="space-y-4">
              <!-- Status -->
              <div class="bg-white dark:bg-gray-800 rounded-lg p-3 border-2 border-gray-200 dark:border-gray-700">
                <label class="flex items-center cursor-pointer">
                  <input
                    v-model="form.is_published"
                    type="checkbox"
                    class="form-checkbox h-5 w-5 text-primary-600 rounded focus:ring-2 focus:ring-primary-500"
                  />
                  <span class="ml-3 text-sm font-medium text-gray-900 dark:text-white">
                    Publish immediately
                  </span>
                </label>
              </div>

              <!-- Publish date -->
              <div v-if="form.is_published">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:calendar-clock" class="h-4 w-4" />
                  Publish Date
                </label>
                <input
                  v-model="form.published_at"
                  type="datetime-local"
                  class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm"
                />
              </div>

              <!-- Expiry date -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:calendar-remove" class="h-4 w-4" />
                  Expiry Date
                  <span class="text-xs text-gray-500 dark:text-gray-400 ml-1">(Optional)</span>
                </label>
                <input
                  v-model="form.expires_at"
                  type="datetime-local"
                  class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm"
                />
                <p class="text-xs text-gray-500 dark:text-gray-400 mt-1.5 flex items-center gap-1">
                  <Icon icon="mdi:information-outline" class="h-3 w-3" />
                  Auto-archive after this date
                </p>
              </div>
            </div>
          </div>

          <!-- Category & Priority -->
          <div class="card">
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
              <Icon icon="mdi:tag-multiple" class="h-5 w-5 text-primary-500" />
              Classification
            </h3>

            <div class="space-y-4">
              <!-- Category -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:folder" class="h-4 w-4" />
                  Category
                </label>
                <div class="flex gap-2">
                  <select v-model="form.category_id" class="flex-1 px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm">
                    <option :value="null">No category</option>
                    <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                      {{ cat.name }}
                    </option>
                  </select>
                  <button
                    type="button"
                    @click="showAddCategoryDialog = true"
                    class="flex-shrink-0 px-3 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors flex items-center gap-1"
                    title="Add new category"
                  >
                    <Icon icon="mdi:plus" class="h-4 w-4" />
                  </button>
                </div>
              </div>

              <!-- Type -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:shape" class="h-4 w-4" />
                  Type
                </label>
                <select v-model="form.type" class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm">
                  <option value="article">📄 Article</option>
                  <option value="tutorial">📚 Tutorial</option>
                  <option value="announcement">📢 Announcement</option>
                  <option value="faq">FAQ</option>
                </select>
              </div>

              <!-- Priority -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:flag" class="h-4 w-4" />
                  Priority
                </label>
                <select v-model="form.priority" class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm">
                  <option value="normal">⚪ Normal</option>
                  <option value="important">🟠 Important</option>
                  <option value="urgent">🔴 Urgent</option>
                </select>
              </div>
            </div>
          </div>

          <!-- Tags -->
          <div class="card">
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
              <Icon icon="mdi:tag-multiple" class="h-5 w-5 text-primary-500" />
              Tags
            </h3>

            <div class="space-y-3">
              <!-- Selected tags -->
              <div v-if="form.tag_ids.length > 0" class="flex flex-wrap gap-2 p-3 bg-gray-50 dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-gray-700">
                <span
                  v-for="tagId in form.tag_ids"
                  :key="tagId"
                  class="inline-flex items-center gap-1 px-3 py-1.5 bg-primary-100 dark:bg-primary-900/30 text-primary-800 dark:text-primary-300 rounded-full text-sm font-medium"
                >
                  <Icon icon="mdi:tag" class="h-3 w-3" />
                  {{ getTagName(tagId) }}
                  <button
                    type="button"
                    @click="removeTag(tagId)"
                    class="ml-1 hover:text-red-600 dark:hover:text-red-400 transition-colors"
                  >
                    <Icon icon="mdi:close-circle" class="h-4 w-4" />
                  </button>
                </span>
              </div>
              <div v-else class="text-sm text-gray-500 dark:text-gray-400 italic p-3 bg-gray-50 dark:bg-gray-900 rounded-lg border border-gray-200 dark:border-gray-700">
                No tags selected
              </div>

              <!-- Add tag -->
              <div class="flex gap-2">
                <select
                  @change="addTag"
                  class="flex-1 px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm"
                >
                  <option value="">+ Add tag</option>
                  <option
                    v-for="tag in availableTags"
                    :key="tag.id"
                    :value="tag.id"
                  >
                    {{ tag.name }}
                  </option>
                </select>
                <button
                  type="button"
                  @click="showAddTagDialog = true"
                  class="flex-shrink-0 px-3 py-2 bg-primary-500 text-white rounded-lg hover:bg-primary-600 transition-colors flex items-center gap-1"
                  title="Add new tag"
                >
                  <Icon icon="mdi:plus" class="h-4 w-4" />
                </button>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="card bg-gradient-to-br from-primary-50 to-primary-100 dark:from-primary-900/20 dark:to-primary-800/20 border-2 border-primary-200 dark:border-primary-800">
            <div class="space-y-3">
              <button
                type="submit"
                :disabled="isSaving"
                class="w-full px-6 py-3 bg-primary-600 hover:bg-primary-700 disabled:bg-gray-400 text-white font-semibold rounded-lg transition-all transform hover:scale-[1.02] active:scale-[0.98] disabled:scale-100 shadow-lg hover:shadow-xl flex items-center justify-center gap-2"
              >
                <Icon v-if="isSaving" icon="mdi:loading" class="h-5 w-5 animate-spin" />
                <Icon v-else icon="mdi:content-save" class="h-5 w-5" />
                <span>{{ isSaving ? 'Saving...' : (isEditMode ? 'Update Article' : 'Create Article') }}</span>
              </button>

              <router-link
                to="/admin/news"
                class="w-full px-6 py-3 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-800 dark:text-gray-200 font-medium rounded-lg transition-all flex items-center justify-center gap-2"
              >
                <Icon icon="mdi:close" class="h-5 w-5" />
                <span>Cancel</span>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </form>

    <!-- Add Category Dialog -->
    <div v-if="showAddCategoryDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 max-w-md w-full mx-4">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Add New Category</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Name</label>
            <input
              v-model="newCategory.name"
              type="text"
              class="input"
              placeholder="Category name"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Icon (Iconify)</label>
            <input
              v-model="newCategory.icon"
              type="text"
              class="input"
              placeholder="mdi:folder"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Color</label>
            <input
              v-model="newCategory.color"
              type="color"
              class="input h-10"
            />
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button
            type="button"
            @click="showAddCategoryDialog = false"
            class="px-4 py-2 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded"
          >
            Cancel
          </button>
          <button
            type="button"
            @click="createCategory"
            class="px-4 py-2 bg-primary-500 text-white rounded hover:bg-primary-600"
          >
            Create
          </button>
        </div>
      </div>
    </div>

    <!-- Add Tag Dialog -->
    <div v-if="showAddTagDialog" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white dark:bg-gray-800 rounded-lg p-6 max-w-md w-full mx-4">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">Add New Tag</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Name</label>
            <input
              v-model="newTag.name"
              type="text"
              class="input"
              placeholder="Tag name"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">Color</label>
            <input
              v-model="newTag.color"
              type="color"
              class="input h-10"
            />
          </div>
        </div>
        <div class="flex justify-end gap-2 mt-6">
          <button
            type="button"
            @click="showAddTagDialog = false"
            class="px-4 py-2 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded"
          >
            Cancel
          </button>
          <button
            type="button"
            @click="createTag"
            class="px-4 py-2 bg-primary-500 text-white rounded hover:bg-primary-600"
          >
            Create
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { useAppStore } from '@/stores/app'
import { newsService } from '@/services/api'
import RichTextEditor from '@/components/news/RichTextEditor.vue'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

const isEditMode = computed(() => route.params.id && route.params.id !== 'new')
const isSaving = ref(false)
const categories = ref([])
const tags = ref([])

// Dialogs
const showAddCategoryDialog = ref(false)
const showAddTagDialog = ref(false)

// New category/tag forms
const newCategory = ref({
  name: '',
  icon: 'mdi:folder',
  color: '#6366f1'
})

const newTag = ref({
  name: '',
  color: '#6366f1'
})

const form = ref({
  title: '',
  summary: '',
  content: '',
  type: 'article',
  priority: 'normal',
  is_published: false,
  published_at: null,
  expires_at: null,
  category_id: null,
  tag_ids: []
})

// Available tags (excluding already selected)
const availableTags = computed(() => {
  return tags.value.filter(tag => !form.value.tag_ids.includes(tag.id))
})

// Get tag name by ID
const getTagName = (tagId) => {
  const tag = tags.value.find(t => t.id === tagId)
  return tag ? tag.name : ''
}

// Add tag
const addTag = (event) => {
  const tagId = parseInt(event.target.value)
  if (tagId && !form.value.tag_ids.includes(tagId)) {
    form.value.tag_ids.push(tagId)
  }
  event.target.value = ''
}

// Remove tag
const removeTag = (tagId) => {
  form.value.tag_ids = form.value.tag_ids.filter(id => id !== tagId)
}

// Create new category
const createCategory = async () => {
  try {
    const category = await newsService.createCategory(newCategory.value)
    categories.value.push(category)
    form.value.category_id = category.id
    showAddCategoryDialog.value = false
    // Reset form
    newCategory.value = {
      name: '',
      icon: 'mdi:folder',
      color: '#6366f1'
    }
  } catch (error) {
    console.error('Error creating category:', error)
    alert('Failed to create category')
  }
}

// Create new tag
const createTag = async () => {
  try {
    const tag = await newsService.createTag(newTag.value)
    tags.value.push(tag)
    form.value.tag_ids.push(tag.id)
    showAddTagDialog.value = false
    // Reset form
    newTag.value = {
      name: '',
      color: '#6366f1'
    }
  } catch (error) {
    console.error('Error creating tag:', error)
    alert('Failed to create tag')
  }
}

// Load news for editing
const loadNews = async () => {
  if (!isEditMode.value) return

  try {
    const news = await newsService.getNewsBySlug(route.params.id)

    // Parse content if it's a JSON string
    let content = news.content || ''
    if (typeof content === 'string' && content.trim().startsWith('{')) {
      try {
        content = JSON.parse(content)
      } catch (e) {
        console.warn('Failed to parse content as JSON, using as-is')
      }
    }

    form.value = {
      title: news.title,
      summary: news.summary || '',
      content: content,
      type: news.type,
      priority: news.priority,
      is_published: news.is_published,
      published_at: news.published_at ? new Date(news.published_at).toISOString().slice(0, 16) : null,
      expires_at: news.expires_at ? new Date(news.expires_at).toISOString().slice(0, 16) : null,
      category_id: news.category_id,
      tag_ids: news.tags ? news.tags.map(t => t.id) : []
    }
  } catch (error) {
    console.error('Error loading news:', error)
    appStore.showError('Failed to load news')
    router.push('/admin/news')
  }
}

// Load categories and tags
const loadMetadata = async () => {
  try {
    [categories.value, tags.value] = await Promise.all([
      newsService.getCategories(),
      newsService.getTags()
    ])
  } catch (error) {
    console.error('Error loading metadata:', error)
  }
}

// Save news
const saveNews = async () => {
  try {
    isSaving.value = true

    const data = {
      ...form.value,
      // Convert Tiptap JSON content to string
      content: typeof form.value.content === 'object'
        ? JSON.stringify(form.value.content)
        : form.value.content,
      published_at: form.value.published_at ? new Date(form.value.published_at).toISOString() : null,
      expires_at: form.value.expires_at ? new Date(form.value.expires_at).toISOString() : null
    }

    if (isEditMode.value) {
      await newsService.updateNews(route.params.id, data)
      appStore.showSuccess('Article updated successfully')
    } else {
      await newsService.createNews(data)
      appStore.showSuccess('Article created successfully')
    }

    router.push('/admin/news')
  } catch (error) {
    console.error('Error saving news:', error)
    appStore.showError('Failed to save article')
  } finally {
    isSaving.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadMetadata()
  loadNews()
})
</script>
