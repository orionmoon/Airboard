<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">{{ $t('applications.title') }}</h1>
          <p class="page-subtitle">{{ $t('applications.subtitle') }}</p>
        </div>
        <button @click="openCreateModal" class="btn btn-primary">
          <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
          {{ $t('applications.new') }}
        </button>
      </div>
    </div>

    <!-- Filters -->
    <div class="filters-container">
      <div class="search-container flex-1">
        <Icon icon="mdi:magnify" class="search-icon" />
        <input
          v-model="searchQuery"
          type="text"
          :placeholder="$t('applications.searchPlaceholder')"
          class="form-input search-input"
        />
      </div>
      <select v-model="groupFilter" class="form-select w-full sm:w-auto">
        <option value="">{{ $t('applications.filter_all_groups') }}</option>
        <option v-for="group in appGroups" :key="group.id" :value="group.id">
          {{ group.name }}
        </option>
      </select>
      <select v-model="statusFilter" class="form-select w-full sm:w-auto">
        <option value="">{{ $t('common.status_all') }}</option>
        <option value="active">{{ $t('common.status_active') }}</option>
        <option value="inactive">{{ $t('common.status_inactive') }}</option>
      </select>
    </div>

    <!-- Content -->
    <div v-if="appStore.loading" class="flex items-center justify-center py-12">
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
            <button @click="editApplication(app)" class="btn-ghost" :title="$t('common.edit')">
              <Icon icon="mdi:pencil" class="h-4 w-4" />
            </button>
            <button @click="confirmDelete(app)" class="btn-ghost text-red-400 hover:text-red-300" :title="$t('common.delete')">
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
              {{ app.is_active ? $t('common.active') : $t('common.inactive') }}
            </span>
            <span class="text-gray-500">{{ $t('applications.order') }}: {{ app.order || 0 }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty States -->
    <div v-else-if="applications.length === 0" class="empty-state">
      <Icon icon="mdi:application" class="empty-state-icon" />
      <h3 class="empty-state-title">{{ $t('applications.emptyTitle') }}</h3>
      <p class="empty-state-description">{{ $t('applications.emptyText') }}</p>
      <button @click="openCreateModal" class="btn btn-primary">
        <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
        {{ $t('applications.add') }}
      </button>
    </div>

    <div v-else class="empty-state">
      <Icon icon="mdi:magnify" class="empty-state-icon" />
      <h3 class="empty-state-title">{{ $t('applications.noResultsTitle') }}</h3>
      <p class="empty-state-description">{{ $t('applications.noResultsText') }}</p>
    </div>

    <!-- Application Modal -->
    <ApplicationModal
      :show="showModal"
      :application="selectedApplication"
      :appGroups="appGroups"
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
                <h3 class="modal-title">{{ $t('applications.confirmDeleteTitle') }}</h3>
                <p class="modal-subtitle">
                  {{ $t('applications.confirmDeleteText', { name: applicationToDelete?.name }) }}
                </p>
              </div>
            </div>
          </div>
          
          <div class="modal-footer">
            <button @click="closeDeleteModal" class="btn btn-secondary w-full sm:w-auto">
              Cancel
            </button>
            <button
              @click="deleteApplication"
              :disabled="deleteLoading"
              class="btn btn-danger w-full sm:w-auto"
            >
              <Icon v-if="deleteLoading" icon="mdi:loading" class="animate-spin h-4 w-4 mr-2" />
              Delete
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
import { adminService } from '@/services/api'
import { useAppStore } from '@/stores/app'
import ApplicationModal from '@/components/admin/ApplicationModal.vue'

const appStore = useAppStore()

// State
const applications = ref([])
const appGroups = ref([])
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
    appStore.setLoading(true)
    // Request all applications by setting a high limit
    const data = await adminService.getApplications({ limit: 1000 })
    console.log('Raw applications API response:', data)
    applications.value = Array.isArray(data) ? data : []
    console.log('Applications loaded:', applications.value)
  } catch (error) {
    console.error('Error loading applications:', error)
    applications.value = []
    appStore.showError('Failed to load applications')
  } finally {
    appStore.setLoading(false)
  }
}

const loadAppGroups = async () => {
  try {
    const data = await adminService.getAppGroups()
    appGroups.value = Array.isArray(data) ? data : []
  } catch (error) {
    console.error('Error loading app groups:', error)
    appGroups.value = []
  }
}

const getAppGroupName = (appGroupId) => {
  const group = appGroups.value.find(g => g.id === appGroupId)
  return group ? group.name : 'No Group'
}

const openCreateModal = () => {
  selectedApplication.value = null
  showModal.value = true
}

const editApplication = (application) => {
  selectedApplication.value = application
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedApplication.value = null
}

const confirmDelete = (application) => {
  applicationToDelete.value = application
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  applicationToDelete.value = null
}

const handleSubmit = async (formData) => {
  try {
    appStore.setLoading(true)
    
    if (formData.id) {
      await adminService.updateApplication(formData.id, formData)
      appStore.showSuccess('Application updated successfully')
    } else {
      await adminService.createApplication(formData)
      appStore.showSuccess('Application created successfully')
    }
    
    closeModal()
    await loadApplications()
  } catch (error) {
    console.error('Error saving application:', error)
    appStore.showError('Failed to save application')
  } finally {
    appStore.setLoading(false)
  }
}

const deleteApplication = async () => {
  if (!applicationToDelete.value || !applicationToDelete.value.id) return
  
  try {
    deleteLoading.value = true
    await adminService.deleteApplication(applicationToDelete.value.id)
    appStore.showSuccess('Application deleted successfully')
    closeDeleteModal()
    await loadApplications()
  } catch (error) {
    console.error('Error deleting application:', error)
    appStore.showError('Failed to delete application')
  } finally {
    deleteLoading.value = false
  }
}

// Lifecycle
onMounted(async () => {
  await Promise.all([
    loadApplications(),
    loadAppGroups()
  ])
})
</script>