<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">{{ isEditMode ? 'Modifier l\'Article' : 'Nouvel Article de Groupe' }}</h1>
          <p class="page-subtitle">{{ isEditMode ? 'Modifiez votre actualité de groupe' : 'Créez une nouvelle actualité pour vos groupes' }}</p>
        </div>
        <router-link to="/group-admin/news" class="btn btn-secondary">
          <Icon icon="mdi:arrow-left" class="h-4 w-4 mr-2" />
          Retour à la liste
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
                Titre *
              </label>
              <input
                v-model="form.title"
                type="text"
                required
                placeholder="Entrez un titre accrocheur pour votre article..."
                class="w-full px-4 py-3 text-lg font-medium border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white transition-colors"
              />
            </div>
          </div>

          <!-- Summary -->
          <div class="card">
            <div class="mb-3">
              <label class="block text-sm font-semibold text-gray-900 dark:text-white mb-2 flex items-center gap-2">
                <Icon icon="mdi:text-short" class="h-5 w-5 text-primary-500" />
                Résumé
              </label>
              <textarea
                v-model="form.summary"
                rows="3"
                maxlength="300"
                placeholder="Écrivez un bref résumé qui apparaîtra dans les listes et prévisualisations..."
                class="w-full px-4 py-3 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white resize-none transition-colors"
              ></textarea>
              <div class="flex justify-between items-center mt-2">
                <p class="text-xs text-gray-500 dark:text-gray-400">
                  Brève description pour les prévisualisations
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
                Contenu de l'Article *
              </label>
              <div class="border-2 border-gray-200 dark:border-gray-700 rounded-lg overflow-hidden">
                <RichTextEditor
                  v-model="form.content"
                  placeholder="Rédigez le contenu de votre article ici. Utilisez la barre d'outils pour formater votre texte, ajouter des liens, des blocs de code, etc..."
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
                    Publier immédiatement
                  </span>
                </label>
              </div>

              <!-- Publish date -->
              <div v-if="form.is_published">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:calendar-clock" class="h-4 w-4" />
                  Date de publication
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
                  Date d'expiration
                  <span class="text-xs text-gray-500 dark:text-gray-400 ml-1">(Optionnel)</span>
                </label>
                <input
                  v-model="form.expires_at"
                  type="datetime-local"
                  class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm"
                />
                <p class="text-xs text-gray-500 dark:text-gray-400 mt-1.5 flex items-center gap-1">
                  <Icon icon="mdi:information-outline" class="h-3 w-3" />
                  Auto-archivage après cette date
                </p>
              </div>
            </div>
          </div>

          <!-- Target Groups (NOUVEAU pour Group Admin) -->
          <div class="card bg-gradient-to-br from-blue-50 to-blue-100 dark:from-blue-900/20 dark:to-blue-800/20 border-2 border-blue-200 dark:border-blue-800">
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white mb-4 flex items-center gap-2">
              <Icon icon="mdi:account-group" class="h-5 w-5 text-blue-500" />
              Groupes Cibles *
            </h3>

            <div class="space-y-3">
              <p class="text-xs text-gray-600 dark:text-gray-400">
                Sélectionnez les groupes qui verront cette actualité
              </p>

              <!-- Selected groups -->
              <div v-if="form.target_group_ids.length > 0" class="flex flex-wrap gap-2 p-3 bg-white dark:bg-gray-800 rounded-lg border border-blue-200 dark:border-blue-700">
                <span
                  v-for="groupId in form.target_group_ids"
                  :key="groupId"
                  class="inline-flex items-center gap-1 px-3 py-1.5 bg-blue-100 dark:bg-blue-900/30 text-blue-800 dark:text-blue-300 rounded-full text-sm font-medium"
                >
                  <Icon icon="mdi:account-group" class="h-3 w-3" />
                  {{ getGroupName(groupId) }}
                  <button
                    type="button"
                    @click="removeGroup(groupId)"
                    class="ml-1 hover:text-red-600 dark:hover:text-red-400 transition-colors"
                  >
                    <Icon icon="mdi:close-circle" class="h-4 w-4" />
                  </button>
                </span>
              </div>
              <div v-else class="text-sm text-red-500 dark:text-red-400 italic p-3 bg-white dark:bg-gray-800 rounded-lg border border-red-200 dark:border-red-700">
                ⚠️ Aucun groupe sélectionné - l'article ne sera visible par personne
              </div>

              <!-- Add group -->
              <select
                @change="addGroup"
                class="w-full px-3 py-2 border-2 border-blue-200 dark:border-blue-700 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-800 dark:text-white text-sm"
              >
                <option value="">+ Ajouter un groupe</option>
                <option
                  v-for="group in availableGroups"
                  :key="group.id"
                  :value="group.id"
                >
                  {{ group.name }}
                </option>
              </select>
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
                  Catégorie
                </label>
                <select v-model="form.category_id" class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm">
                  <option :value="null">Aucune catégorie</option>
                  <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                    {{ cat.name }}
                  </option>
                </select>
              </div>

              <!-- Type -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:shape" class="h-4 w-4" />
                  Type
                </label>
                <select v-model="form.type" class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm">
                  <option value="article">📄 Article</option>
                  <option value="tutorial">📚 Tutoriel</option>
                  <option value="announcement">📢 Annonce</option>
                  <option value="faq">❓ FAQ</option>
                </select>
              </div>

              <!-- Priority -->
              <div>
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2 flex items-center gap-1">
                  <Icon icon="mdi:flag" class="h-4 w-4" />
                  Priorité
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
                Aucun tag sélectionné
              </div>

              <!-- Add tag -->
              <select
                @change="addTag"
                class="w-full px-3 py-2 border-2 border-gray-200 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-800 dark:text-white text-sm"
              >
                <option value="">+ Ajouter un tag</option>
                <option
                  v-for="tag in availableTags"
                  :key="tag.id"
                  :value="tag.id"
                >
                  {{ tag.name }}
                </option>
              </select>
            </div>
          </div>

          <!-- Actions -->
          <div class="card bg-gradient-to-br from-primary-50 to-primary-100 dark:from-primary-900/20 dark:to-primary-800/20 border-2 border-primary-200 dark:border-primary-800">
            <div class="space-y-3">
              <button
                type="submit"
                :disabled="isSaving || form.target_group_ids.length === 0"
                class="w-full px-6 py-3 bg-primary-600 hover:bg-primary-700 disabled:bg-gray-400 text-white font-semibold rounded-lg transition-all transform hover:scale-[1.02] active:scale-[0.98] disabled:scale-100 shadow-lg hover:shadow-xl flex items-center justify-center gap-2"
              >
                <Icon v-if="isSaving" icon="mdi:loading" class="h-5 w-5 animate-spin" />
                <Icon v-else icon="mdi:content-save" class="h-5 w-5" />
                <span>{{ isSaving ? 'Enregistrement...' : (isEditMode ? 'Mettre à jour' : 'Créer l\'article') }}</span>
              </button>

              <p v-if="form.target_group_ids.length === 0" class="text-xs text-center text-red-500 dark:text-red-400">
                Sélectionnez au moins un groupe pour publier
              </p>

              <router-link
                to="/group-admin/news"
                class="w-full px-6 py-3 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-800 dark:text-gray-200 font-medium rounded-lg transition-all flex items-center justify-center gap-2"
              >
                <Icon icon="mdi:close" class="h-5 w-5" />
                <span>Annuler</span>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Icon } from '@iconify/vue'
import { useAppStore } from '@/stores/app'
import { newsService, groupAdminService } from '@/services/api'
import RichTextEditor from '@/components/news/RichTextEditor.vue'

const route = useRoute()
const router = useRouter()
const appStore = useAppStore()

const isEditMode = computed(() => route.params.slug && route.params.slug !== 'new')
const isSaving = ref(false)
const categories = ref([])
const tags = ref([])
const managedGroups = ref([]) // Groupes gérés par le group admin

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
  tag_ids: [],
  target_group_ids: [] // NOUVEAU: groupes cibles
})

// Available groups (excluding already selected)
const availableGroups = computed(() => {
  return managedGroups.value.filter(group => !form.value.target_group_ids.includes(group.id))
})

// Available tags (excluding already selected)
const availableTags = computed(() => {
  return tags.value.filter(tag => !form.value.tag_ids.includes(tag.id))
})

// Get group name by ID
const getGroupName = (groupId) => {
  const group = managedGroups.value.find(g => g.id === groupId)
  return group ? group.name : ''
}

// Get tag name by ID
const getTagName = (tagId) => {
  const tag = tags.value.find(t => t.id === tagId)
  return tag ? tag.name : ''
}

// Add group
const addGroup = (event) => {
  const groupId = parseInt(event.target.value)
  if (groupId && !form.value.target_group_ids.includes(groupId)) {
    form.value.target_group_ids.push(groupId)
  }
  event.target.value = ''
}

// Remove group
const removeGroup = (groupId) => {
  form.value.target_group_ids = form.value.target_group_ids.filter(id => id !== groupId)
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

// Load managed groups (groups that this group admin manages)
const loadManagedGroups = async () => {
  try {
    const data = await groupAdminService.getManagedGroups()
    managedGroups.value = Array.isArray(data) ? data : (data.groups || data.data || [])
  } catch (error) {
    console.error('Error loading managed groups:', error)
    managedGroups.value = []
  }
}

// Load news for editing
const loadNews = async () => {
  if (!isEditMode.value) return

  try {
    const news = await newsService.getNewsBySlug(route.params.slug)

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
      tag_ids: news.tags ? news.tags.map(t => t.id) : [],
      target_group_ids: news.target_groups ? news.target_groups.map(g => g.id) : []
    }
  } catch (error) {
    console.error('Error loading news:', error)
    appStore.showError('Échec du chargement de l\'actualité')
    router.push('/group-admin/news')
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
  // Validation
  if (form.value.target_group_ids.length === 0) {
    appStore.showError('Veuillez sélectionner au moins un groupe cible')
    return
  }

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
      await groupAdminService.updateNews(route.params.slug, data)
      appStore.showSuccess('Article mis à jour avec succès')
    } else {
      await groupAdminService.createNews(data)
      appStore.showSuccess('Article créé avec succès')
    }

    router.push('/group-admin/news')
  } catch (error) {
    console.error('Error saving news:', error)
    appStore.showError('Échec de l\'enregistrement de l\'article')
  } finally {
    isSaving.value = false
  }
}

// Lifecycle
onMounted(async () => {
  await Promise.all([
    loadMetadata(),
    loadManagedGroups(),
    loadNews()
  ])
})
</script>
