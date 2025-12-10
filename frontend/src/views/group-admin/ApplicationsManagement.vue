<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">Applications</h1>
          <p class="page-subtitle">Gérez les applications de vos groupes</p>
        </div>
        <button @click="openCreateModal" class="btn btn-primary" :disabled="appGroups.length === 0">
          <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
          Nouvelle Application
        </button>
      </div>
    </div>

    <!-- Info Alert -->
    <div v-if="appGroups.length === 0" class="mb-6 p-4 bg-yellow-900/20 border border-yellow-600 rounded-lg">
      <div class="flex items-start">
        <Icon icon="mdi:information" class="h-5 w-5 text-yellow-500 mr-3 mt-0.5" />
        <div>
          <h3 class="text-yellow-500 font-medium">Aucun groupe d'applications</h3>
          <p class="text-sm text-yellow-200 mt-1">
            Vous devez d'abord créer un groupe d'applications avant de pouvoir ajouter des applications.
          </p>
        </div>
      </div>
    </div>

    <!-- Filters -->
    <div class="filters-container">
      <div class="search-container flex-1">
        <Icon icon="mdi:magnify" class="search-icon" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Rechercher des applications..."
          class="form-input search-input"
        />
      </div>
      <select v-model="groupFilter" class="form-select w-full sm:w-auto">
        <option value="">Tous les groupes</option>
        <option v-for="group in appGroups" :key="group.id" :value="group.id">
          {{ group.name }}
        </option>
      </select>
      <select v-model="statusFilter" class="form-select w-full sm:w-auto">
        <option value="">Tous les statuts</option>
        <option value="active">Actifs uniquement</option>
        <option value="inactive">Inactifs uniquement</option>
      </select>
    </div>

    <!-- Content -->
    <div v-if="loading" class="flex items-center justify-center py-12">
      <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-gray-400" />
    </div>

    <div v-else-if="filteredApplications.length > 0" class="grid-container grid-cols-auto">
      <div
        v-for="app in filteredApplications"
        :key="app.id"
        class="card hover:border-gray-600 transition-all duration-200 fade-in"
      >
        <div class="card-header">
          <div class="flex items-center gap-3">
            <div
              class="h-10 w-10 rounded-lg flex items-center justify-center"
              :style="{ backgroundColor: app.color || '#6366f1' }"
            >
              <Icon :icon="app.icon || 'mdi:application'" class="h-5 w-5 text-white" />
            </div>
            <div class="flex-1">
              <h3 class="card-title">{{ app.name }}</h3>
              <p class="text-sm text-gray-400">{{ getAppGroupName(app.app_group_id) }}</p>
            </div>
          </div>

          <div class="flex gap-2">
            <button @click="editApplication(app)" class="btn-ghost" title="Modifier">
              <Icon icon="mdi:pencil" class="h-4 w-4" />
            </button>
            <button @click="confirmDelete(app)" class="btn-ghost text-red-400 hover:text-red-300" title="Supprimer">
              <Icon icon="mdi:delete" class="h-4 w-4" />
            </button>
          </div>
        </div>

        <div class="card-content">
          <p v-if="app.description" class="text-sm mb-3">{{ app.description }}</p>

          <div class="mb-3">
            <a
              :href="app.url"
              :target="app.open_in_new_tab ? '_blank' : '_self'"
              class="text-sm text-green-400 hover:text-green-300 flex items-center space-x-1 group"
              :title="app.url"
            >
              <span class="truncate">{{ app.url }}</span>
              <Icon v-if="app.open_in_new_tab" icon="mdi:open-in-new" class="h-3 w-3 flex-shrink-0" />
            </a>
          </div>

          <div class="flex items-center justify-between text-sm">
            <span :class="app.is_active ? 'badge badge-success' : 'badge badge-secondary'">
              {{ app.is_active ? 'Actif' : 'Inactif' }}
            </span>
            <span class="text-gray-500">Ordre: {{ app.order || 0 }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty States -->
    <div v-else-if="applications.length === 0" class="empty-state">
      <Icon icon="mdi:application" class="empty-state-icon" />
      <h3 class="empty-state-title">Aucune application</h3>
      <p class="empty-state-description">Créez votre première application pour commencer.</p>
      <button
        @click="openCreateModal"
        class="btn btn-primary"
        :disabled="appGroups.length === 0"
      >
        <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
        Ajouter une application
      </button>
    </div>

    <div v-else class="empty-state">
      <Icon icon="mdi:magnify" class="empty-state-icon" />
      <h3 class="empty-state-title">Aucun résultat</h3>
      <p class="empty-state-description">Aucune application ne correspond à vos critères de recherche.</p>
    </div>

    <!-- Application Modal -->
    <ApplicationModal
      :show="showModal"
      :application="selectedApplication"
      :app-groups="appGroups"
      @close="closeModal"
      @submit="handleSubmit"
    />

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
                <h3 class="modal-title">Supprimer l'application</h3>
                <p class="modal-subtitle">
                  Êtes-vous sûr de vouloir supprimer "<strong>{{ applicationToDelete?.name }}</strong>" ? Cette action est irréversible.
                </p>
              </div>
            </div>
          </div>

          <div class="modal-footer">
            <button @click="closeDeleteModal" class="btn btn-secondary w-full sm:w-auto">
              Annuler
            </button>
            <button
              @click="deleteApplication"
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
import { ref, computed, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { groupAdminService } from '@/services/api'
import ApplicationModal from '@/components/group-admin/ApplicationModal.vue'

// State
const applications = ref([])
const appGroups = ref([])
const loading = ref(false)
const showModal = ref(false)
const selectedApplication = ref(null)
const showDeleteModal = ref(false)
const applicationToDelete = ref(null)
const deleteLoading = ref(false)
const searchQuery = ref('')
const groupFilter = ref('')
const statusFilter = ref('')

// Computed
const filteredApplications = computed(() => {
  let filtered = applications.value

  if (!Array.isArray(filtered)) {
    return []
  }

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(app =>
      app.name.toLowerCase().includes(query) ||
      app.description?.toLowerCase().includes(query) ||
      app.url.toLowerCase().includes(query)
    )
  }

  // Filter by app group
  if (groupFilter.value) {
    filtered = filtered.filter(app => app.app_group_id === parseInt(groupFilter.value))
  }

  // Filter by status
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(app => app.is_active)
  } else if (statusFilter.value === 'inactive') {
    filtered = filtered.filter(app => !app.is_active)
  }

  return filtered.sort((a, b) => (a.order || 0) - (b.order || 0))
})

// Methods
const loadApplications = async () => {
  try {
    loading.value = true
    const data = await groupAdminService.getApplications({ limit: 1000 })
    applications.value = Array.isArray(data) ? data : (data.data || [])
  } catch (error) {
    console.error('Error loading applications:', error)
    applications.value = []
  } finally {
    loading.value = false
  }
}

const loadAppGroups = async () => {
  try {
    const data = await groupAdminService.getAppGroups()
    appGroups.value = Array.isArray(data) ? data : (data.data || [])
  } catch (error) {
    console.error('Error loading app groups:', error)
    appGroups.value = []
  }
}

const getAppGroupName = (appGroupId) => {
  const group = appGroups.value.find(g => g.id === appGroupId)
  return group ? group.name : 'Aucun groupe'
}

const openCreateModal = () => {
  if (appGroups.value.length === 0) {
    return
  }
  selectedApplication.value = null
  showModal.value = true
}

const editApplication = (app) => {
  selectedApplication.value = { ...app }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedApplication.value = null
}

const handleSubmit = async (formData) => {
  try {
    if (formData.id) {
      await groupAdminService.updateApplication(formData.id, formData)
    } else {
      await groupAdminService.createApplication(formData)
    }
    closeModal()
    await loadApplications()
  } catch (error) {
    console.error('Error saving application:', error)
    throw error
  }
}

const confirmDelete = (app) => {
  applicationToDelete.value = app
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  applicationToDelete.value = null
}

const deleteApplication = async () => {
  if (!applicationToDelete.value) return

  try {
    deleteLoading.value = true
    await groupAdminService.deleteApplication(applicationToDelete.value.id)
    closeDeleteModal()
    await loadApplications()
  } catch (error) {
    console.error('Error deleting application:', error)
    alert('Erreur lors de la suppression de l\'application')
  } finally {
    deleteLoading.value = false
  }
}

onMounted(async () => {
  await loadAppGroups()
  await loadApplications()
})
</script>
