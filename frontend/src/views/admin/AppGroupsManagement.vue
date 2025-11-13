<template>
  <div class="content-area">
    <!-- Header -->
    <div class="page-header">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="page-title">App Groups</h1>
          <p class="page-subtitle">Organize your applications into groups</p>
        </div>
        <button @click="openCreateModal" class="btn btn-primary">
          <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
          New Group
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
          placeholder="Search groups..."
          class="form-input search-input"
        />
      </div>
      <select v-model="statusFilter" class="form-select w-full sm:w-auto">
        <option value="">All Status</option>
        <option value="active">Active Only</option>
        <option value="inactive">Inactive Only</option>
      </select>
    </div>

    <!-- Content -->
    <div v-if="appStore.loading" class="flex items-center justify-center py-12">
      <Icon icon="mdi:loading" class="h-8 w-8 animate-spin text-gray-400" />
    </div>

    <div v-else-if="filteredAppGroups.length > 0" class="grid-container grid-cols-auto">
      <div
        v-for="group in filteredAppGroups"
        :key="group.id"
        class="card hover:border-gray-600 transition-all duration-200 fade-in"
      >
        <div class="card-header">
          <div class="flex items-center space-x-3">
            <div 
              class="h-10 w-10 rounded-lg flex items-center justify-center"
              :style="{ backgroundColor: group.color || '#10b981' }"
            >
              <Icon :icon="group.icon || 'mdi:folder'" class="h-5 w-5 text-white" />
            </div>
            <div>
              <h3 class="card-title">{{ group.name }}</h3>
              <p class="text-sm text-gray-400">{{ group.applications?.length || 0 }} applications</p>
            </div>
          </div>
          
          <div class="flex space-x-2">
            <button @click="editGroup(group)" class="btn-ghost" title="Edit">
              <Icon icon="mdi:pencil" class="h-4 w-4" />
            </button>
            <button @click="confirmDelete(group)" class="btn-ghost text-red-400 hover:text-red-300" title="Delete">
              <Icon icon="mdi:delete" class="h-4 w-4" />
            </button>
          </div>
        </div>
        
        <div class="card-content">
          <p v-if="group.description" class="text-sm mb-4">{{ group.description }}</p>
          
          <div class="flex items-center justify-between text-sm">
            <span :class="group.is_active ? 'badge badge-success' : 'badge badge-secondary'">
              {{ group.is_active ? 'Active' : 'Inactive' }}
            </span>
            <span class="text-gray-500">Order: {{ group.order || 0 }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty States -->
    <div v-else-if="appGroups.length === 0" class="empty-state">
      <Icon icon="mdi:folder-multiple" class="empty-state-icon" />
      <h3 class="empty-state-title">No App Groups</h3>
      <p class="empty-state-description">Create your first application group to get started.</p>
      <button @click="openCreateModal" class="btn btn-primary">
        <Icon icon="mdi:plus" class="h-4 w-4 mr-2" />
        Create Group
      </button>
    </div>

    <div v-else class="empty-state">
      <Icon icon="mdi:magnify" class="empty-state-icon" />
      <h3 class="empty-state-title">No Results</h3>
      <p class="empty-state-description">No groups match your search criteria.</p>
    </div>

    <!-- App Group Modal -->
    <AppGroupModal
      :show="showModal"
      :app-group="selectedAppGroup"
      @close="closeModal"
      @submit="handleSubmit"
    />

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="modal-overlay">
      <div class="modal-container">
        <div class="modal-panel sm:max-w-lg sm:w-full">
          <div class="modal-header">
            <div class="flex items-start space-x-3">
              <div class="flex-shrink-0">
                <Icon icon="mdi:alert" class="h-6 w-6 text-red-500" />
              </div>
              <div>
                <h3 class="modal-title">Delete App Group</h3>
                <p class="modal-subtitle">
                  Are you sure you want to delete "<strong>{{ appGroupToDelete?.name }}</strong>"? This action cannot be undone.
                </p>
              </div>
            </div>
          </div>
          
          <div class="modal-footer">
            <button @click="closeDeleteModal" class="btn btn-secondary w-full sm:w-auto">
              Cancel
            </button>
            <button
              @click="deleteGroup"
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
import AppGroupModal from '@/components/admin/AppGroupModal.vue'

const appStore = useAppStore()

// State
const appGroups = ref([])
const showModal = ref(false)
const selectedAppGroup = ref(null)
const showDeleteModal = ref(false)
const appGroupToDelete = ref(null)
const deleteLoading = ref(false)
const searchQuery = ref('')
const statusFilter = ref('')

// Computed
const filteredAppGroups = computed(() => {
  let filtered = appGroups.value

  if (!Array.isArray(filtered)) {
    return []
  }

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(group => 
      group.name.toLowerCase().includes(query) ||
      group.description?.toLowerCase().includes(query)
    )
  }

  // Filter by status
  if (statusFilter.value === 'active') {
    filtered = filtered.filter(group => group.is_active)
  } else if (statusFilter.value === 'inactive') {
    filtered = filtered.filter(group => !group.is_active)
  }

  return filtered.sort((a, b) => (a.order || 0) - (b.order || 0))
})

// Methods
const loadAppGroups = async () => {
  try {
    appStore.setLoading(true)
    console.log('🔄 Starting to load app groups...')
    const data = await adminService.getAppGroups()
    console.log('📊 Raw API response:', data)
    console.log('📊 Type of data:', typeof data)
    console.log('📊 Is array?', Array.isArray(data))
    
    if (Array.isArray(data)) {
      appGroups.value = data
      console.log('✅ App groups set to:', appGroups.value)
    } else if (data && Array.isArray(data.data)) {
      appGroups.value = data.data
      console.log('✅ App groups set from data.data:', appGroups.value)
    } else {
      console.warn('⚠️ Data is not an array, setting empty:', data)
      appGroups.value = []
    }
    
    console.log('📈 Final appGroups.value:', appGroups.value)
    console.log('📈 Final appGroups.value length:', appGroups.value.length)
  } catch (error) {
    console.error('❌ Error loading app groups:', error)
    appGroups.value = []
    appStore.showError('Failed to load app groups')
  } finally {
    appStore.setLoading(false)
  }
}

const openCreateModal = () => {
  selectedAppGroup.value = null
  showModal.value = true
}

const editGroup = (group) => {
  selectedAppGroup.value = group
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  selectedAppGroup.value = null
}

const confirmDelete = (group) => {
  appGroupToDelete.value = group
  showDeleteModal.value = true
}

const closeDeleteModal = () => {
  showDeleteModal.value = false
  appGroupToDelete.value = null
}

const handleSubmit = async (formData) => {
  try {
    appStore.setLoading(true)
    
    if (formData.id) {
      await adminService.updateAppGroup(formData.id, formData)
      appStore.showSuccess('App group updated successfully')
    } else {
      await adminService.createAppGroup(formData)
      appStore.showSuccess('App group created successfully')
    }
    
    closeModal()
    await loadAppGroups()
  } catch (error) {
    console.error('Error saving app group:', error)
    appStore.showError('Failed to save app group')
  } finally {
    appStore.setLoading(false)
  }
}

const deleteGroup = async () => {
  if (!appGroupToDelete.value || !appGroupToDelete.value.id) return
  
  try {
    deleteLoading.value = true
    await adminService.deleteAppGroup(appGroupToDelete.value.id)
    appStore.showSuccess('App group deleted successfully')
    closeDeleteModal()
    await loadAppGroups()
  } catch (error) {
    console.error('Error deleting app group:', error)
    appStore.showError('Failed to delete app group')
  } finally {
    deleteLoading.value = false
  }
}

// Lifecycle
onMounted(() => {
  loadAppGroups()
})
</script>